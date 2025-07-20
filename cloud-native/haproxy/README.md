
# Installation

## Ubuntu 24.04

```shell
apt-get install --no-install-recommends software-properties-common
add-apt-repository ppa:vbernat/haproxy-3.1
```

### Install command
```shell
apt-get install haproxy=3.1.\*
```
## Debian 12

Follow quick guide at https://haproxy.debian.net/.

## Haproxy 3.1

### Enable repository
```bash
curl https://haproxy.debian.net/haproxy-archive-keyring.gpg \
      > /usr/share/keyrings/haproxy-archive-keyring.gpg
echo deb "[signed-by=/usr/share/keyrings/haproxy-archive-keyring.gpg]" \
      http://haproxy.debian.net bookworm-backports-3.1 main \
      > /etc/apt/sources.list.d/haproxy.list
```


### apt install

```bash
apt-get update
apt-get install haproxy=3.1.\*
```


# Configuration

The configuration file is located at `/etc/haproxy/haproxy.cfg`

## haproxy configuration guide

A more user friendly guide can be found at:
- https://www.haproxy.com/documentation/haproxy-configuration-tutorials/proxying-essentials/configuration-basics/
- https://www.haproxy.com/documentation/haproxy-configuration-manual/latest/

## TLS Certificates

SSL certificates are located at `/etc/haproxy/ssl/certs`.