# Virtual Machine Management

## Overview

U-Server uses KVM/QEMU with libvirt for full virtual machine support.

The virtualization stack includes:

```text
KVM
QEMU
libvirt
virsh
virt-install
bridge-utils
```

Intel VT-x hardware virtualization is enabled on U-Server.

Verify virtualization support:

```bash
lscpu | grep -E 'Virtualization|Hypervisor'
```

Verify KVM kernel modules:

```bash
lsmod | grep kvm
```

Expected modules include:

```text
kvm_intel
kvm
```

## libvirt Service

Check libvirt:

```bash
systemctl status libvirtd --no-pager
```

Check group access:

```bash
getent group libvirt
getent group kvm
```

The `jonathon` user belongs to both groups and can use normal `virsh` commands without `sudo`.

## List Virtual Machines

List running VMs:

```bash
virsh list
```

List all VMs:

```bash
virsh list --all
```

## View VM Information

```bash
virsh dominfo <vm-name>
```

Example:

```bash
virsh dominfo ubuntu-test
```

This displays information such as:

- Current state
- Virtual CPUs
- Memory
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

Request a normal guest shutdown:

```bash
virsh shutdown <vm-name>
```

A normal shutdown should be preferred whenever possible.

## Force Stop a VM

If the guest cannot shut down normally:

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

Connect to the guest serial console:

```bash
virsh console <vm-name>
```

Example:

```bash
virsh console ubuntu-test
```

Detach from the console without stopping the VM:

```text
Ctrl + ]
```

## VM Networking

U-Server uses the default libvirt NAT network for the current test VM.

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

View a VM's network interfaces:

```bash
virsh domiflist <vm-name>
```

Attempt to determine its guest IP address:

```bash
virsh domifaddr <vm-name>
```

libvirt provides DHCP through `dnsmasq` for the default NAT network.

## VM Storage

View block devices assigned to a VM:

```bash
virsh domblklist <vm-name>
```

Virtual machine disk images are normally stored under:

```text
/var/lib/libvirt/images/
```

Installation ISO images are stored under:

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

The current `ubuntu-test` VM has autostart disabled because it is a test environment.

## Remove a VM

Before removing a VM, shut it down and confirm its data is no longer required.

Remove the VM definition while retaining storage:

```bash
virsh undefine <vm-name>
```

Remove the VM definition and associated storage:

```bash
virsh undefine <vm-name> --remove-all-storage
```

The second command permanently removes associated virtual disks and should be used carefully.

## Test Virtual Machine

The current virtualization test VM is:

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

## Troubleshooting

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

Check its network interfaces:

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
