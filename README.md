# CnosDB
CnosDB is a high-performance, high-compression, and easy-to-use open-source distributed time-series database. It is primarily used in fields such as IoT, industrial internet, connected cars, and IT operations. All of the code is open-sourced and available on GitHub.

English | [简体中文](./README_CN.md)
## tl;dr
```sh
helm repo add cnosdb https://cnosdb.github.io/helm-chart/
helm repo update
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
| image.cnosdb.repository          | Cnosdb image repository              | cnosdb/cnosdb            |
| image.cnosdb.tag                 | Cnosdb image tag                     | community-latest         |
| image.clusterConfTool.repository | Cnosdb cluster tool image repository | wyxok/cnosdb-init-config |
| image.clusterConfTool.tag        | Cnosdb cluster tool image tag        | latest                   |
| image.pullPolicy                 | Image pull policy                    | IfNotPresent             |
| image.pullSecrets                | Image pull secrets                   | []                       |


### Common Parameters

| Name             | Description                                        | Value   |
| ---------------- | -------------------------------------------------- | ------- |
| nameOverride     | String to partially override common.names.fullname | ""      |
| fullnameOverride | String to fully override common.names.fullname     | ""      |
| architecture     | The architecture of installing                     | cluster |

### Meta Parameters

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
| meta.persistence.existingClaim        | Use a existing PVC which must be created manually before bound                  | ""                |

### Tskv Parameters

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
| tskv.service.nodePorts.http           | http Node port for Cnosdb tskv replicas                                         | ""                |
| tskv.service.nodePorts.grpc           | grpc Node port for Cnosdb tskv replicas                                         | ""                |
| tskv.service.nodePorts.flight         | flight Node port for Cnosdb tskv replicas                                       | ""                |
| tskv.service.nodePorts.tcp            | tcp Node port for Cnosdb tskv replicas                                          | ""                |
| tskv.service.clusterIP                | Cnosdb tskv replicas service Cluster IP                                         | ""                |
| tskv.service.externalTrafficPolicy    | Cnosdb tskv replicas service external traffic policy                            | Cluster           |
| tskv.service.annotations              | Additional custom annotations for Cnosdb tskv replicas service                  | {}                |
| tskv.service.loadBalancerIP           | Cnosdb tskv replicas service Load Balancer IP                                   | ""                |
| tskv.service.loadBalancerSourceRanges | Cnosdb tskv replicas service Load Balancer sources                              | []                |
| tskv.persistence.enabled              | Enable persistence on Cnosdb tskv replicas nodes using Persistent Volume Claims | false             |
| tskv.persistence.storageClass         | Persistent Volume storage class                                                 | ""                |
| tskv.persistence.accessModes          | Persistent Volume access modes                                                  | ["ReadWriteOnce"] |
| tskv.persistence.size                 | Persistent Volume size                                                          | 1Gi               |
| tskv.persistence.existingClaim        | Use a existing PVC which must be created manually before bound                  | ""                |

### Query Parameters

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
| query.service.ports.flight             | Cnosdb query replicas service flight port                        | 8904      |
| query.service.nodePorts.http           | http Node port for Cnosdb query replicas                         | ""        |
| query.service.nodePorts.flight         | flight Node port for Cnosdb query replicas                       | ""        |
| query.service.clusterIP                | Cnosdb query replicas service Cluster IP                         | ""        |
| query.service.externalTrafficPolicy    | Cnosdb query replicas service external traffic policy            | Cluster   |
| query.service.annotations              | Additional custom annotations for Cnosdb query replicas service  | {}        |
| query.service.loadBalancerIP           | Cnosdb query replicas service Load Balancer IP                   | ""        |
| query.service.loadBalancerSourceRanges | Cnosdb query replicas service Load Balancer sources              | []        |


## Tips
PV will not be removed when delete a cnosdb helm release

### Persistence

The chart mounts a [Persistent Volume](https://kubernetes.io/docs/concepts/storage/persistent-volumes/) volume at the `/var/lib/cnosdb` path. The volume is created using dynamic volume provisioning, by default. 
An existing PersistentVolumeClaim can be defined. If a Persistent Volume Claim already exists, specify it during installation.

### Existing PersistentVolumeClaims

1. Create the PersistentVolume
1. Create the PersistentVolumeClaim
1. Install the chart
```sh
helm install --set meta.persistence.existingClaim=PVC_NAME my-cnosdb cnosdb/cnosdb
```

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