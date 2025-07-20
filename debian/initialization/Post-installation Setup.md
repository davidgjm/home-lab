> [!INFO]
> [Disable root login](disable-root-login) should be followed if you didn't disable root login during OS installation.

# Bring-up second network interface

## Manually add the interface `enp1s0`

Add the following text for the new interface `enp1s0`
```
# The secondary network interfac -- manually addede
allow-hotplug enp1s0
iface enp1s0 inet dhcp
```
### The resulting `/etc/network/interfaces`

```
# This file describes the network interfaces available on your system
# and how to activate them. For more information, see interfaces(5).

source /etc/network/interfaces.d/*

# The loopback network interface
auto lo
iface lo inet loopback

# The primary network interface
allow-hotplug eno1
iface eno1 inet dhcp

# The secondary network interfac -- manually addede
allow-hotplug enp1s0
iface enp1s0 inet dhcp
```


- investigate linux `netplan` to work with multiple NICs
# Setting up required tools
>[!NOTE] 
> [`ansible` automation](ansible/README.md) already contains `ansible` based automation to install these packages

## `vim` Installation
1. apt install -y vim
2. change to vim by using sudo update-alternatives --config editor

## Install `sudo`
1. `apt install -y sudo
2. `echo "export PATH=\$PATH:/usr/sbin" >> ~/.bashrc
3. `source ~/.bashrcc

### Add users to sudo

1. `sudo usermod -aG sudo david
2. `visudo
	 - Change `%sudo        ALL=(ALL:ALL) ALL` to the following
	 - `%sudo        ALL=(ALL:ALL) NOPASSWD: ALL`


# Installing through shell
```embed-shell
PATH: "vault://debian/install-packages.sh"

TITLE: "Install apt packages"
```


#apt 