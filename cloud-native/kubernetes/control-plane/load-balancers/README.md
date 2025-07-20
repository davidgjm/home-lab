
# Bare metal deployment issues

## validation hook url not reachable from kube-api server

> Error from server (InternalError): error when creating "config.yaml": Internal error occurred: failed calling webhook "ipaddresspoolvalidationwebhook.metallb.io": failed to call webhook: Post "https://metallb-webhook-service.metallb-system.svc:443/validate-metallb-io-v1beta1-ipaddresspool?timeout=10s": context deadline exceeded

References:
* https://github.com/metallb/metallb/issues/1547
* https://github.com/metallb/metallb/issues/1597
* https://github.com/metallb/metallb/issues/1597#issuecomment-2783904978
* [Admission webhooks and API-server (controller) networking](https://github.com/kelseyhightower/kubernetes-the-hard-way/issues/588)
* [Build a managed Kubernetes cluster from scratch — part 5](https://medium.com/@norlin.t/build-a-managed-kubernetes-cluster-from-scratch-part-5-a4c22f0c0245)

> I encountered a similar issue and found it was related to missing routing for the Kubernetes service IP range.
> 
> In my setup, I followed [Kubernetes the Hard Way](https://github.com/kelseyhightower/kubernetes-the-hard-way), and the root cause turned out to be the service CIDR not being properly routed. Specifically, traffic to the webhook service IP was not reaching the controller pod, which caused the InternalError (failed calling webhook "ipaddresspoolvalidationwebhook.metallb.io").
>  
> The fix was to add a route for the service IP range (--service-cluster-ip-range) to the node network, as described in this comment:  
[kelseyhightower/kubernetes-the-hard-way#588 (comment)](https://github.com/kelseyhightower/kubernetes-the-hard-way/issues/588#issuecomment-2783825235)
> 
> Once the route was in place, the apiserver was able to reach the webhook pod, and the error was resolved.
> 
> Might be worth checking if the service CIDR is correctly routed in your environment, especially if you're using a custom or manually configured cluster.

![](Pasted%20image%2020250609215622.png)