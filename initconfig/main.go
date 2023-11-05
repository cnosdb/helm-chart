package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pelletier/go-toml"
	"github.com/tidwall/gjson"
	"k8s.io/apimachinery/pkg/util/json"

	"os"
)

const QueryTskv = "query_tskv"
const Query = "query"
const TSKV = "tskv"
const META = "meta"
const SINGLETON = "singleton"

const checkJsonPath = "membership_config.membership.configs"
const checkJsonPathUnderOk = "Ok.membership_config.membership.configs"

type StartType string
type ContextType string

const Conf StartType = "conf"
const Completion StartType = "completion"

const Helm ContextType = "helm"
const Operator ContextType = "operator"

var clusterNotReadyErr = errors.New("cluster not ready")

func main() {
	var starterType string
	var contextType string
	flag.StringVar(&starterType, "type", "conf", "starter type: conf or completion")
	flag.StringVar(&contextType, "context", "helm", "starter type: conf or completion")
	flag.Parse()
	if starterType == string(Conf) {
		generateConf(contextType)
	} else if starterType == string(Completion) {
		completion()
	} else {
		exitErr(errors.New("unsupported start type: " + starterType))
	}
}

func completion() {
	msg, pass := checkCompletionEnv()
	if !pass {
		exitErr(errors.New(fmt.Sprintf("[completion]The necessary environment variables are missing: %s", msg)))
	}
	client := resty.New()
	upgrade := os.Getenv("UPGRADE")
	namespace := os.Getenv("NAMESPACE")
	metaSvcName := os.Getenv("META_SVC_NAME")
	metaStsName := os.Getenv("META_STS_NAME")
	metaSvcPort := os.Getenv("META_SVC_PORT")
	metaReplicasStr := os.Getenv("META_REPLICAS")
	metaReplicas, err := strconv.Atoi(metaReplicasStr)
	if err != nil {
		exitErr(err)
	}
	metaAddrs := generateMetaAddrs(metaReplicas, metaSvcName, metaStsName, metaSvcPort, namespace)
	baseUrl := fmt.Sprintf("http://%s", metaAddrs[0])
	if upgrade != "true" {
		fmt.Println("=------------install cluster------------=")
		// install - init
		body := make(map[string]string)
		_, reqErr := client.R().SetBody(body).Post(fmt.Sprintf("%s/init", baseUrl))
		if reqErr != nil {
			fmt.Println("init meta failed: " + reqErr.Error())
			for reqErr != nil {
				fmt.Println("retry after 10s...")
				time.Sleep(10 * time.Second)
				_, reqErr = client.R().SetBody(body).Post(fmt.Sprintf("%s/init", baseUrl))
			}
		}
		fmt.Println("=------------init success------------=")
	} else {
		fmt.Println("=------------upgrade cluster------------=")
	}
	// upgrade/install - add-learner,change-membership
	if l := len(metaAddrs); l > 1 {
		var changeMembershipBody []int
		changeMembershipBody = append(changeMembershipBody, 0)
		for i := 1; i < l; i++ {
			changeMembershipBody = append(changeMembershipBody, i)
			var addLearnerBody []interface{}
			addLearnerBody = append(addLearnerBody, i, metaAddrs[i])
			_, reqErr := client.R().SetHeader("Content-Type", "application/json").SetBody(addLearnerBody).Post(fmt.Sprintf("%s/add-learner", baseUrl))
			if reqErr != nil {
				fmt.Println("meta add-learner failed: " + reqErr.Error())
				for reqErr != nil {
					fmt.Println("retry after 10s...")
					time.Sleep(10 * time.Second)
					_, reqErr = client.R().SetHeader("Content-Type", "application/json").SetBody(addLearnerBody).Post(fmt.Sprintf("%s/add-learner", baseUrl))
				}
			}
			fmt.Printf("=------------add-learner %d success------------=\n", i)
		}
		_, reqErr := client.R().SetHeader("Content-Type", "application/json").SetBody(changeMembershipBody).Post(fmt.Sprintf("%s/change-membership", baseUrl))
		if reqErr != nil {
			fmt.Println("meta change-membership failed: " + reqErr.Error())
			for reqErr != nil {
				fmt.Println("retry after 10s...")
				time.Sleep(10 * time.Second)
				_, reqErr = client.R().SetHeader("Content-Type", "application/json").SetBody(changeMembershipBody).Post(fmt.Sprintf("%s/change-membership", baseUrl))
			}
		}
		fmt.Println("=------------change-membership success------------=")
	}
	fmt.Println("=------------all finished------------=")
}

func generateConf(contextType string) {
	role := os.Getenv("CNOSDB_ROLE")
	if role == "" {
		exitErr(errors.New("env CNOSDB_ROLE is missing"))
	}
	msg, pass := checkConfEnv(role)
	if !pass {
		exitErr(errors.New(fmt.Sprintf("[%s]The necessary environment variables are missing: %s", role, msg)))
	}
	baseConfPath := "/etc/initconf/default.conf"
	conf, err := toml.LoadFile(baseConfPath)
	exitErr(err)
	targetConfPath := ""
	if role == META {
		targetConfPath = "/etc/initconf/cnosdb-meta.conf"
	} else {
		targetConfPath = "/etc/initconf/cnosdb.conf"
	}
	f, err := os.Create(targetConfPath)
	exitErr(err)
	defer f.Close()
	err = setConfFromUser(conf, contextType)
	exitErr(err)
	var metaAddr string
	switch role {
	case META:
		err = setMeta(conf)
	case SINGLETON:
		err = setSingleton(conf)
	default:
		metaAddr, err = setTskvOrQuery(role, conf)
	}
	exitErr(err)
	conf.WriteTo(f)
	fmt.Println("=------------generate config finished------------=")
	if os.Getenv("DEBUG") == "true" {
		fmt.Println(conf.String())
	}
	if role != META && role != SINGLETON {
		waitingMeta(metaAddr)
	}
}

func setSingleton(conf *toml.Tree) error {
	clusterName := os.Getenv("CLUSTER_INSTANCE_NAME")
	namespace := os.Getenv("NAMESPACE")
	svcName := os.Getenv("SVC_NAME")
	conf.Set("host", fmt.Sprintf("%s.%s", svcName, namespace))
	conf.Set("deployment.mode", "singleton")
	conf.Set("cluster.name", clusterName)
	return nil
}
func setMeta(conf *toml.Tree) error {
	hostname := os.Getenv("HOSTNAME")
	namespace := os.Getenv("NAMESPACE")
	metaSvcName := os.Getenv("META_SVC_NAME")
	clusterName := os.Getenv("CLUSTER_INSTANCE_NAME")
	id, err := getId(META, hostname)
	if err != nil {
		return err
	}
	conf.Set("host", generateHost(hostname, metaSvcName, namespace, false))
	conf.Set("id", int64(id))
	conf.Set("meta_init.cluster_name", clusterName)
	return nil
}

func setTskvOrQuery(role string, conf *toml.Tree) (string, error) {
	hostname := os.Getenv("HOSTNAME")
	namespace := os.Getenv("NAMESPACE")
	metaSvcName := os.Getenv("META_SVC_NAME")
	metaStsName := os.Getenv("META_STS_NAME")
	metaSvcPort := os.Getenv("META_SVC_PORT")
	metaReplicasStr := os.Getenv("META_REPLICAS")
	metaReplicas, err := strconv.Atoi(metaReplicasStr)
	if err != nil {
		return "", err
	}
	svcName := os.Getenv("SVC_NAME")
	clusterName := os.Getenv("CLUSTER_INSTANCE_NAME")
	id, err := getId(role, hostname)
	if err != nil {
		return "", err
	}
	metaAddrs := generateMetaAddrs(metaReplicas, metaSvcName, metaStsName, metaSvcPort, namespace)
	if err != nil {
		return "", err
	}
	conf.Set("host", generateHost(hostname, svcName, namespace, role == Query))
	conf.Set("node_basic.node_id", int64(id))
	conf.Set("cluster.meta_service_addr", metaAddrs)
	conf.Set("cluster.name", clusterName)
	return metaAddrs[0], nil
}

func setConfFromUser(conf *toml.Tree, contextType string) error {
	if contextType == string(Helm) {
		userConf := os.Getenv("CONF_FROM_USER")
		if userConf == "" || userConf == "{}" {
			return nil
		}
		confMap := make(map[string]interface{})
		err := json.Unmarshal([]byte(userConf), &confMap)
		if err != nil {
			return err
		}
		for k, v := range confMap {
			switch v.(type) {
			case int:
				conf.Set(k, v.(int64))
			case string:
				conf.Set(k, v.(string))
			default:
				conf.Set(k, v)
			}
		}
	} else if contextType == string(Operator) {
		userConfPath := "/etc/initconf/user.conf"
		userConf, err := toml.LoadFile(userConfPath)
		if err != nil {
			return err
		}
		paths := getTomlPaths(userConf)
		for _, k := range paths {
			conf.Set(k, userConf.Get(k))
		}
	} else {
		return errors.New("unsupported start type: " + contextType)
	}
	return nil
}

func getId(role, hostname string) (int, error) {
	idx := strings.LastIndex(hostname, "-")
	idstr := hostname[idx+1:]
	if role == Query {
		id := 1000
		for _, a := range idstr {
			id = id + int(a)
		}
		return id, nil
	}
	return strconv.Atoi(idstr)
}

func generateMetaAddrs(replicas int, metaSvcName, metaStsName, metaSvcPort, namespace string) []string {
	var addrs []string
	for i := 0; i < replicas; i++ {
		addrs = append(addrs, fmt.Sprintf("%s-%d.%s.%s:%s", metaStsName, i, metaSvcName, namespace, metaSvcPort))
	}
	return addrs
}

func generateHost(hostname, svcName, namespace string, isQuery bool) string {
	// query 是deployment部署的
	if isQuery {
		return fmt.Sprintf("%s.%s", svcName, namespace)
	}
	return fmt.Sprintf("%s.%s.%s", hostname, svcName, namespace)
}

func checkConfEnv(role string) (string, bool) {
	var envs []string
	switch role {
	case Query:
		fallthrough
	case TSKV:
		fallthrough
	case QueryTskv:
		envs = append(envs, "HOSTNAME", "NAMESPACE", "META_SVC_NAME", "META_STS_NAME", "META_SVC_PORT", "META_REPLICAS", "SVC_NAME", "CLUSTER_INSTANCE_NAME")
	case META:
		envs = append(envs, "HOSTNAME", "NAMESPACE", "META_SVC_NAME", "CLUSTER_INSTANCE_NAME")
	case SINGLETON:
		envs = append(envs, "NAMESPACE", "SVC_NAME", "CLUSTER_INSTANCE_NAME")
	}
	return doCheckEnv(envs)
}

func checkCompletionEnv() (string, bool) {
	var envs []string
	envs = append(envs, "NAMESPACE", "META_SVC_NAME", "META_STS_NAME", "META_SVC_PORT", "META_REPLICAS", "UPGRADE")
	return doCheckEnv(envs)
}

func doCheckEnv(envs []string) (string, bool) {
	var msgs []string
	for _, env := range envs {
		v := os.Getenv(env)
		if v == "" {
			msgs = append(msgs, env)
		}
	}
	if len(msgs) > 0 {
		return strings.Join(msgs, ","), false
	}
	return "", true
}

func waitingMeta(metaAddr string) {
	client := resty.New()
	url := fmt.Sprintf("http://%s/metrics", metaAddr)
	resp, reqErr := client.R().Get(url)
	var configs gjson.Result
	if reqErr == nil {
		ok := gjson.Get(resp.String(), "Ok")
		if ok.Exists() {
			configs = gjson.Get(resp.String(), checkJsonPathUnderOk)
		} else {
			configs = gjson.Get(resp.String(), checkJsonPath)
		}
		if !configs.Exists() || len(configs.Array()) == 0 {
			reqErr = clusterNotReadyErr
		}
	}
	for reqErr != nil {
		fmt.Printf("meta not ready: %s\n", reqErr.Error())
		fmt.Println("waiting meta for 10s....")
		time.Sleep(10 * time.Second)
		fmt.Printf("check meta: %s\n", url)
		resp, reqErr = client.R().Get(url)
		if reqErr == nil {
			fmt.Println(resp.String())
			ok := gjson.Get(resp.String(), "Ok")
			if ok.Exists() {
				configs = gjson.Get(resp.String(), checkJsonPathUnderOk)
			} else {
				configs = gjson.Get(resp.String(), checkJsonPath)
			}
			if configs.Exists() && len(configs.Array()) > 0 {
				continue
			}
			reqErr = clusterNotReadyErr
		}
	}

	fmt.Println("meta is up let's go!")
}

func exitErr(err error) {
	if err != nil {
		fmt.Printf("found error：%v", err)
		os.Exit(1)
	}
}

func getTomlPaths(tree *toml.Tree) []string {
	var result []string
	m := tree.ToMap()
	getTomlPathsRecursive(m, "", &result)
	return result
}
func getTomlPathsRecursive(current any, path string, result *[]string) {
	switch c := current.(type) {
	case nil:
	case map[string]interface{}:
		for k := range c {
			if path == "" {
				getTomlPathsRecursive(c[k], k, result)
			} else {
				getTomlPathsRecursive(c[k], path+"."+k, result)
			}
		}
	default:
		*result = append(*result, path)
	}
	return
}
