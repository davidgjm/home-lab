
# home lab servers


| Gateway      | DNS server   |
| ------------ | ------------ |
| 172.16.100.1 | 172.31.255.2 |

## issues
1. `m5` - hardware issue with NIC. tweaks needed to pin the MAC address. See [GMKtec M5 Plus NIC issue fix](networking/GMKtec%20M5%20Plus%20NIC%20issue%20fix.md)
2. `m7` - the ssd sequence is not correct.

| Hostname      | CPU/RAM          | Storage                                                        | network interfaces                                                |
| ------------- | ---------------- | -------------------------------------------------------------- | ----------------------------------------------------------------- |
| m5.home.lab   | 16 core 64GB RAM | - nvme0n1: 1TB (system disk)<br>- nvme1n1: 2TB (data disk)     | - enp1s0: 172.16.100.33/24<br>- enp2s0: 172.16.100.34/24          |
| m7.home.lab   | 16 core 64GB RAM | - **nvme1n1**: 2TB (system disk)<br>- nvme0n1: 2TB (data disk) | - enp1s0: 172.16.100.35/24<br>- eno1: 172.16.100.36/24            |
| ser6.home.lab | 16 core 64GB RAM | - nvme0n1: 1TB (system disk)<br>- sda: 2TB (data disk)         | - enp2s0: 172.16.100.37/24<br>- enx000fc929fe3f: 172.16.100.38/24 |
