# Preparation - HA cluster with kubeadm

## Node Information

| Role                    |              DNS Name | IP Address     |
| ----------------------- | --------------------: | -------------- |
| jump server             | jumpserver01.home.lab | 172.16.100.200 |
| Load Balancer (haproxy) |          klb.home.lab | 172.16.100.21  |
| Controller 1            |       adm-c0.home.lab | 172.16.100.217 |
| Controller 2            |       adm-c1.home.lab | 172.16.100.218 |
| Controller 3            |       adm-c2.home.lab | 172.16.100.219 |
| Node 1                  |       adm-n0.home.lab | 172.16.100.230 |
| Node 2                  |       adm-n1.home.lab | 172.16.100.231 |
| Node 3                  |       adm-n2.home.lab | 172.16.100.232 |

## Install `kubeadm`

Following the guide at https://v1-32.docs.kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/


# Control Plane

## bootstrap controller

```bash
sudo kubeadm init --control-plane-endpoint "klb.home.lab:6443" --pod-network-cidr 10.244.0.0/16 --upload-certs --kubernetes-version v1.32.5
```


### Cluster joining information

```text
You can now join any number of control-plane nodes running the following command on each as root:

  kubeadm join klb.home.lab:6443 --token 88civd.44m90wi1ocva59q5 \
	--discovery-token-ca-cert-hash sha256:b2b0da15a01b4ad1e0fa168e610c0fb5f798d99ad897b7ecca71ae5cd295188d \
	--control-plane --certificate-key a9e679f31c3aa1517730c5858fce672041a8f4921b04eb75cf93da46a3aad863

Please note that the certificate-key gives access to cluster sensitive data, keep it secret!
As a safeguard, uploaded-certs will be deleted in two hours; If necessary, you can use
"kubeadm init phase upload-certs --upload-certs" to reload certs afterward.

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join klb.home.lab:6443 --token 88civd.44m90wi1ocva59q5 \
	--discovery-token-ca-cert-hash sha256:b2b0da15a01b4ad1e0fa168e610c0fb5f798d99ad897b7ecca71ae5cd295188d
```


### kubeconfig for cluster admin

```text
To start administering your cluster from this node, you need to run the following as a regular user:

	mkdir -p $HOME/.kube
	sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
	sudo chown $(id -u):$(id -g) $HOME/.kube/config
```


# Cluster Configuration
## Kubernetes cluster metadata

| Item                           | Value           | Configuration Source                                   | Configuration Information                                                                                                 |
| ------------------------------ | --------------- | ------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------- |
| Service cluster IP CIDR        | `10.96.0.0/12`  | /etc/kubernetes/manifests/kube-apiserver.yaml          | --service-cluster-ip-range=10.96.0.0/12                                                                                   |
| Pod IP CIDR                    | `10.244.0.0/16` | /etc/kubernetes/manifests/kube-controller-manager.yaml | `--cluster-cidr=10.200.0.0/16`                                                                                            |
| internal kubernetes service IP | `10.96.0.1/24`  | kubernetes service object yaml                         | `clusterIP: 10.96.0.1`                                                                                                    |
| Cluster DNS Server IP          | 10.96.0.10/24`  | coreDNS set by kubeadm                                 | [defaults.go](https://github.com/kubernetes/kubernetes/blob/v1.32.5/cmd/kubeadm/app/apis/kubeadm/v1beta4/defaults.go#L37) |
| DNS Domain                     | `cluster.local` | kubelet-config.yaml                                    | [defaults.go](https://github.com/kubernetes/kubernetes/blob/v1.32.5/cmd/kubeadm/app/apis/kubeadm/v1beta4/defaults.go#L37) |




# Troubleshooting

## flannel add-on

###  Error registering network: failed to acquire lease: node "adm-n0" pod cidr not assigned

flannel reports the following error when `--pod-network-cidr` argument is missing when bootstrapping the cluster.

``
```text
I0602 11:03:13.955718       1 vxlan.go:141] VXLAN config: VNI=1 Port=0 GBP=false Learning=false DirectRouting=false
I0602 11:03:13.956989       1 kube.go:636] List of node(adm-n0) annotations: map[string]string{"kubeadm.alpha.kubernetes.io/cri-socket":"unix:///var/run/containerd/containerd.sock", "node.alpha.kubernetes.io/ttl":"0", "volumes.kubernetes.io/controller-managed-attach-detach":"true"}
E0602 11:03:13.957123       1 main.go:359] Error registering network: failed to acquire lease: node "adm-n0" pod cidr not assigned
I0602 11:03:13.957165       1 main.go:448] Stopping shutdownHandler...
```

To fix this, need to manually update the `kube-controller-manager` static pod manifest. If the cluster is HA with multiple controller nodes, all nodes need to be updated.

Add the following args:
* `--allocate-node-cidrs=true`
* `--cluster-cidr=10.244.0.0/16`

```yaml
spec:
  containers:
  - command:
    - kube-controller-manager
    - --allocate-node-cidrs=true
    - --cluster-cidr=10.244.0.0/16
```


## Resetting endpoints for master service "kubernetes" to ...

## `kubernetes.default.svc` endpoints reset
With the HA cluster bootstrapped with kubeadm (stacked etcd mode), the default IP for `kubernetes.default.svc` is `10.96.0.1`. However, the service endpoints are reset and changed to node IPs when deploying metalLB IpAddressPool. 

Steps to reproduce:
1. On a newly bootstrapped cluster, deploy **MetalLB** from `kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.14.9/config/manifests/metallb-native.yaml`
2. Deploy `IPAddressPool` to see the resource can be created successfully.
3. The logs from `kube-apiserver` will contain information about resetting the endpoints.


```text
I0602 10:18:45.172384       1 policy_source.go:240] refreshing policies
I0602 10:18:45.175622       1 handler_discovery.go:451] Starting ResourceDiscoveryManager
I0602 10:18:45.176851       1 cache.go:39] Caches are synced for LocalAvailability controller
I0602 10:18:45.180983       1 apf_controller.go:382] Running API Priority and Fairness config worker
I0602 10:18:45.181019       1 apf_controller.go:385] Running API Priority and Fairness periodic rebalancing process
I0602 10:18:45.181052       1 shared_informer.go:320] Caches are synced for cluster_authentication_trust_controller
I0602 10:18:45.186342       1 shared_informer.go:320] Caches are synced for crd-autoregister
I0602 10:18:45.186415       1 aggregator.go:171] initial CRD sync complete...
I0602 10:18:45.186436       1 autoregister_controller.go:144] Starting autoregister controller
I0602 10:18:45.186449       1 cache.go:32] Waiting for caches to sync for autoregister controller
I0602 10:18:45.186479       1 cache.go:39] Caches are synced for autoregister controller
I0602 10:18:45.221300       1 controller.go:615] quota admission added evaluator for: leases.coordination.k8s.io
I0602 10:18:46.079209       1 storage_scheduling.go:111] all system priority classes are created successfully or already exist.
I0602 10:18:46.175021       1 shared_informer.go:320] Caches are synced for configmaps
W0602 10:18:46.290685       1 lease.go:265] Resetting endpoints for master service "kubernetes" to [172.16.100.217 172.16.100.218 172.16.100.219]
I0602 10:18:46.291945       1 controller.go:615] quota admission added evaluator for: endpoints
I0602 10:18:46.296827       1 controller.go:615] quota admission added evaluator for: endpointslices.discovery.k8s.io
```

### IpAddressPool

```yaml
apiVersion: metallb.io/v1beta1
kind: IPAddressPool
metadata:
  name: ip-pool-basic
  namespace: metallb-system
spec:
  addresses:
  - 172.16.100.11-172.16.100.29
  avoidBuggyIPs: true
```

----



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