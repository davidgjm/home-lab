
# Run `ansible` playbook

```shell
cd $(git rev-parse --show-toplevel)
ansible-playbook debian-playbook.yaml  -i inventory/inventory.yaml
```

# `ansible` tricks
## Provide `sudo` password

```shell
ansible-playbook debian-playbook.yaml  -i inventory/inventory.yaml -K
```


## Shutdown all servers

```shell
ansible all -m ansible.builtin.command -a "/sbin/shutdown -P now" --become --ask-become-pass --inventory=inventory/inventory.yaml
```

### Simplified version when inventory is configured through `~/.ansible.cfg`

```shell
ansible all -m ansible.builtin.command -a "/sbin/shutdown -P now" --become --ask-become-pass
```