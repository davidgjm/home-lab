
# Corosync

## `/etc/pve/corosync.conf` read-only leads to "Operation not permitted"

https://forum.proxmox.com/threads/corosync-not-running-file-etc-pve-corosync-conf-impossible-write.63360/

### Root cause
The cluster member/node is in an invalid status. The configuration becomes read-only because of this.

### Solution
Need to set the node to maintenance mode by running:
```shell
pvecm expected 1
```

