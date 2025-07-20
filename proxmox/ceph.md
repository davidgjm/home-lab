
# [Hyper-Converged Ceph Clustering](https://pve.proxmox.com/pve-docs/pve-admin-guide.html#chapter_pveceph)


## Joining cluster for new PVE nodes

### General steps
1. After joining the PVE cluster, install ceph your new node from web console. Key options:
	1. Squid 19.2.1
	2. No-subscription
2. Configure Ceph for the new node
3. Get bootstrap auth from other existing node
	1. `scp m7:/var/lib/ceph/bootstrap-osd/ceph.keyring /var/lib/ceph/bootstrap-osd/ceph.keyring`
4. Add Ceph components as needed
	1. OSD
	2. Monitor
	3. Managers


### Configure ceph
 Link the cluster ceph configuration to the local configuration

```bash
ln -s /etc/pve/ceph.conf /etc/ceph/ceph.conf
```

### bootstrap-osd auth

Get bootstrap auth from other existing node
```shell
scp m7:/var/lib/ceph/bootstrap-osd/ceph.keyring /var/lib/ceph/bootstrap-osd/ceph.keyring
```

### Add other components
```shell
pveceph mon create --mon-address 172.16.100.39

```


# Ceph HOW-TO


## Create LVM volumes to add as OSDs

### Create PV & VG from an entire disk (optional)

```shell
pvcreate /dev/nvme1n1
vgcreate data /dev/nvme1n1
```

### Create data logical volumes for the data disk

```shell
lvcreate -L 500G -n data/lvpool-01
lvcreate -L 500G -n data/lvpool-02
lvcreate -L 500G -n data/lvpool-03
lvcreate -l 100%FREE -n data/lvpool-04
```


## Add LVM volumes as OSDs for ceph cluster

### create ceph OSDs from LV

Use the logical volumes created previously to create ceph volumes.

```shell
ceph-volume lvm create --data data/lvpool-01
ceph-volume lvm create --data data/lvpool-02
ceph-volume lvm create --data data/lvpool-03
ceph-volume lvm create --data data/lvpool-04
```


### Reference Example
#### Add lvm logical volume as OSD
* https://forum.proxmox.com/threads/ceph-osd-on-lvm-logical-volume.68618/

1. During install set maxvz to 0 to not create local storage and keep free space for Ceph on the OS drive. [GUIDE, 2.3.1 Advanced LVM Configuration Options ]  
2. Setup Proxmox like usual and create a cluster  
3. Install Ceph packages and do initial setup (network interfaces etc.) via GUI, also create Managers and Monitors  
4. To create OSDs open a shell on each node and
	1. bootstrap auth [4]:  
		`ceph auth get client.bootstrap-osd > /var/lib/ceph/bootstrap-osd/ceph.keyring`
	2. Create new logical volume with the remaining free space:  
		`lvcreate -l 100%FREE -n pve/vz`
	3. Create (= prepare and activate) the logical volume for OSD [2] [3]  
		`ceph-volume lvm create --data pve/vz`
5. That's it. Now you can keep using GUI to:  
	- create Metadata servers,
	- by clicking on a node in the cluster in "Ceph" create a CephFS. And then add in "Datacenter-Storage" for CD images and backups. This will be mounted in /mnt/pve/cephfs/
	- and in "Datacenter-Storage" add an "RDS" block device for virtual VM HDDs.



[GUIDE] [https://pve.proxmox.com/pve-docs/pve-admin-guide.pdf](https://pve.proxmox.com/pve-docs/pve-admin-guide.pdf)  
[2] [https://docs.ceph.com/docs/master/ceph-volume/lvm/create/#ceph-volume-lvm-create](https://docs.ceph.com/docs/master/ceph-volume/lvm/create/#ceph-volume-lvm-create)  
[3] [https://docs.ceph.com/docs/master/ceph-volume/](https://docs.ceph.com/docs/master/ceph-volume/)  
[4] [https://forum.proxmox.com/threads/p...ble-to-create-a-new-osd-id.55730/#post-257533](https://forum.proxmox.com/threads/proxmox-6-0-unable-to-create-ceph-osds-unable-to-create-a-new-osd-id.55730/#post-257533)



