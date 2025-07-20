#!/bin/env bash

sudo apt-get install -y vim curl wget bash-completion gpg git openssh-server chrony

echo "installing network packages"
sudo apt-get install -y net-tools bridge-utils tcpdump traceroute bind9-dnsutils
