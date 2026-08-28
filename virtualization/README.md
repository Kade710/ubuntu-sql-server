# Virtualization and Containers

## Overview

U-Server uses two forms of application isolation and virtualization:

- KVM/QEMU with libvirt for full virtual machines
- Docker and Docker Compose for containers

KVM provides hardware-assisted virtualization for running complete guest operating systems.

Docker provides isolated application environments while sharing the U-Server Linux kernel.

This document contains the primary management and troubleshooting commands for both environments.

---

# Virtual Machines

## Platform

U-Server uses the following virtualization stack:

```text
KVM
QEMU
libvirt
virsh
virt-install
bridge-utils
```

Intel VT-x hardware virtualization is enabled on the host.

Verify virtualization support:

```bash
lscpu | grep -E 'Virtualization|Hypervisor'
```

Expected virtualization type:

```text
VT-x
```

Verify the KVM kernel modules:

```bash
lsmod | grep kvm
```

The host currently uses:

```text
kvm_intel
kvm
```

## libvirt

libvirt manages the virtual machines running on U-Server.

Check the service:

```bash
systemctl status libvirtd --no-pager
```

Check the libvirt and KVM groups:

```bash
getent group libvirt
getent group kvm
```

The `jonathon` user belongs to both groups and can perform normal `virsh` management without `sudo`.

## List Virtual Machines

List running VMs:

```bash
virsh list
```

List all VMs:

```bash
virsh list --all
```

## VM Information

Display detailed information about a VM:

```bash
virsh dominfo <vm-name>
```

Example:

```bash
virsh dominfo ubuntu-test
```

This displays information such as:

- Current state
- Number of virtual CPUs
- Memory allocation
- Persistence
- Autostart status

## Start a VM

```bash
virsh start <vm-name>
```

Example:

```bash
virsh start ubuntu-test
```

## Shut Down a VM

Request a normal operating system shutdown:

```bash
virsh shutdown <vm-name>
```

A normal shutdown should be preferred whenever possible.

## Force Stop a VM

If a VM cannot shut down normally:

```bash
virsh destroy <vm-name>
```

This immediately stops the VM and is similar to removing power from a physical computer.

Use it only when necessary.

## Reboot a VM

```bash
virsh reboot <vm-name>
```

## VM Console

Connect to a guest serial console:

```bash
virsh console <vm-name>
```

Example:

```bash
virsh console ubuntu-test
```

Detach from the console without stopping the VM with:

```text
Ctrl + ]
```

## VM Networking

The current virtual machines use the default libvirt NAT network.

View libvirt networks:

```bash
virsh net-list --all
```

The default network is configured as:

```text
Name: default
State: active
Autostart: yes
Persistent: yes
Network: 192.168.122.0/24
```

View the virtual network interfaces assigned to a VM:

```bash
virsh domiflist <vm-name>
```

Attempt to determine a guest IP address:

```bash
virsh domifaddr <vm-name>
```

libvirt provides DHCP through `dnsmasq` for the default NAT network.

## VM Storage

View the block devices assigned to a VM:

```bash
virsh domblklist <vm-name>
```

Virtual machine disk images are normally stored in:

```text
/var/lib/libvirt/images/
```

Installation ISO images are stored in:

```text
/var/lib/libvirt/iso/
```

The current Ubuntu Server installation ISO is:

```text
/var/lib/libvirt/iso/ubuntu-24.04.4-live-server-amd64.iso
```

The ISO was verified against the official Ubuntu SHA256 checksum before installation.

## VM Autostart

Check autostart status:

```bash
virsh dominfo <vm-name>
```

Enable autostart:

```bash
virsh autostart <vm-name>
```

Disable autostart:

```bash
virsh autostart --disable <vm-name>
```

Test VMs should normally remain disabled unless they need to run continuously.

## Remove a VM

Before removing a VM, shut it down and confirm that its data is no longer required.

Remove the persistent VM definition while retaining its storage:

```bash
virsh undefine <vm-name>
```

Remove the definition and associated storage:

```bash
virsh undefine <vm-name> --remove-all-storage
```

The second command permanently removes associated virtual disks and should be used carefully.

---

# Test Virtual Machine

The current virtualization test system is:

```text
Name: ubuntu-test
Guest OS: Ubuntu Server 24.04.4 LTS
Guest Kernel: 6.8.0-138-generic
vCPUs: 2
Memory: 3 GiB
Virtual Disk: 20 GiB
Network: libvirt default NAT
IPv4 Address: 192.168.122.154
Persistent: yes
Autostart: disabled
```

The VM uses VirtIO virtual hardware for its disk and network interface.

OpenSSH Server is installed in the guest.

The test VM has been verified to:

- Install successfully with KVM/QEMU and libvirt
- Boot from its virtual disk
- Boot without the installation ISO attached
- Run as a persistent libvirt domain
- Use two virtual CPUs
- Use 3 GiB of memory
- Connect to the default libvirt NAT network
- Receive an IPv4 address through DHCP
- Provide a working serial console
- Run Ubuntu Server 24.04.4 LTS
- Use Intel VT-x hardware-assisted virtualization

---

# Docker

## Platform

Docker and Docker Compose are installed on U-Server.

Check Docker:

```bash
docker --version
```

Check Docker Compose:

```bash
docker compose version
```

Docker stores its data under:

```text
/var/lib/docker
```

## Container Management

List running containers:

```bash
docker ps
```

List all containers:

```bash
docker ps -a
```

Start a container:
=======
## Current Docker Images

Docker images currently stored on U-Server include:

```text
docker-bittensor-node:latest
nvidia/cuda:12.0.0-base-ubuntu22.04
```

```bash
docker start <container-name>
```

Stop a container:

```bash
docker stop <container-name>
=======
Unused images may remain available so containers can be recreated later without downloading or rebuilding the image again.

## Compose Projects

Two Docker Compose configurations are currently stored under the Projects directory.

```text
/home/jonathon/Projects/bittensor-node/docker/docker-compose.yaml
>>>>>>> a738d80 (udated virtualization readme.md)
```

Restart a container:

```bash
docker restart <container-name>
```

Inspect a container:

```bash
docker inspect <container-name>
```

## Container Logs

View logs:

```bash
docker logs <container-name>
```

View recent logs:

```bash
docker logs --tail 50 <container-name>
```

Follow live logs:

```bash
docker logs -f <container-name>
```

Press `Ctrl+C` to stop following logs.

## Docker Compose

Compose projects should normally be managed from the directory containing their Compose configuration.

Validate a Compose configuration:

```bash
docker compose config
```

Start a project:

```bash
docker compose up -d
```

View project containers:

```bash
docker compose ps
```

Stop a project:

```bash
docker compose stop
```

Restart a project:

```bash
docker compose restart
```

Remove project containers and networks:

```bash
docker compose down
```

Persistent storage should be understood before removing volumes or application data.

## Bittensor Development Container

The Bittensor Docker project is stored at:

```text
~/Projects/bittensor-node/docker/
```

Its Compose configuration is:

```text
~/Projects/bittensor-node/docker/docker-compose.yaml
```

The development container is:

```text
bittensor-node-dev
```

The project uses:

```text
docker-bittensor-node:latest
```

Check the container:

```bash
docker ps
```

View its logs:

```bash
docker logs --tail 50 bittensor-node-dev
```

<<<<<<< HEAD
=======
Live logs can be followed with:

```bash
docker logs -f bittensor-node-dev
```

Press `Ctrl+C` to stop following the logs.

>>>>>>> a738d80 (udated virtualization readme.md)
## Docker Networks

View Docker networks:

```bash
docker network ls
```

Inspect a network:

```bash
docker network inspect <network-name>
```

Common Docker network types include:

```text
bridge
host
none
```

Docker Compose may also create project-specific bridge networks.

Docker networking creates virtual Linux interfaces such as:

```text
docker0
br-*
```

These interfaces may be detected by the Ubuntu SQL Server monitoring system.

## Docker Storage

View named volumes:

```bash
docker volume ls
```

Check Docker disk usage:

```bash
docker system df
```

View detailed disk usage:

```bash
docker system df -v
```

Important application data should use persistent storage such as:

- Bind mounts
- Host directories
- Docker volumes

Important persistent data should also be included in the server backup strategy.

## Docker Resource Monitoring

Monitor running containers:

```bash
docker stats
```

This reports information such as:

- CPU usage
- Memory usage
- Network activity
- Block I/O
- Process count

Press `Ctrl+C` to exit.

## Docker Service

Check Docker:

```bash
systemctl status docker
```

Start Docker:

```bash
sudo systemctl start docker
```

Restart Docker:

```bash
sudo systemctl restart docker
```

View recent Docker service logs:

```bash
journalctl -u docker -n 50 --no-pager
```

---

# Security

Virtual machines and containers should follow the same security principles as the rest of U-Server.

- Only trusted users should receive virtualization or Docker access.
- Only required network services should be exposed.
- Guest operating systems should receive security updates.
- Container images should be kept updated.
- Passwords, tokens, and private keys must not be committed to Git.
- Environment files containing secrets must be protected.
- Containers should not receive unnecessary privileges.
- VM and container storage should be backed up when it contains important data.
- Firewall rules should be reviewed whenever a new network service is deployed.
- Logs should be reviewed after unexpected behavior.

Docker group membership should be treated as privileged host access.

---

# Troubleshooting

## Virtual Machines

Check libvirt:

```bash
systemctl status libvirtd --no-pager
```

Check all VMs:

```bash
virsh list --all
```

Inspect a VM:

```bash
virsh dominfo <vm-name>
```

Check its storage:

```bash
virsh domblklist <vm-name>
```

Check its network interface:

```bash
virsh domiflist <vm-name>
```

Check libvirt networks:

```bash
virsh net-list --all
```

Attempt console access:

```bash
virsh console <vm-name>
```

## Containers

Check Docker:

```bash
systemctl status docker
```

Check containers:

```bash
docker ps -a
```

Check container logs:

```bash
docker logs --tail 50 <container-name>
```

Check Docker networks:

```bash
docker network ls
```

Check listening host ports:

```bash
sudo ss -lntp
```

Check the firewall:

```bash
sudo ufw status numbered
```

For Compose projects:

```bash
docker compose config
docker compose ps
```

---

# Current Environment

```text
Host:
- U-Server
- Ubuntu 24.04.4 LTS

Virtual Machine Platform:
- KVM
- QEMU
- libvirt

VM Management:
- virsh
- virt-install

libvirt Network:
- default
- NAT
- 192.168.122.0/24
- autostart enabled

Test VM:
- ubuntu-test
- Ubuntu Server 24.04.4 LTS
- 2 vCPUs
- 3 GiB RAM
- 20 GiB virtual disk
- 192.168.122.154
- autostart disabled

Container Platform:
- Docker
- Docker Compose

Docker Root:
- /var/lib/docker

Bittensor Container:
- bittensor-node-dev

Bittensor Image:
- docker-bittensor-node:latest

CUDA Image:
- nvidia/cuda:12.0.0-base-ubuntu22.04

Docker Networks:
- bridge
- docker_default
- host
- none

Named Docker Volumes:
- None currently listed

Compose Projects:
- bittensor-node/docker
```

This document should be updated whenever major virtual machines, containers, networks, storage configurations, or virtualization technologies are added, removed, or changed.
