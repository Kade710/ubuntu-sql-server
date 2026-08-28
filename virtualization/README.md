# Virtualization and Containers

## Overview

U-Server uses KVM/QEMU with libvirt for full virtual machines and Docker with Docker Compose for containerized applications.

Management procedures are separated into dedicated documentation:

- [Virtual Machine Management](vm_management.md)
- [Container Management](container_management.md)

## Virtual Machine Platform

U-Server uses:

```text
KVM
QEMU
libvirt
virsh
virt-install
bridge-utils
```

Intel VT-x hardware-assisted virtualization is enabled.

The current test virtual machine is:

```text
Name: ubuntu-test
Guest OS: Ubuntu Server 24.04.4 LTS
vCPUs: 2
Memory: 3 GiB
Virtual Disk: 20 GiB
Network: libvirt default NAT
IPv4 Address: 192.168.122.154
Persistent: yes
Autostart: disabled
```

For VM administration, networking, storage, console access, autostart, removal, and troubleshooting, see:

[Virtual Machine Management](vm_management.md)

## Container Platform

U-Server uses:

```text
Docker
Docker Compose
```

Docker data is stored under:

```text
/var/lib/docker
```

The current documented container environment includes the Bittensor development environment.

For Docker administration, Compose, networking, storage, logging, security, and troubleshooting, see:

[Container Management](container_management.md)

## Documentation Structure

```text
virtualization/
├── README.md
├── vm_management.md
└── container_management.md
```

The documentation in this directory should be updated whenever major virtual machines, containers, networks, storage configurations, or virtualization technologies are added, removed, or changed.