
# bootstrapping

## Open up 8443
When the boostrap server is initialized, run the following so that the remote members will be able to join the cluster

```shell
incus config set core.https_address :8443
```

# Joining a cluster

When joining cluster is successful, remember to run the following so that the heartbeating work as expected:

```shell
incus config set core.https_address :8443
```