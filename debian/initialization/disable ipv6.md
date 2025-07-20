
[https://www.howtoraspberry.com/2020/04/disable-ipv6-on-raspberry-pi/](https://www.howtoraspberry.com/2020/04/disable-ipv6-on-raspberry-pi/)

# Edit `/add file /etc/sysctl.d/10-local.conf

Add new `local.conf `to `/etc/sysctl.d/` if it does not exist

```c
net.ipv6.conf.all.disable_ipv6 = 1
net.ipv6.conf.default.disable_ipv6=1
net.ipv6.conf.lo.disable_ipv6 = 1
```

# Apply changes

`sudo sysctl -p`

After making any changes, please run "`service procps force-reload`" (or, from a Debian package maintainer script "`deb-systemd-invoke restart procps.service`").

## Test the configuration

```shell
cat /proc/sys/net/ipv6/conf/all/disable_ipv6
1
```

### `ip a`

```shell
david@netsrv1:~ $ ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    inet 172.16.100.4/24 brd 172.16.100.255 scope global dynamic noprefixroute eth0
       valid_lft 21474677sec preferred_lft 21474677sec
```


#debian #ubuntu #linux #ipv6
#networking

