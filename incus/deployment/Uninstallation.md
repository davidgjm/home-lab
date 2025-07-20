
# Manual uninstallation
1. `systemctl stop incus.socket incus.service`
2. `apt remove -y incus incus-base incus-client`
3. `rm -rf /var/lib/incus*`


  

incus config set core.https_address :8443


