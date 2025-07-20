
# VM ID Allocation

The max length of PVE VM ID is **9** digits.

## Convention

| Min       | Max       | Description                                     |
| --------- | --------- | ----------------------------------------------- |
| 10000-    |           | VM IDs starting with 10000 are for VM templates |
| 100001000 | 100001999 | ubuntu templates                                |
| 100002000 | 100002999 | debian templates                                |
| 100003000 | 100003999 | coreos                                          |
| 100005000 | 100005999 | other - placeholder                             |
| 1000      | 1999      | VM IDs for routerhost.home.lab                  |
| 3000      | 3999      | VM IDs for m5.home.lab                          |
| 5000      | 5999      | VM IDs for m7.home.lab                          |
| 7000      | 7999      | VM IDs for ser6.home.lab                        |
| 9000      | 9999      | VM IDs for gem12.home.lab                       |



# Key Concepts

## Storage

- https://pve.proxmox.com/pve-docs/pve-admin-guide.html#_storage_types
#### 7.2.2. Common Storage Properties [](https://pve.proxmox.com/pve-docs/pve-admin-guide.html#_common_storage_properties "Permalink to this heading")

A few storage properties are common among different storage types.

nodes

List of cluster node names where this storage is usable/accessible. One can use this property to restrict storage access to a limited set of nodes.

`content`

A storage can support several content types, for example virtual disk images, cdrom iso images, container templates or container root directories. Not all storage types support all content types. One can set this property to select what this storage is used for.

images

QEMU/KVM VM images.

rootdir

Allow to store container data.

vztmpl

Container templates.

backup

Backup files (vzdump).

iso

ISO images

snippets

Snippet files, for example guest hook scripts

shared

Indicate that this is a single storage with the same contents on all nodes (or all listed in the _nodes_ option). It will not make the contents of a local storage automatically accessible to other nodes, it just marks an already shared storage as such!

disable

You can use this flag to disable the storage completely.

maxfiles

Deprecated, please use prune-backups instead. Maximum number of backup files per VM. Use 0 for unlimited.

prune-backups

Retention options for backups. For details, see [Backup Retention](https://pve.proxmox.com/pve-docs/pve-admin-guide.html#vzdump_retention).

format

Default image format (raw|qcow2|vmdk)

preallocation

Preallocation mode (off|metadata|falloc|full) for raw and qcow2 images on file-based storages. The default is metadata, which is treated like off for raw images. When using network storages in combination with large qcow2 images, using off can help to avoid timeouts.

|   |   |
|---|---|
|![Warning](data:image/png;base64,%0A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%0A/F/lgJiyQFHragAAAABJRU5ErkJggg==)|It is not advisable to use the same storage pool on different Proxmox VE clusters. Some storage operation need exclusive access to the storage, so proper locking is required. While this is implemented within a cluster, it does not work between different clusters.|