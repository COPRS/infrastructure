# Image maintenance

The image created with this playbook is available in the directory `roles/image/files/output/image`.

## Requirements

- `packer` v1.7.8
- `qemu-system` v4.2.1
- `openstacksdk` >= v0.12.0 (necessary for the image upload)

## Original image

The image upon which the cluster operating system image is written is defined in the file `roles/image/defaults/main.yaml`. It is defined by:

- The remote URL of the orignal ISO
- The sha256 checksum of this ISO

## Softwares version

A complete manifest of the installed softwares is available in the directory `roles/image/files/output` at the end of the image creation. The majority of the installed softwares are installed through the *aptitude package manager* to ensure the latest security updates are used. The exception to that rule is `containerd`.

The `containerd` version installed in the image is defined in the file `roles/image/defaults/main.yaml`. This version is chosen to better suit the needs of the kubernetes version installed later on.

## Image name

The name under which this image is deployed on the cloud provider is "csc-rs-ubuntu" by default. This name is defined in the file `roles/image/defaults/main.yaml`.

The name for the *qcow2* image created in defined in `roles/image/defaults/main.yaml`. This name is not significant considering it will not appear anywhere but on the localhost.
