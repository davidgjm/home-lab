
# Installation

## apt installation

```shell
apt install -y incus incus-base incus-ui-canonical
```

# Cheatsheet

## admi

| Command                                                    | Description                     |
| ---------------------------------------------------------- | ------------------------------- |
| `incus monitor --pretty`                                   | Realtime monitoring             |
| `incus admin sql global "select * from nodes"`             | Get node list in a cluster node |
| `openssl x509 -noout -text -in /var/lib/incus/cluster.crt` | Review cluster TLS certificate  |


## Images

|Operation|Command|Note|
|----|----|----|
|List *Debian* 12 VM images|`incus image list images: debian/12 architecture=amd64 type=virtual-machine`|