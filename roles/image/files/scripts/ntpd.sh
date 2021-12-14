#!/bin/bash

# Install ntp server

sudo apt-get install -y ntp

sudo systemctl enable ntp
sudo systemctl restart ntp
