#!/bin/bash

sudo apt-get autoremove -y
sudo apt-get autoclean -y

sudo snap install ubuntu-package-manifest
snap connect ubuntu-package-manifest:system-data
sudo ubuntu-package-manifest > final-manifest.txt
sudo snap remove ubuntu-package-manifest