# CnosDB
CnosDB is a high-performance, high-compression, and easy-to-use open-source distributed time-series database. It is primarily used in fields such as IoT, industrial internet, connected cars, and IT operations. All of the code is open-sourced and available on GitHub.

English | [简体中文](./README_CN.md)
## TL;DR
```sh
helm repo add cnosdb https://cnosdb.github.io/helm-chart/
helm repo update cnosdb
helm install my-cnosdb cnosdb/cnosdb
```

## Introduction
This chart bootstraps a Cnosdb deployment on a cluster using the package manager.

## Prerequisites
- Kubernetes 1.19+ 
- PV provisioner support in the underlying infrastructure

## Installing the Chart
To install the chart with the release name `my-cnosdb` In `cnosdb` namespace
```sh
helm install my-cnosdb cnosdb/cnosdb -ncnosdb
```

## Uninstalling the Chart

To uninstall/delete the `my-cnosdb` deployment:

```sh
helm delete my-cnosdb -ncnosdb
```
The command removes all the Kubernetes components associated with the chart and deletes the release.

## Parameters

Specify each parameter using the --set key=value[,key=value] argument to helm install. For example,
```sh
helm install my-cnosdb \
--set image.pullPolicy=IfNotPresent \
cnosdb/cnosdb -ncnosdb
```
Alternatively, a YAML file that specifies the values for the parameters can be provided while installing the chart. For example,
```sh
helm install my-cnosdb -f values.yaml cnosdb/cnosdb -ncnosdb
```
### Image Parameters

| Name                             | Description                          | Value                    |
| -------------------------------- | ------------------------------------ | ------------------------ |
| image.cnosdbMeta.repository      | Cnosdb meta image repository         | cnosdb/cnosdb-meta       |
| image.cnosdbMeta.tag             | Cnosdb meta image tag                | community-latest         |
| image.cnosdb.repository          | Cnosdb image repository              | cnosdb/cnosdb            |
| image.cnosdb.tag                 | Cnosdb image tag                     | community-latest         |
| image.clusterConfTool.repository | Cnosdb cluster tool image repository | wyxok/cnosdb-init-config |
| image.clusterConfTool.tag        | Cnosdb cluster tool image tag        | latest                   |
| image.pullPolicy                 | Image pull policy                    | IfNotPresent             |
| image.pullSecrets                | Image pull secrets                   | []                       |


### Common Parameters

| Name             | Description                                                                     | Value      |
| ---------------- | ------------------------------------------------------------------------------- | ---------- |
| nameOverride     | String to partially override common.names.fullname                              | ""         |
| fullnameOverride | String to fully override common.names.fullname                                  | ""         |
| architecture     | The architecture of installing. Allowed values: separation, bundle or singleton | separation |

### Meta Parameters
**Active when architecture is `separation` or `bundle`**

| Name                                  | Description                                                                     | Value             |
| ------------------------------------- | ------------------------------------------------------------------------------- | ----------------- |
| meta.replicaCount                     | Number of Cnosdb meta replicas to deploy                                        | 3                 |
| meta.terminationGracePeriodSeconds    | Integer setting the termination grace period for the Cnosdb meta-replicas pods  | 10                |
| meta.extraConf                        | Configuration for Cnosdb meta replicas nodes                                    | {}                |
| meta.resources.limits                 | The resources limits for the Cnosdb meta replicas containers                    | {}                |
| meta.resources.requests               | The requested resources for the Cnosdb meta replicas containers                 | {}                |
| meta.affinity                         | Affinity for Cnosdb meta replicas pods assignment                               | {}                |
| meta.nodeSelector                     | Node labels for Cnosdb meta replicas pods assignment                            | {}                |
| meta.tolerations                      | Tolerations for Cnosdb meta replicas pods assignment                            | []                |
| meta.service.type                     | Cnosdb meta replicas service type                                               | ClusterIP         |
| meta.service.port                     | Cnosdb meta replicas service port                                               | 8901              |
| meta.service.nodePort                 | Node port for Cnosdb meta replicas                                              | ""                |
| meta.service.clusterIP                | Cnosdb meta replicas service Cluster IP                                         | ""                |
| meta.service.externalTrafficPolicy    | Cnosdb meta replicas service external traffic policy                            | Cluster           |
| meta.service.annotations              | Additional custom annotations for Cnosdb meta replicas service                  | {}                |
| meta.service.loadBalancerIP           | Cnosdb meta replicas service Load Balancer IP                                   | ""                |
| meta.service.loadBalancerSourceRanges | Cnosdb meta replicas service Load Balancer sources                              | []                |
| meta.persistence.enabled              | Enable persistence on Cnosdb meta replicas nodes using Persistent Volume Claims | false             |
| meta.persistence.storageClass         | Persistent Volume storage class                                                 | ""                |
| meta.persistence.accessModes          | Persistent Volume access modes                                                  | ["ReadWriteOnce"] |
| meta.persistence.size                 | Persistent Volume size                                                          | 1Gi               |

### Tskv Parameters
**Active when architecture is `separation`**

| Name                                  | Description                                                                     | Value             |
| ------------------------------------- | ------------------------------------------------------------------------------- | ----------------- |
| tskv.replicaCount                     | Number of Cnosdb tskv replicas to deploy                                        | 2                 |
| tskv.terminationGracePeriodSeconds    | Integer setting the termination grace period for the Cnosdb tskv-replicas pods  | 10                |
| tskv.extraConf                        | Configuration for Cnosdb tskv replicas nodes                                    | {}                |
| tskv.resources.limits                 | The resources limits for the Cnosdb tskv replicas containers                    | {}                |
| tskv.resources.requests               | The requested resources for the Cnosdb tskv replicas containers                 | {}                |
| tskv.affinity                         | Affinity for Cnosdb tskv replicas pods assignment                               | {}                |
| tskv.nodeSelector                     | Node labels for Cnosdb tskv replicas pods assignment                            | {}                |
| tskv.tolerations                      | Tolerations for Cnosdb tskv replicas pods assignment                            | []                |
| tskv.service.type                     | Cnosdb tskv replicas service type                                               | ClusterIP         |
| tskv.service.ports.http               | Cnosdb tskv replicas service http port                                          | 8902              |
| tskv.service.ports.grpc               | Cnosdb tskv replicas service grpc port                                          | 8903              |
| tskv.service.ports.flight             | Cnosdb tskv replicas service flight port                                        | 8904              |
| tskv.service.ports.tcp                | Cnosdb tskv replicas service tcp port                                           | 8905              |
| tskv.service.ports.vector             | Cnosdb tskv replicas service vector port                                        | 8906              |
| tskv.service.nodePorts.http           | http Node port for Cnosdb tskv replicas                                         | ""                |
| tskv.service.nodePorts.grpc           | grpc Node port for Cnosdb tskv replicas                                         | ""                |
| tskv.service.nodePorts.flight         | flight Node port for Cnosdb tskv replicas                                       | ""                |
| tskv.service.nodePorts.tcp            | tcp Node port for Cnosdb tskv replicas                                          | ""                |
| tskv.service.nodePorts.vector         | vector Node port for Cnosdb tskv replicas                                       | ""                |
| tskv.service.clusterIP                | Cnosdb tskv replicas service Cluster IP                                         | ""                |
| tskv.service.externalTrafficPolicy    | Cnosdb tskv replicas service external traffic policy                            | Cluster           |
| tskv.service.annotations              | Additional custom annotations for Cnosdb tskv replicas service                  | {}                |
| tskv.service.loadBalancerIP           | Cnosdb tskv replicas service Load Balancer IP                                   | ""                |
| tskv.service.loadBalancerSourceRanges | Cnosdb tskv replicas service Load Balancer sources                              | []                |
| tskv.persistence.enabled              | Enable persistence on Cnosdb tskv replicas nodes using Persistent Volume Claims | false             |
| tskv.persistence.storageClass         | Persistent Volume storage class                                                 | ""                |
| tskv.persistence.accessModes          | Persistent Volume access modes                                                  | ["ReadWriteOnce"] |
| tskv.persistence.size                 | Persistent Volume size                                                          | 1Gi               |

### Query Parameters
**Active when architecture is `separation`**

| Name                                   | Description                                                      | Value     |
| -------------------------------------- | ---------------------------------------------------------------- | --------- |
| query.replicaCount                     | Number of Cnosdb query replicas to deploy                        | 2         |
| query.extraConf                        | Configuration for Cnosdb query replicas nodes                    | {}        |
| query.resources.limits                 | The resources limits for the Cnosdb query replicas containers    | {}        |
| query.resources.requests               | The requested resources for the Cnosdb query replicas containers | {}        |
| query.affinity                         | Affinity for Cnosdb query replicas pods assignment               | {}        |
| query.nodeSelector                     | Node labels for Cnosdb query replicas pods assignment            | {}        |
| query.tolerations                      | Tolerations for Cnosdb query replicas pods assignment            | []        |
| query.service.type                     | Cnosdb query replicas service type                               | ClusterIP |
| query.service.ports.http               | Cnosdb query replicas service http port                          | 8902      |
| query.service.ports.grpc               | Cnosdb query replicas service grpc port                          | 8903      |
| query.service.ports.flight             | Cnosdb query replicas service flight port                        | 8904      |
| query.service.ports.tcp                | Cnosdb query replicas service tcp port                           | 8905      |
| query.service.ports.vector             | Cnosdb query replicas service vector port                        | 8906      |
| query.service.nodePorts.http           | http Node port for Cnosdb query replicas                         | ""        |
| query.service.nodePorts.grpc           | grpc Node port for Cnosdb query replicas                         | ""        |
| query.service.nodePorts.flight         | flight Node port for Cnosdb query replicas                       | ""        |
| query.service.nodePorts.tcp            | tcp Node port for Cnosdb query replicas                          | ""        |
| query.service.nodePorts.vector         | vector Node port for Cnosdb query replicas                       | ""        |
| query.service.clusterIP                | Cnosdb query replicas service Cluster IP                         | ""        |
| query.service.externalTrafficPolicy    | Cnosdb query replicas service external traffic policy            | Cluster   |
| query.service.annotations              | Additional custom annotations for Cnosdb query replicas service  | {}        |
| query.service.loadBalancerIP           | Cnosdb query replicas service Load Balancer IP                   | ""        |
| query.service.loadBalancerSourceRanges | Cnosdb query replicas service Load Balancer sources              | []        |

### QueryTskv Parameters
**Active when architecture is `bundle`**

| Name                                       | Description                                                                           | Value             |
| ------------------------------------------ | ------------------------------------------------------------------------------------- | ----------------- |
| queryTskv.replicaCount                     | Number of Cnosdb query_tskv replicas to deploy                                        | 2                 |
| queryTskv.terminationGracePeriodSeconds    | Integer setting the termination grace period for the Cnosdb query_tskv-replicas pods  | 10                |
| queryTskv.extraConf                        | Configuration for Cnosdb query_tskv replicas nodes                                    | {}                |
| queryTskv.resources.limits                 | The resources limits for the Cnosdb query_tskv replicas containers                    | {}                |
| queryTskv.resources.requests               | The requested resources for the Cnosdb query_tskv replicas containers                 | {}                |
| queryTskv.affinity                         | Affinity for Cnosdb query_tskv replicas pods assignment                               | {}                |
| queryTskv.nodeSelector                     | Node labels for Cnosdb query_tskv replicas pods assignment                            | {}                |
| queryTskv.tolerations                      | Tolerations for Cnosdb query_tskv replicas pods assignment                            | []                |
| queryTskv.service.type                     | Cnosdb query_tskv replicas service type                                               | ClusterIP         |
| queryTskv.service.ports.http               | Cnosdb query_tskv replicas service http port                                          | 8902              |
| queryTskv.service.ports.grpc               | Cnosdb query_tskv replicas service grpc port                                          | 8903              |
| queryTskv.service.ports.flight             | Cnosdb query_tskv replicas service flight port                                        | 8904              |
| queryTskv.service.ports.tcp                | Cnosdb query_tskv replicas service tcp port                                           | 8905              |
| queryTskv.service.ports.vector             | Cnosdb query_tskv replicas service vector port                                        | 8906              |
| queryTskv.service.nodePorts.http           | http Node port for Cnosdb query_tskv replicas                                         | ""                |
| queryTskv.service.nodePorts.grpc           | grpc Node port for Cnosdb query_tskv replicas                                         | ""                |
| queryTskv.service.nodePorts.flight         | flight Node port for Cnosdb query_tskv replicas                                       | ""                |
| queryTskv.service.nodePorts.tcp            | tcp Node port for Cnosdb query_tskv replicas                                          | ""                |
| queryTskv.service.nodePorts.vector         | vector Node port for Cnosdb query_tskv replicas                                       | ""                |
| queryTskv.service.clusterIP                | Cnosdb query_tskv replicas service Cluster IP                                         | ""                |
| queryTskv.service.externalTrafficPolicy    | Cnosdb query_tskv replicas service external traffic policy                            | Cluster           |
| queryTskv.service.annotations              | Additional custom annotations for Cnosdb query_tskv replicas service                  | {}                |
| queryTskv.service.loadBalancerIP           | Cnosdb query_tskv replicas service Load Balancer IP                                   | ""                |
| queryTskv.service.loadBalancerSourceRanges | Cnosdb query_tskv replicas service Load Balancer sources                              | []                |
| queryTskv.persistence.enabled              | Enable persistence on Cnosdb query_tskv replicas nodes using Persistent Volume Claims | false             |
| queryTskv.persistence.storageClass         | Persistent Volume storage class                                                       | ""                |
| queryTskv.persistence.accessModes          | Persistent Volume access modes                                                        | ["ReadWriteOnce"] |
| queryTskv.persistence.size                 | Persistent Volume size                                                                | 1Gi               |

## Tips
`PV` will not be removed when delete a cnosdb helm release unless delete `PVC` manully.

### Persistence

If you have enabled persistence by specifying `[queryTskv | tskv | meta].persistence.enabled=true`, The chart will mounts a [Persistent Volume](https://kubernetes.io/docs/concepts/storage/persistent-volumes/) volume at the `/var/lib/cnosdb` path. The volume is created using dynamic volume provisioning, by default. 

### Extra Configuration
A toml configuration can be expressed by inline expression, `demo.foo=bar` is equal to this:
```toml
# demo.foo=bar
[demo]
foo = bar
```
Overwrite the default settings for any replicas:
```sh
helm install \
--set meta.extraConf.'storage\.maxsummary_size'='64M' \
--set tskv.extraConf.'storage\.max_level'=1 \
my-cnosdb cnosdb/cnosdb
```

### Upgrade the chart release

Upgrade Cnosdb version only.

```sh
helm upgrade my-cnosdb cnosdb/cnosdb -ncnosdb --reuse-values --set image.cnosdb.tag=new.version
```

Perform a horizontal scale.

```sh
helm upgrade my-cnosdb cnosdb/cnosdb -ncnosdb --reuse-values --set meta.replicaCount=3 --set tskv.replicaCount=5 
```

If the underlying infrastructure supports dynamically [expand PVC](https://kubernetes.io/docs/concepts/storage/persistent-volumes/#expanding-persistent-volumes-claims), you could perform a `vertical scale` for storage. However helm cannot do that, because helm cannot deal with statefulset's validation error, and statefulset cannot change pvc size right now. So you can only perform `vertical scale` for `resources` such as `cpu` and `memory`.

```sh
helm upgrade my-cnosdb cnosdb/cnosdb -ncnosdb --reuse-values --set tskv.resources.limits.cpu=1
```

Chart Information is cached locally, if you wanna upgrade chart(not app) version, you should update information of available charts locally from chart repositories.

```sh
helm repo update cnosdb
```
Then upgrade the chart as needed

```sh
helm upgrade my-cnosdb cnosdb/cnosdb -ncnosdb --set foo=bar
```