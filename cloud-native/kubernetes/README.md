
# Kubernetes learning key points & labs


| Item                   | Area                   | Note                                                                              | Status |
| ---------------------- | ---------------------- | --------------------------------------------------------------------------------- | ------ |
| TLS bootstrapping      | Cluster administration | OK for kubelet, serving certificate                                               | ✅      |
| Gateway API            | Add-on                 |                                                                                   |        |
| Calico                 | CNI add-on             |                                                                                   |        |
| Cilium                 | CNI add-on             |                                                                                   |        |
| Istio                  | Service Mesh           |                                                                                   |        |
| Envoy                  | networking/proxy       |                                                                                   |        |
| Device Plugin          | Extending Cluster      |                                                                                   |        |
| Storage solution       | CSI (Storage)          |                                                                                   |        |
| External load balancer | networking             | OpenELB/MetalLB<br>MetaLB is preferred as no change is necessary on service side. | ✅      |




# Kubernetes installation on PVE VMs
## Preparation

## Node Information

### HA cluster nodes

| Role                    |              DNS Name | IP Address     |
| ----------------------- | --------------------: | -------------- |
| jump server             | jumpserver01.home.lab | 172.16.100.200 |
| Load Balancer (haproxy) |          k8s.home.lab | 172.16.100.21  |
| Controller 1            | controller-0.home.lab | 172.16.100.210 |
| Controller 2            | controller-1.home.lab | 172.16.100.211 |
| Controller 3            | controller-2.home.lab | 172.16.100.212 |
| Node 1                  |       node-0.home.lab | 172.16.100.220 |
| Node 2                  |       node-1.home.lab | 172.16.100.221 |
| Node 3                  |       node-2.home.lab | 172.16.100.222 |

### Single Controller Cluster


| Role        |        DNS Name | IP Address     | Description |
| ----------- | --------------: | -------------- | ----------- |
| jump server |             N/A | 172.16.100.201 |             |
| Server      | server.home.lab | 172.16.100.215 |             |
| Node 0      | node-0.home.lab | 172.16.100.225 |             |
| Node 1      | node-1.home.lab | 172.16.100.226 |             |

## Kubernetes cluster metadata


| Item                           | Value           | Configuration Source                | Configuration Information               |
| ------------------------------ | --------------- | ----------------------------------- | --------------------------------------- |
| Service cluster IP CIDR        | `10.32.0.0/24`  | `kube-apiserver.service`            | --service-cluster-ip-range=10.32.0.0/24 |
| Pod IP CIDR                    | `10.200.0.0/16` | `kube-controller-manager`           | `--cluster-cidr=10.200.0.0/16`          |
| internal kubernetes service IP | `10.32.0.1/24`  | kubernetes service object yaml      | `clusterIP: 10.32.0.1`                  |
| Cluster DNS Server IP          | `10.32.0.10/24` | coreDNS helm installation parameter |                                         |
| DNS Domain                     | `cluster.local` | kubelet-config.yaml                 | `clusterDomain: "cluster.local"`        |


### Internal kubernetes service

```yaml
apiVersion: v1
kind: Service
metadata:
  labels:
    component: apiserver
    provider: kubernetes
  name: kubernetes
  namespace: default

spec:
  clusterIP: 10.32.0.1
  internalTrafficPolicy: Cluster
  ipFamilies:
  - IPv4
  ipFamilyPolicy: SingleStack
  ports:
  - name: https
    port: 443
    protocol: TCP
    targetPort: 6443
  sessionAffinity: None
  type: ClusterIP

```


# Working with nodes running on PVE

## Starting/stopping nodes through PVE `qm` command

```bash
./k8s-pve-nodes-op.sh --action=start
```

Possible values for `--action`:
* `start`
* `stop`
* `shutdown`

# Troubleshooting

## Unable to delete namespace (`terminating`)

### dump & edit the namespace file

```bash
NAMESPACE=kube-flannel
kubectl get namespace ${NAMESPACE} -o json > tmp.json
```

edit `tmp.json` to make the finalizers array empty
```bash
vim tmp.json
```


```json
"spec": {
	"finalizers": [
	]
}
```
### proxy

```bash
kubectl proxy
```
### finalize
```bash
curl -k -H "Content-Type: application/json" -X PUT --data-binary @tmp.json http://127.0.0.1:8001/api/v1/namespaces/${NAMESPACE}/finalize
```