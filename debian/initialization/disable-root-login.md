# Disable `root` account
## Resources

- [https://www.howtogeek.com/828538/how-and-why-to-disable-root-login-over-ssh-on-linux/](https://www.howtogeek.com/828538/how-and-why-to-disable-root-login-over-ssh-on-linux/)
- [https://unix.stackexchange.com/questions/383301/should-i-disable-the-root-account-on-my-debian-pc-for-security](https://unix.stackexchange.com/questions/383301/should-i-disable-the-root-account-on-my-debian-pc-for-security)

## Disable root login

1. `vim /etc/ssh/sshd_config`
2. Add/uncomment `PermitRootLogin no` in the file
3. `sudo systemctl restart ssh`

## Disable password login

1. `sudo passwd root -ld`