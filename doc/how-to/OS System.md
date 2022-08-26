
# Image maintenance

The image created with this playbook is available in the directory `platform/roles/image/files/output/image`.

## Requirements

- packer v1.7.8
- qemu-system v4.2.1
- openstacksdk >= v0.12.0 (necessary for the image upload)

## Original image

The image upon which the cluster operating system image is written is defined in the file `platform/roles/image/defaults/main`. It is defined by:

- the iso url to use
- the sha256 checksum of this iso

## Softwares version

A complete manifest of the installed softwares is available in the directory `platform/roles/image/files/output` at the end of the image creation. The majority of the installed softwares are installed through the aptitude package manager to ensure the latest security updates are used. The exception to that rule is containerd.

The containerd version installed in the image is defined in the file `platform/roles/image/defaults/main`. The version is chosen to better suit the needs of the kubernetes version installed later on.

## Image name

The name under which this image is deployed on the cloud provider is "csc-rs-ubuntu" by default. This name is defined in the file `inventory/mycluster/host_vars/localhost/image.yaml`.

The name for the qcow2 image created in defined in `platform/roles/image/defaults/main`. This name is not very important considering it will not appear anywhere but on the localhost.
