
## Installation (not fresh installation)

## Troubleshooting for flannel installation failure


If the `kubernetes.svc.default` service cluster IP is not included in SAN for the api server, flannel will not work because TLS verification will fail. To fix this,

1. Make sure the `--service-cluser-ip-range` is configured correctly at `kube-apiserver`
2. Recreate the default internal service and manually assign an cluster IP.


## Error registering network: failed to acquire lease: node "node-1" pod cidr not assigned

Need to patch the existing nodes with command like below

```shell
kubectl patch node node-1 -p '{"spec":{"podCIDR":"10.200.1.0/24"}}'
```

# flannel configuration (kubeadm)

## Configmap

```shell
kubectl get cm/kube-flannel-cfg -n kube-flannel -o yaml
```

output
```yaml
apiVersion: v1
data:
  cni-conf.json: |
    {
      "name": "cbr0",
      "cniVersion": "0.3.1",
      "plugins": [
        {
          "type": "flannel",
          "delegate": {
            "hairpinMode": true,
            "isDefaultGateway": true
          }
        },
        {
          "type": "portmap",
          "capabilities": {
            "portMappings": true
          }
        }
      ]
    }
  net-conf.json: |
    {
      "Network": "10.244.0.0/16",
      "EnableNFTables": false,
      "Backend": {
        "Type": "vxlan"
      }
    }
kind: ConfigMap
```


# flannel configuration traffic analysis


## check vxlan information


```shell
ip -d link show flannel.1
```

output
```text
3: flannel.1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1450 qdisc noqueue state UNKNOWN mode DEFAULT group default
    link/ether e6:28:e5:cd:55:90 brd ff:ff:ff:ff:ff:ff promiscuity 0  allmulti 0 minmtu 68 maxmtu 65535
    vxlan id 1 local 172.16.100.231 dev eth0 srcport 0 0 dstport 8472 nolearning ttl auto ageing 300 udpcsum noudp6zerocsumtx noudp6zerocsumrx addrgenmode eui64 numtxqueues 1 numrxqueues 1 gso_max_size 65536 gso_max_segs 65535 tso_max_size 65536 tso_max_segs 65535 gro_max_size 65536
```

## Check iptables

### On controller created with kubeadm

```shell
sudo iptables -L
```

Output
```text
Chain FLANNEL-FWD (1 references)
target     prot opt source               destination
ACCEPT     0    --  10.244.0.0/16        0.0.0.0/0            /* flanneld forward */
ACCEPT     0    --  0.0.0.0/0            10.244.0.0/16        /* flanneld forward */
```


## Checking vxlan traffic with tcpdump

#### from source host (controller)
```bash
sudo tcpdump port 8472 and udp -l -A -xxx -i any -n
```


#### from target node (worker)

```shell
sudo tcpdump port 8472 and udp -l -A -xxx -i any -n
```


#### Example dump

```bash
tcpdump: data link type LINUX_SLL2
tcpdump: verbose output suppressed, use -v[v]... for full protocol decode
listening on any, link-type LINUX_SLL2 (Linux cooked v2), snapshot length 262144 bytes
22:49:14.933565 eth0  Out IP 172.16.100.217.42533 > 172.16.100.231.8472: OTV, flags [I] (0x08), overlay 0, instance 1
IP 10.244.0.0.40555 > 10.244.2.2.9443: Flags [S], seq 1570154806, win 64240, options [mss 1460,sackOK,TS val 2814197129 ecr 0,nop,wscale 7], length 0
	0x0000:  0800 0000 0000 0002 0001 0406 bc24 11e3
	0x0010:  0063 0000 4500 006e 19cc 0000 4011 3ed2
	0x0020:  ac10 64d9 ac10 64e7 a625 2118 005a 5a14
	0x0030:  0800 0000 0000 0100 e628 e5cd 5590 ca58
	0x0040:  469f 909f 0800 4500 003c 1351 4000 4006
	0x0050:  0f82 0af4 0000 0af4 0202 9e6b 24e3 5d96
	0x0060:  a936 0000 0000 a002 faf0 85c3 0000 0204
	0x0070:  05b4 0402 080a a7bd 3d89 0000 0000 0103
	0x0080:  0307
22:49:14.933887 eth0  In  IP 172.16.100.231.47229 > 172.16.100.217.8472: OTV, flags [I] (0x08), overlay 0, instance 1
IP 10.244.2.2.9443 > 10.244.0.0.40555: Flags [S.], seq 2097266526, ack 1570154807, win 64308, options [mss 1410,sackOK,TS val 3185890486 ecr 2814197129,nop,wscale 7], length 0
	0x0000:  0800 0000 0000 0002 0001 0006 bc24 1146
	0x0010:  e1ad 0000 4500 006e c025 0000 4011 9878
	0x0020:  ac10 64e7 ac10 64d9 b87d 2118 005a 47bc
	0x0030:  0800 0000 0000 0100 ca58 469f 909f e628
	0x0040:  e5cd 5590 0800 4500 003c 0000 4000 3f06
	0x0050:  23d3 0af4 0202 0af4 0000 24e3 9e6b 7d01
	0x0060:  bf5e 5d96 a937 a012 fb34 b6a4 0000 0204
	0x0070:  0582 0402 080a bde4 d4b6 a7bd 3d89 0103
	0x0080:  0307
22:49:14.933912 eth0  Out IP 172.16.100.217.42533 > 172.16.100.231.8472: OTV, flags [I] (0x08), overlay 0, instance 1
IP 10.244.0.0.40555 > 10.244.2.2.9443: Flags [.], ack 1, win 502, options [nop,nop,TS val 2814197130 ecr 3185890486], length 0
	0x0000:  0800 0000 0000 0002 0001 0406 bc24 11e3
	0x0010:  0063 0000 4500 0066 19cd 0000 4011 3ed9
	0x0020:  ac10 64d9 ac10 64e7 a625 2118 0052 5a1c
	0x0030:  0800 0000 0000 0100 e628 e5cd 5590 ca58
	0x0040:  469f 909f 0800 4500 0034 1352 4000 4006
	0x0050:  0f89 0af4 0000 0af4 0202 9e6b 24e3 5d96
	0x0060:  a937 7d01 bf5f 8010 01f6 de7c 0000 0101
	0x0070:  080a a7bd 3d8a bde4 d4b6
22:49:14.933954 eth0  Out IP 172.16.100.217.42533 > 172.16.100.231.8472: OTV, flags [I] (0x08), overlay 0, instance 1
IP 10.244.0.0.40555 > 10.244.2.2.9443: Flags [F.], seq 1, ack 1, win 502, options [nop,nop,TS val 2814197130 ecr 3185890486], length 0
	0x0000:  0800 0000 0000 0002 0001 0406 bc24 11e3
	0x0010:  0063 0000 4500 0066 19ce 0000 4011 3ed8
	0x0020:  ac10 64d9 ac10 64e7 a625 2118 0052 5a1c
	0x0030:  0800 0000 0000 0100 e628 e5cd 5590 ca58
	0x0040:  469f 909f 0800 4500 0034 1353 4000 4006
	0x0050:  0f88 0af4 0000 0af4 0202 9e6b 24e3 5d96
	0x0060:  a937 7d01 bf5f 8011 01f6 de7b 0000 0101
	0x0070:  080a a7bd 3d8a bde4 d4b6
22:49:14.934437 eth0  In  IP 172.16.100.231.47229 > 172.16.100.217.8472: OTV, flags [I] (0x08), overlay 0, instance 1
IP 10.244.2.2.9443 > 10.244.0.0.40555: Flags [F.], seq 1, ack 2, win 503, options [nop,nop,TS val 3185890486 ecr 2814197130], length 0
	0x0000:  0800 0000 0000 0002 0001 0006 bc24 1146
	0x0010:  e1ad 0000 4500 0066 c026 0000 4011 987f
	0x0020:  ac10 64e7 ac10 64d9 b87d 2118 0052 47c4
	0x0030:  0800 0000 0000 0100 ca58 469f 909f e628
	0x0040:  e5cd 5590 0800 4500 0034 4f7c 4000 3f06
	0x0050:  d45e 0af4 0202 0af4 0000 24e3 9e6b 7d01
	0x0060:  bf5f 5d96 a938 8011 01f7 de79 0000 0101
	0x0070:  080a bde4 d4b6 a7bd 3d8a
22:49:14.934448 eth0  Out IP 172.16.100.217.42533 > 172.16.100.231.8472: OTV, flags [I] (0x08), overlay 0, instance 1
IP 10.244.0.0.40555 > 10.244.2.2.9443: Flags [.], ack 2, win 502, options [nop,nop,TS val 2814197130 ecr 3185890486], length 0
	0x0000:  0800 0000 0000 0002 0001 0406 bc24 11e3
	0x0010:  0063 0000 4500 0066 19cf 0000 4011 3ed7
	0x0020:  ac10 64d9 ac10 64e7 a625 2118 0052 5a1c
	0x0030:  0800 0000 0000 0100 e628 e5cd 5590 ca58
	0x0040:  469f 909f 0800 4500 0034 1354 4000 4006
	0x0050:  0f87 0af4 0000 0af4 0202 9e6b 24e3 5d96
	0x0060:  a938 7d01 bf60 8010 01f6 de7a 0000 0101
	0x0070:  080a a7bd 3d8a bde4 d4b6

```