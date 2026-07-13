# Environment Requirements

## Overview

This document defines the hardware, software, and access requirements needed to build and operate the Ubuntu SQL Server environment.

The requirements are based on the current U-Server development environment and are intended to support PostgreSQL database hosting, Linux administration, automation development, and future application deployment.

# Hardware Requirements

## Minimum Requirements

| Component | Requirement                             |
| --------- | --------------------------------------- |
| Processor | 64-bit x86 CPU                          |
| Memory    | 8GB RAM minimum                         |
| Storage   | 100GB available storage minimum         |
| Network   | Ethernet or reliable network connection |

## Recommended Requirements

| Component | Recommendation                     |
| --------- | ---------------------------------- |
| Processor | Multi-core 64-bit processor        |
| Memory    | 16GB RAM or higher                 |
| Storage   | 1TB+ storage depending on workload |
| Network   | Gigabit Ethernet connection        |

# Current Hardware Environment

The current production development server uses:

| Component | Specification                                  |
| --------- | ---------------------------------------------- |
| Processor | Intel Core i5-4690K @ 3.50GHz                  |
| Memory    | 16GB DDR3 RAM                                  |
| Storage   | 2TB drive                                      |
| Graphics  | NVIDIA GeForce GTX 750                         |
| Network   | Qualcomm Atheros Killer E220x Gigabit Ethernet |

# Operating System Requirements

Required operating system:

```text
Ubuntu 24.04.4 LTS
```

Recommended:

* Updated system packages
* Active security updates
* sudo-enabled administrative user
* SSH access enabled

# Required Software

The following software is required:

## Database

* PostgreSQL 16+
* PostgreSQL client tools

## System Administration

* OpenSSH Server
* Git
* curl
* wget
* unzip
* nano
* vim
* tree
* htop
* net-tools

## Development Tools

Future application development may require:

* Python
* Go
* Rust
* Node.js
* Docker

# User Requirements

The server requires:

## Linux Administrative User

A Linux user account with:

* sudo privileges
* SSH access
* Permission to manage project files

## PostgreSQL Administrative User

A PostgreSQL administrative account with:

* Database creation privileges
* Role management privileges
* Schema management privileges

Current PostgreSQL administrative account:

```text
jonathon_admin
```

# Network Requirements

The environment requires:

* Local network connectivity
* Static or reserved IP address recommended
* SSH access on port 22
* Database access configured according to security requirements

# Storage Requirements

The server should provide enough storage for:

* Operating system files
* PostgreSQL databases
* Database backups
* Application files
* Logs
* Future monitoring data

# Security Requirements

The environment should include:

* Separate Linux and PostgreSQL accounts
* Role-based database permissions
* Protected credentials
* Regular backups
* System updates
* SSH security configuration

# Backup Requirements

The environment requires:

* Automated database backups
* Tested restore procedures
* Backup storage location

Current backup method:

```text
pg_dump custom format
```

Backup scripts:

```text
scripts/
├── backup_database.sh
└── restore_database.sh
```

# Future Requirements

As the project expands, additional requirements may include:

* Container runtime
* Monitoring stack
* Web application dependencies
* Infrastructure-as-Code tooling
* Additional storage devices

