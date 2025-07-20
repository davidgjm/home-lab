# Setting apt proxy

## Resources

- [https://linuxconfig.org/setting-up-an-apt-proxy-server-on-debian-linux](https://linuxconfig.org/setting-up-an-apt-proxy-server-on-debian-linux)
- [https://linuxiac.com/how-to-use-apt-with-proxy/](https://linuxiac.com/how-to-use-apt-with-proxy/)

### Configuration file

`vim /etc/apt/apt.conf.d/proxy.conf`

```EditorConfig
Acquire::http::Proxy "[http://172.16.100.17:7890/](http://172.16.100.17:7890/)";
Acquire::https::Proxy "[http://172.16.100.17:7890/](http://172.16.100.17:7890/)";
```


# Change apt repository mirrors
## Sources

|                 |                                                                                                                                                       |
| --------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| China mirror    | [https://mirrors.aliyun.com/debian/](https://mirrors.aliyun.com/debian/?spm=a2c6h.13651104.d-1001.1.4ee127078qrliF)                                   |
| Debian-security | [https://developer.aliyun.com/mirror/debian-security](https://developer.aliyun.com/mirror/debian-security?spm=a2c6h.13651104.d-1001.3.4ee127078qrliF) |

[https://developer.aliyun.com/mirror/debian?spm=a2c6h.13651102.0.0.27331b11woYqlb](https://developer.aliyun.com/mirror/debian?spm=a2c6h.13651102.0.0.27331b11woYqlb)

## Update repository

1. `cp /etc/apt/sources.list{,.orig}`
2. `apt edit-sources`. Change to
	- Main
		- `deb [https://mirrors.aliyun.com/debian/](https://mirrors.aliyun.com/debian/) bookworm contrib main non-free-firmware`

	- Security

		- `deb [https://mirrors.aliyun.com/debian-security](https://mirrors.aliyun.com/debian-security) bookworm-security main non-free-firmware contrib non-free

3. `apt-get update`
4. `apt upgrade -y

## Updated file content

```conf
deb [https://mirrors.aliyun.com/debian/](https://mirrors.aliyun.com/debian/) bookworm contrib main non-free-firmware
deb [https://mirrors.aliyun.com/debian/](https://mirrors.aliyun.com/debian/) bookworm-updates contrib main non-free-firmware
deb [https://mirrors.aliyun.com/debian-security](https://mirrors.aliyun.com/debian-security) bookworm-security main non-free-firmware contrib non-free
```


## Upgrade

1. `sudo apt list --upgradable`


#debian #apt
#ubuntu
