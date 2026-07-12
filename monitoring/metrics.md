# Monitoring Metrics

## Overview

This document defines the metrics used to monitor the health and performance of U-Server.

Server:

* U-Server

Operating System:

* Ubuntu 24.04.4 LTS Desktop

Monitoring Goals:

* Verify system availability
* Monitor hardware utilization
* Verify PostgreSQL service status
* Monitor database connectivity
* Detect resource utilization issues

# Current System Baseline

## CPU

Status:

* Intel Core i5-4690K
* 4 Physical Cores
* 4 Threads

Current Status:

* Operating normally

## Memory

Installed:

* 16 GB DDR3

Current Usage:

* Used: 1.3 GB
* Free: 10 GB
* Available: 14 GB
* Swap: 4 GB (0 GB in use)

## Storage

Primary Disk:

* /dev/sda
* 1.8 TB

Current Usage:

* Used: 14 GB
* Available: 1.7 TB
* Utilization: 1%

## System Uptime

Current Uptime:

* 21 hours

Load Average:

* 1 Minute: 0.00
* 5 Minutes: 0.00
* 15 Minutes: 0.00

# Network Metrics

Primary Interface:

* enp3s0

IPv4 Address:

* 192.168.1.100

Speed:

* 1 Gigabit Ethernet

Connectivity Test:

* 4/4 packets received
* 0% packet loss
* Average latency: 33.3 ms

# PostgreSQL Metrics

Service:

* postgresql@16-main

Status:

* Active (running)

Database:

* ubuntu_sql_server

Database Owner:

* jonathon_admin

Current Tables:

* hardware_components
* maintenance_logs
* network_interfaces
* operating_systems
* server_inventory

# Monitoring Commands

## CPU

```bash
lscpu
```

## Memory

```bash
free -h
```

## Storage

```bash
df -h
```

## Disk Layout

```bash
lsblk
```

## System Uptime

```bash
uptime
```

## Network Configuration

```bash
ip addr show
```

## Network Connectivity

```bash
ping -c 4 8.8.8.8
```

## PostgreSQL Status

```bash
systemctl status postgresql@16-main
```

## PostgreSQL Login

```bash
psql -U jonathon_admin -d ubuntu_sql_server
```

## Database Tables

```sql
\dt server_management.*
```
