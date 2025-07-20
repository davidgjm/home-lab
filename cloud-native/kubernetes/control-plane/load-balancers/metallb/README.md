# metallb installation & configuration

## Installation

### Create namespace

```bash
kubectl create namespace metallb-system
```

### Helm installation
```bash
helm repo add metallb https://metallb.github.io/metallb
helm install metallb metallb/metallb
```

