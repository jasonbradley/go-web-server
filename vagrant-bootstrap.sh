#!/usr/bin/env bash

VAGRANT_PROVISIONED=/etc/vagrant-provisioned

if [ -e $VAGRANT_PROVISIONED ]
then
    exit 0
fi

apt-get update

apt-get install -y golang git

mkdir /var/apps/goserver; echo "export GOPATH=/var/apps/gowebserver" >> ~/.bashrc
echo "export PATH=$PATH:$HOME/go/bin:/usr/local/go/bin" >> ~/.bashrc
