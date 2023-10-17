# Cnosdb
CnosDB 是一款高性能、高压缩率、高易用性的开源分布式时序数据库。主要应用场景为物联网，工业互联网，车联网，IT运维等。所有代码均已在GitHub开源。

[English](./README.md) | 简体中文

## 快速体验
```sh
helm repo add cnosdb https://cnosdb.github.io/helm-chart/
helm repo update cnosdb
helm install my-cnosdb cnosdb/cnosdb
```
## 介绍
这个chart 可以在一个集群中部署一个cnosdb

## 环境要求
- Kubernetes 1.19+
- 集群基础设施支持PV供应商

## 安装Chart
在`cnosdb`命名空间中安装一个叫`my-cnosdb`的实例
```sh
helm install my-cnosdb cnosdb/cnosdb -ncnosdb
```

## 卸载Chart

将 `my-cnosdb` 实例删除:

```sh
helm delete my-cnosdb -ncnosdb
```
这个命令会移除所有和这个chart相关联的k8s资源并且删除`release`

## 参数配置

在安装helm chart的时候可以使用 `--set key=value[,key=value]`指定任意的参数. 例如:
```sh
helm install my-cnosdb \
--set image.pullPolicy=IfNotPresent \
cnosdb/cnosdb -ncnosdb
```
或者,使用一个`values.yaml`这样的YAML文件来指定安装的时候所使用的参数,例如:
```sh
helm install my-cnosdb -f values.yaml cnosdb/cnosdb -ncnosdb
```
### 镜像参数

| 名称                             | 描述                      | 默认值                   |
| -------------------------------- | ------------------------- | ------------------------ |
| image.cnosdbMeta.repository      | Cnosdb meta 镜像仓库      | cnosdb/cnosdb-meta       |
| image.cnosdbMeta.tag             | Cnosdb meta 镜像tag       | community-latest         |
| image.cnosdb.repository          | Cnosdb 镜像仓库           | cnosdb/cnosdb            |
| image.cnosdb.tag                 | Cnosdb 镜像tag            | community-latest         |
| image.clusterConfTool.repository | Cnosdb 集群工具的镜像仓库 | wyxok/cnosdb-init-config |
| image.clusterConfTool.tag        | Cnosdb 集群工具的镜像tag  | latest                   |
| image.pullPolicy                 | 镜像拉取策略              | IfNotPresent             |
| image.pullSecrets                | 镜像拉取秘钥              | []                       |

### 通用参数

| 名称             | 描述                                   | 默认值  |
| ---------------- | -------------------------------------- | ------- |
| nameOverride     | 部分覆盖 common.names.fullname的字符串 | ""      |
| fullnameOverride | 完全覆盖 common.names.fullname的字符串 | ""      |
| architecture     | 安装的架构(目前只支持cluster)          | cluster |

### Meta参数

| 名称                                  | 描述                               | 默认值            |
| ------------------------------------- | ---------------------------------- | ----------------- |
| meta.replicaCount                     | Cnosdb meta 部署的副本数           | 3                 |
| meta.terminationGracePeriodSeconds    | 优雅终结Cnosdb meta 副本pod的时间  | 10                |
| meta.extraConf                        | Cnosdb meta 副本节点的覆盖配置     | {}                |
| meta.resources.limits                 | Cnosdb meta 副本容器的资源限制     | {}                |
| meta.resources.requests               | Cnosdb meta 副本容器资源请求       | {}                |
| meta.affinity                         | Cnosdb meta 副本pods的亲和性配置   | {}                |
| meta.nodeSelector                     | Cnosdb meta 副本pods的节点选择配置 | {}                |
| meta.tolerations                      | Cnosdb meta 副本pods容忍配置       | []                |
| meta.service.type                     | Cnosdb meta 服务类型               | ClusterIP         |
| meta.service.port                     | Cnosdb meta 服务的端口             | 8901              |
| meta.service.nodePort                 | Cnosdb meta 服务的节点端口         | ""                |
| meta.service.clusterIP                | Cnosdb meta 服务的集群IP           | ""                |
| meta.service.externalTrafficPolicy    | Cnosdb meta 服务外部流量策略       | Cluster           |
| meta.service.annotations              | Cnosdb meta 服务的额外注释         | {}                |
| meta.service.loadBalancerIP           | Cnosdb meta 服务负载均衡 IP        | ""                |
| meta.service.loadBalancerSourceRanges | Cnosdb meta 服务负载均衡源         | []                |
| meta.persistence.enabled              | 持久化存储的开关                   | false             |
| meta.persistence.storageClass         | 持久化存储的存储类                 | ""                |
| meta.persistence.accessModes          | 持久化存储的访问模式               | ["ReadWriteOnce"] |
| meta.persistence.size                 | 持久化存储的大小                   | 1Gi               |
| meta.persistence.existingClaim        | 使用已存在的持久化存储声明         | ""                |

### Tskv参数

| 名称                                  | 描述                                 | 默认值            |
| ------------------------------------- | ------------------------------------ | ----------------- |
| tskv.replicaCount                     | Cnosdb tskv 部署的副本数             | 2                 |
| tskv.terminationGracePeriodSeconds    | 优雅终结Cnosdb tskv 副本pod的时间    | 10                |
| tskv.extraConf                        | Cnosdb tskv 副本节点的覆盖配置       | {}                |
| tskv.resources.limits                 | Cnosdb tskv 副本容器的资源限制       | {}                |
| tskv.resources.requests               | Cnosdb tskv 副本容器资源请求         | {}                |
| tskv.affinity                         | Cnosdb tskv 副本pods的亲和性配置     | {}                |
| tskv.nodeSelector                     | Cnosdb tskv 副本pods的节点选择配置   | {}                |
| tskv.tolerations                      | Cnosdb tskv 副本pods容忍配置         | []                |
| tskv.service.type                     | Cnosdb tskv 服务类型                 | ClusterIP         |
| tskv.service.ports.http               | Cnosdb tskv 服务的http端口           | 8902              |
| tskv.service.ports.grpc               | Cnosdb tskv 服务的grpc端口           | 8903              |
| tskv.service.ports.flight             | Cnosdb tskv 服务的flight rpc端口     | 8904              |
| tskv.service.ports.tcp                | Cnosdb tskv 服务的tcp端口            | 8905              |
| tskv.service.ports.vector             | Cnosdb tskv 服务的vector端口         | 8906              |
| tskv.service.nodePorts.http           | Cnosdb tskv 服务的http节点端口       | ""                |
| tskv.service.nodePorts.grpc           | Cnosdb tskv 服务的grpc节点端口       | ""                |
| tskv.service.nodePorts.flight         | Cnosdb tskv 服务的flight rpc节点端口 | ""                |
| tskv.service.nodePorts.tcp            | Cnosdb tskv 服务的tcp节点端口        | ""                |
| tskv.service.nodePorts.vector         | Cnosdb tskv 服务的vector节点端口     | ""                |
| tskv.service.clusterIP                | Cnosdb tskv 服务的集群IP             | ""                |
| tskv.service.externalTrafficPolicy    | Cnosdb tskv 服务外部流量策略         | Cluster           |
| tskv.service.annotations              | Cnosdb tskv 服务的额外注释           | {}                |
| tskv.service.loadBalancerIP           | Cnosdb tskv 服务负载均衡 IP          | ""                |
| tskv.service.loadBalancerSourceRanges | Cnosdb tskv 服务负载均衡源           | []                |
| tskv.persistence.enabled              | 持久化存储的开关                     | false             |
| tskv.persistence.storageClass         | 持久化存储的存储类                   | ""                |
| tskv.persistence.accessModes          | 持久化存储的访问模式                 | ["ReadWriteOnce"] |
| tskv.persistence.size                 | 持久化存储的大小                     | 1Gi               |
| tskv.persistence.existingClaim        | 使用已存在的持久化存储声明           | ""                |

### Query参数

| 名称                                   | 描述                                  | 默认值    |
| -------------------------------------- | ------------------------------------- | --------- |
| query.replicaCount                     | Cnosdb query 部署的副本数             | 2         |
| query.extraConf                        | Cnosdb query 副本节点的覆盖配置       | {}        |
| query.resources.limits                 | Cnosdb query 副本容器的资源限制       | {}        |
| query.resources.requests               | Cnosdb query 副本容器资源请求         | {}        |
| query.affinity                         | Cnosdb query 副本pods的亲和性配置     | {}        |
| query.nodeSelector                     | Cnosdb query 副本pods的节点选择配置   | {}        |
| query.tolerations                      | Cnosdb query 副本pods容忍配置         | []        |
| query.service.type                     | Cnosdb query 服务类型                 | ClusterIP |
| query.service.ports.http               | Cnosdb query 服务的http端口           | 8902      |
| query.service.ports.grpc               | Cnosdb query 服务的grpc端口           | 8903      |
| query.service.ports.flight             | Cnosdb query 服务的flight rpc端口     | 8904      |
| query.service.ports.tcp                | Cnosdb query 服务的tcp端口            | 8905      |
| query.service.ports.vector             | Cnosdb query 服务的vector端口         | 8906      |
| query.service.nodePorts.http           | Cnosdb query 服务的http节点端口       | ""        |
| query.service.nodePorts.grpc           | Cnosdb query 服务的grpc节点端口       | ""        |
| query.service.nodePorts.flight         | Cnosdb query 服务的flight rpc节点端口 | ""        |
| query.service.nodePorts.tcp            | Cnosdb query 服务的tcp节点端口        | ""        |
| query.service.nodePorts.vector         | Cnosdb query 服务的vector节点端口     | ""        |
| query.service.clusterIP                | Cnosdb query 服务的集群IP             | ""        |
| query.service.externalTrafficPolicy    | Cnosdb query 服务外部流量策略         | Cluster   |
| query.service.annotations              | Cnosdb query 服务的额外注释           | {}        |
| query.service.loadBalancerIP           | Cnosdb query 服务负载均衡 IP          | ""        |
| query.service.loadBalancerSourceRanges | Cnosdb query 服务负载均衡源           | []        |


## 提示
在删除chart实例的时候不会移除PV

### Persistence

chart 会挂载 [持久卷](https://kubernetes.io/docs/concepts/storage/persistent-volumes/) 到`/var/lib/cnosdb` 目录. 默认情况下,会动态创建`pvc`,但是也可以指定一个已存在的`pvc`, 如果已经有现成的pvc了, 那么可以在安装的时候可以指定

### Existing PersistentVolumeClaims

1. 创建 PersistentVolume
2. 创建 PersistentVolumeClaim
3. 安装 chart
```sh
helm install --set meta.persistence.existingClaim=PVC_NAME my-cnosdb cnosdb/cnosdb
```

### 额外配置
toml 配置可以通过一行表达式表示, `demo.foo=bar` 和下面是相等的
```toml
# demo.foo=bar
[demo]
foo = bar
```
可以通过设置extraConf的值来覆盖默认配置:
```sh
helm install \
--set meta.extraConf.'storage\.maxsummary_size'='64M' \
--set tskv.extraConf.'storage\.max_level'=1 \
my-cnosdb cnosdb/cnosdb
```


### 升级Chart的实例

只更新镜像的版本

```sh
helm upgrade my-cnosdb cnosdb/cnosdb -ncnosdb --reuse-values --set image.cnosdb.tag=new.version
```

执行水平扩容

```sh
helm upgrade my-cnosdb cnosdb/cnosdb -ncnosdb --reuse-values --set meta.replicaCount=3 --set tskv.replicaCount=5 
```

如果集群支持动态 [扩展 PVC](https://kubernetes.io/docs/concepts/storage/persistent-volumes/#expanding-persistent-volumes-claims), 你可以对存储进行垂直扩容(不借助helm). 但是目前helm做不到, 因为helm无法处理statefulset的验证错误, 而目前statefulset不支持修改vct中的pvc的大小. 所以你只能对 `resources` 例如 `cpu` 和 `memory`进行垂直扩容.

```sh
helm upgrade my-cnosdb cnosdb/cnosdb -ncnosdb --reuse-values --set tskv.resources.limits.cpu=1
```

Chart会在本地缓存一份, 如果你想更新Chart(不是应用)的版本, 你应该从远程仓库同步一下Chart的信息.

```sh
helm repo update cnosdb
```

然后就可以根据需要进行升级

```sh
helm upgrade my-cnosdb cnosdb/cnosdb -ncnosdb --set foo=bar
```