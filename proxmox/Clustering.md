
# corosync configuration
## Separate network for corosync

Follow the guides from proxmox documentation:
- https://pve.proxmox.com/wiki/Cluster_Manager#pvecm_cluster_network
- https://pve.proxmox.com/wiki/Cluster_Manager#_corosync_configuration

## Update `corosync.conf`

>There are two on each cluster node, one in `/etc/pve/corosync.conf` and the other in `/etc/corosync/corosync.conf`. Editing the one in our cluster file system will propagate the changes to the local one, but not vice versa.

> [!IMPORTANT]
> Update `/etc/pve/corosync.conf` instead of local `/etc/corosync/corosync.conf`

https://pve.proxmox.com/pve-docs/pve-admin-guide.html#pvecm_edit_corosync_conf

# Add a new PVE Node

> [!TIP]
> For PVE hosts with multiple NICs, it's more reliable to join the cluster than through Web UI.

## Joining PVE cluster

Take the following as an example, the new node has 2 network interfaces which has the following IPs:

| interface | IP            | Usage                              |
| --------- | ------------- | ---------------------------------- |
| eno1      | 172.17.100.39 | primary interface for cluster data |
| enp3s0    | 172.17.100.40 | management interface. For corosync |


```shell
pvecm add m7.home.lab --force --link0=172.16.100.40
```


## Update SDN network settings

In the Web UI, do the following:
1. Go to **Datacenter -> SDN**
2. You will see pending changes in the new node
3. Click on *Apply*


## ceph storage setup

See [ceph.md](ceph)