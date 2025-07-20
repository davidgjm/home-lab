
See [VM IDs](../README.md#vm-id-allocation) to see VM ID allocation convention

# Cloud-init VM Templates

## Create template

See details at https://pve.proxmox.com/pve-docs/pve-admin-guide.html#_preparing_cloud_init_templates



### Reference snippets

```shell
# download the image
wget https://cloud-images.ubuntu.com/bionic/current/bionic-server-cloudimg-amd64.img

# create a new VM with VirtIO SCSI controller
qm create 9000 --memory 2048 --net0 virtio,bridge=vmbr0 --scsihw virtio-scsi-pci

# import the downloaded disk to the local-lvm storage, attaching it as a SCSI drive
qm set 9000 --scsi0 local-lvm:0,import-from=/path/to/bionic-server-cloudimg-amd64.img
```

## Create VM without disk

```shell
vmid=100001000
qm create $vmid --cores 2 --cpu x86-64-v2-AES --memory 2048 --net0 virtio,bridge=vmbr0 --scsihw virtio-scsi-pci
```



## Import from cloud images (cloud-init enabled)


There're 2 VM disk options to import the import the images:
1. Use local disk: `local-lvm`
2. Use ceph RDB: `vms` (ceph RDB storage pool)
![](PVE%20storage%20pools.png)


### Debian
Debian cloud images are in `qcow2` format.


#### Image location

```shell
# Debian 12
image=/mnt/pve/cephfs/template/qcow/debian-12-genericcloud-amd64.qcow2
```

#### Import
#### Debian cloud init image
```shell
qm set $vmid --scsi0 local-lvm:0,import-from=$image
```

OR
```shell
qm importdisk $vmid $image local --format qcow2 --target-disk scsi0
```




### Ubuntu

The ubuntu cloud images are stored in cephfs at `/mnt/pve/cephfs/template/iso/`.
#### Image location

| Version | Path                                                                 |
| ------- | -------------------------------------------------------------------- |
| 24.04   | `image=/mnt/pve/cephfs/template/iso/noble-server-cloudimg-amd64.img` |
| 22.04   | `image=/mnt/pve/cephfs/template/iso/jammy-server-cloudimg-amd64.img` |

#### Import into local lvm

```shell
qm set $vmid --scsi0 local-lvm:0,import-from=$image
```

OR
#### Import into ceph RDB
```shell
qm set $vmid --scsi0 vms:0,import-from=$image
```


## Template configuration

```shell
qm set $vmid --ostype l26
qm set $vmid --agent enabled=1
qm set $vmid --ciupgrade false
qm set $vmid --ipconfig0 ip=dhcp
```

### Custom user config file


#### Debian
```shell
qm set $vmid --cicustom "user=cephfs:snippets/debian-userconfig.yaml"
```


### Ubuntu
```shell
qm set $vmid --cicustom "user=cephfs:snippets/ubuntu-userconfig.yaml"
```



### Add cloud-init CD-ROM drive
```shell
qm set $vmid --ide2 local-lvm:cloudinit
qm set $vmid --boot order=scsi0
```


## Convert to VM template
```shell
qm template $vmid
```

## Clone VM from template
```shell
root@m5:~# qm list
      vmid NAME                 STATUS     MEM(MB)    BOOTDISK(GB) PID
       100 vm1                  running    2048              22.20 80720
      9000 VM 9000              stopped    2048              22.20 0
```

### Clone VM from the template

```shell
root@m5:~# qm clone $vmid 7000 --name citest001
create full clone of drive ide2 (local-lvm:vm-9000-cloudinit)
  Logical volume "vm-500-cloudinit" created.
create linked clone of drive scsi0 (local-lvm:base-9000-disk-0)
  Logical volume "vm-500-disk-0" created.
```

## Further configuration

### Resize disk
```shell
qm disk resize 3100 scsi0 +1G
```