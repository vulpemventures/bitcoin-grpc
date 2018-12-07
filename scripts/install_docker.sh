#!/bin/bash

set -e

# add docker GPG to the system
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

# add docker repo to apt
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable edge"

# update package db
apt-cache policy docker-ce

# install docker
apt-get install -y docker-ce

