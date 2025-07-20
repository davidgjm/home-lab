
# clustering with `cephadm`


## Console output when complete

```shell
Ceph Dashboard is now available at:

	     URL: https://vm04:8443/
	    User: admin
	Password: xxxxxxxxx

Enabling client.admin keyring and conf on hosts with "admin" label
Saving cluster configuration to /var/lib/ceph/d50ccbc6-afcc-11ef-b019-00163e47f9ef/config directory
Enabling autotune for osd_memory_target
You can access the Ceph CLI as following in case of multi-cluster or non-default config:

	sudo /usr/sbin/cephadm shell --fsid d50ccbc6-afcc-11ef-b019-00163e47f9ef -c /etc/ceph/ceph.conf -k /etc/ceph/ceph.client.admin.keyring

Or, if you are only running a single cluster on this host:

	sudo /usr/sbin/cephadm shell

Please consider enabling telemetry to help improve Ceph:

	ceph telemetry on

For more information see:

	https://docs.ceph.com/docs/master/mgr/telemetry/

Bootstrap complete.
```