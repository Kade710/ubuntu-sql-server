# Hardware Specifications

## Server Information

Hostname:
- U-Server

Operating System:
- Ubuntu 24.04.4 LTS

Kernel:
- Linux 6.17.0-35-generic

Architecture:
- x86_64

Hardware Vendor:
- MSI

Motherboard:
- MSI MS-7917
- Intel Z97 Chipset
- Firmware Version: V1.8


# Processor

Model:
- Intel Core i5-4690K CPU @ 3.50GHz

Architecture:
- x86_64

Cores:
- 4 physical cores

Threads:
- 4 threads

Base Frequency:
- 3.50 GHz

Maximum Turbo Frequency:
- 4.00 GHz

Cache:
- L1: 128 KiB Data + 128 KiB Instruction
- L2: 1 MiB
- L3: 6 MiB

Virtualization:
- Intel VT-x Supported


# Memory

Installed RAM:
- 16 GiB

Available RAM:
- 13 GiB

Swap:
- 4 GiB


# Storage

Primary Storage:
- Capacity: 1.8 TB

Partition Layout:

- EFI Partition: 1 GB
- Root Partition: 1.8 TB

Filesystem:
- Linux root filesystem


# Graphics

GPU:
- NVIDIA GeForce GTX 750

PCI Device:
- NVIDIA GM107


# Networking

Ethernet:
- Qualcomm Atheros Killer E220x Gigabit Ethernet Controller

Network Capability:
- 1 Gigabit Ethernet


# Server Purpose

Current Goals:
- Ubuntu Linux server
- PostgreSQL database server
- SQL development environment
- Linux administration practice
- Future application hosting

Future Expansion:
- Additional storage
- Monitoring
- Automation
- Application deployment

ip addr

1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host noprefixroute 
       valid_lft forever preferred_lft forever
2: enp3s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP group default qlen 1000
    link/ether 02:00:00:00:00:01 brd ff:ff:ff:ff:ff:ff
    inet 192.168.1.100/24 brd 192.168.12.255 scope global dynamic noprefixroute enp3s0
       valid_lft 602243sec preferred_lft 602243sec
    inet6 2607:fb91:4e85:90ac:9b9f:2240:2539:30c4/64 scope global temporary dynamic 
       valid_lft 1786sec preferred_lft 1786sec
    inet6 2607:fb91:4e85:90ac:dacb:8aff:fe3b:f912/64 scope global dynamic mngtmpaddr 
       valid_lft 1786sec preferred_lft 1786sec
    inet6 fdfe:5e6c:2e69:98de:f8af:6f6:79c6:7ebe/64 scope global temporary dynamic 
       valid_lft 3586sec preferred_lft 3586sec
    inet6 fdfe:5e6c:2e69:98de:dacb:8aff:fe3b:f912/64 scope global dynamic mngtmpaddr 
       valid_lft 3586sec preferred_lft 3586sec
    inet6 fe80::dacb:8aff:fe3b:f912/64 scope link 
       valid_lft forever preferred_lft forever
