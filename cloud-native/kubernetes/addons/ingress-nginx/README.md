
# Installation

```shell
helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --set controller.progressDeadlineSeconds=300 \
  --namespace ingress-nginx --create-namespace

```