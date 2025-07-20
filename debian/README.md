# Installation

## Prepare Installation Image from USB device on MacOs

1. Download the DVD image from primary mirror repository at https://cdimage.debian.org/debian-cd/
2. From MacOS terminal, run command like below

```shell
cd ~/Downloads
sudo dd if=debian-12.10.0-amd64-DVD-1.iso bs=4M of=/dev/disk4 status=progress oflag=sync
```


# List files installed by a package

```shell
dpkg -L <package>
```




#debian

