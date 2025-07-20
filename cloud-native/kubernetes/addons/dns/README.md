
# Deployment through helm chart


## Add helm repo

```shell
helm repo add coredns https://coredns.github.io/helm
```


## Install coredns


```shell
helm --namespace=kube-system install coredns coredns/coredns \
--set service.clusterIP=10.32.0.10
```


# Deployment from `kubeadm` template

The `kubeadm` coreDNS templates are defined at https://github.com/kubernetes/kubernetes/blob/master/cmd/kubeadm/app/phases/addons/dns/manifests.go.



# Debug coredns

#### Turn on logging

Add `log` in the corefile from coredns configmap. See details at [Are DNS queries being received/processed?](https://kubernetes.io/docs/tasks/administer-cluster/dns-debugging-resolution/#are-dns-queries-being-received-processed)

```yaml
apiVersion: v1
data:
  Corefile: |-
    .:53 {
        log
        errors
        health {
            lameduck 10s
        }
        ready
        kubernetes cluster.local in-addr.arpa ip6.arpa {
            pods insecure
            fallthrough in-addr.arpa ip6.arpa
            ttl 30
        }
        prometheus 0.0.0.0:9153
        forward . /etc/resolv.conf
        cache 30
        loop
        reload
        loadbalance
    }
kind: ConfigMap
metadata:
  annotations:
    meta.helm.sh/release-name: coredns
    meta.helm.sh/release-namespace: kube-system
  creationTimestamp: "2025-05-25T13:55:16Z"
  labels:
    app.kubernetes.io/instance: coredns
    app.kubernetes.io/managed-by: Helm
    app.kubernetes.io/name: coredns
    helm.sh/chart: coredns-1.42.1
    k8s-app: coredns
    kubernetes.io/cluster-service: "true"
    kubernetes.io/name: CoreDNS
  name: coredns
  namespace: kube-system
  resourceVersion: "188743"
  uid: 13d6976e-aa76-4625-8295-0b54b472d7e3

```