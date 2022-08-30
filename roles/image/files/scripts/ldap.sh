#!/bin/bash

# Install necessary ldap packages
sudo DEBIAN_FRONTEND=nointerractive \
apt install -y ldap-auth-client=0.5.4 \
ldap-auth-config=0.5.4 \
libnss-ldap=265-5ubuntu1 \
libpam-ldap=186-4ubuntu1 \
ldap-utils
