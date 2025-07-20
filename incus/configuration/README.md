# Incus Configuration

## Setup bridge to connect physical network interface

### Requirement

The managed network bridge should be set up in a way that all VMs connected to this bridge can get IP addresses from the network where the physical network interfaces are connected.

### Solution

1. Shutdown the target physical interface `enp2s0` and delete assigned ipv4 address
2. Create a bridge `incusbr1` from incus
3. Key configurations for the bridge
  * set `bridge.external_interfaces=enp2s0`
  * dns should be completely turned off. No ip should be assigned to the bridge. See incus code snippet below.
  * assign dhcp gateway to the value of the physical network DHCP server


https://github.com/lxc/incus/blob/main/internal/server/network/driver_bridge.go#L2685-L2688


### Configure physical network interface `enp2s0`

```bash
sudo ifdown enp2s0
sudo ip addr del 172.16.100.34/24 dev enp2s0
```

### Configure target network bridge

```bash
incus network create incusbr1
incus network set incusbr1 ipv6.address=none
incus network set incusbr1 ipv4.nat=false
incus network set incusbr1 ipv4.firewall=false ipv6.firewall=false
incus network set incusbr1 ipv4.dhcp.gateway=172.16.100.1
incus network set incusbr1 bridge.external_interfaces=enp2s0
incus network set incusbr1 dns.mode=none

```

### End state

##### IP configuration
```shell
david@vmhost1:~$ ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
2: enp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether aa:09:cc:de:ce:e6 brd ff:ff:ff:ff:ff:ff
    inet 172.16.100.33/24 brd 172.16.100.255 scope global dynamic enp1s0
       valid_lft 415023523sec preferred_lft 415023523sec
3: enp2s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel master incusbr1 state UP group default qlen 1000
    link/ether 2a:9e:83:34:85:9f brd ff:ff:ff:ff:ff:ff
    inet 172.16.100.34/24 brd 172.16.100.255 scope global dynamic enp2s0
       valid_lft 415023524sec preferred_lft 415023524sec
4: wlp3s0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default qlen 1000
    link/ether 94:8c:d7:c0:f0:56 brd ff:ff:ff:ff:ff:ff
5: incusbr0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default qlen 1000
    link/ether 00:16:3e:ec:1e:f2 brd ff:ff:ff:ff:ff:ff
    inet 10.59.219.1/24 scope global incusbr0
       valid_lft forever preferred_lft forever
7: incusbr1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 00:16:3e:67:5a:8a brd ff:ff:ff:ff:ff:ff
    inet 172.16.100.54/24 brd 172.16.100.255 scope global dynamic incusbr1
       valid_lft 6865sec preferred_lft 6865sec
10: tap4615fde0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq master incusbr1 state UP group default qlen 1000
    link/ether ca:b6:fc:a8:17:6a brd ff:ff:ff:ff:ff:ff
12: tap822a48e2: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq master incusbr1 state UP group default qlen 1000
    link/ether 26:ee:d3:cc:f6:e7 brd ff:ff:ff:ff:ff:ff
```

##### Incus network bridge configuration

```yaml
config:
  bridge.external_interfaces: enp2s0
  dns.mode: none
  ipv4.dhcp.gateway: 172.16.100.1
  ipv4.firewall: "false"
  ipv4.nat: "false"
  ipv6.address: none
  ipv6.firewall: "false"
description: ""
name: incusbr1
type: bridge
used_by:
- /1.0/instances/vm03
- /1.0/instances/vm04
- /1.0/profiles/direct-bridge
managed: true
status: Created
locations:
- none
project: default
```

Profile `direct-bridge`

```yaml
config: {}
description: ""
devices:
  eth0:
    network: incusbr1
    type: nic
  root:
    path: /
    pool: default
    size: 10GiB
    type: disk
name: direct-bridge
project: default
```
