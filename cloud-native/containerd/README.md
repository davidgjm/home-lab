
# Install dependencies

```bash
{
  sudo apt-get update
  sudo apt-get -y install socat conntrack ipset kmod
}
```

# Tools for container runtime


```bash
ARCH=$(dpkg --print-architecture)
cat downloads-${ARCH}.txt
```


| Tool          | Usage                         |
| ------------- | ----------------------------- |
| `containerd`  | Container Runtime             |
| `crictl`      | Container Runtime client tool |
| `runc`        | Low level container runtime   |
| `cni-plugins` | CNI plugins                   |


# Download Tools


```bash
wget -q --show-progress \
  --https-only \
  --timestamping \
  -P downloads \
  -i downloads-$(dpkg --print-architecture).txt
```


## Extract

```bash
{
  ARCH=$(dpkg --print-architecture)
  mkdir -p downloads/{cni-plugins,worker}
  tar -xvf downloads/crictl-v1.32.0-linux-${ARCH}.tar.gz \
    -C downloads/worker/
  tar -xvf downloads/containerd-2.1.1-linux-${ARCH}.tar.gz \
    --strip-components 1 \
    -C downloads/worker/
  tar -xvf downloads/cni-plugins-linux-${ARCH}-v1.7.0.tgz \
    -C downloads/cni-plugins/
  mv downloads/runc.${ARCH} downloads/worker/runc
}
```

# Install

```bash
sudo mkdir -p \
  /etc/cni/net.d \
  /opt/cni/bin
```

Install binaries:
```bash
{
  cd downloads
  sudo mv worker/crictl worker/ctr worker/runc \
    /usr/local/bin/
  sudo mv worker/containerd worker/containerd-shim-runc-v2 worker/containerd-stress /bin/
  sudo mv cni-plugins/* /opt/cni/bin/
}
```

# Configure

## Configure CNI Networking

```bash
{
  modprobe br-netfilter
  echo "br-netfilter" >> /etc/modules-load.d/modules.conf
}
```

```bash
{
  echo "net.bridge.bridge-nf-call-iptables = 1" \
    >> /etc/sysctl.d/kubernetes.conf
  echo "net.bridge.bridge-nf-call-ip6tables = 1" \
    >> /etc/sysctl.d/kubernetes.conf
  sysctl -p /etc/sysctl.d/kubernetes.conf
}
```

## Configure `containerd`

upload config files to node

```bash
scp -r configs debian@172.16.100.233:~/
```


```bash
{
  sudo mkdir -p /etc/containerd/
  sudo mv configs/containerd-config.toml /etc/containerd/config.toml
  sudo mv configs/containerd.service /etc/systemd/system/
}
```


### Start services

```bash
{
  sudo systemctl daemon-reload
  sudo systemctl enable containerd
  sudo systemctl start containerd
```

