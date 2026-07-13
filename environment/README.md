# Environment Overview

## Purpose

This directory documents the current operating environment for the Ubuntu SQL Server project.

The environment documentation provides an overview of the server hardware, operating system, software requirements, and installed software used to support database development, Linux administration, and future application hosting.

# Server Information

## Server Name

```text
U-Server
```

## Operating System

```text
Ubuntu 24.04.4 LTS Desktop
```

## Environment Type

```text
Development Server
```

## Primary Purpose

U-Server is used as a self-hosted Linux development environment for:

* PostgreSQL database administration
* SQL development and testing
* Linux system administration practice
* Server automation development
* Future application hosting

# Hardware Environment

The server is built using custom desktop hardware.

Hardware summary:

| Component   | Specification                                  |
| ----------- | ---------------------------------------------- |
| Motherboard | MSI Z97 GAMING 5 (MS-7917)                     |
| Processor   | Intel Core i5-4690K @ 3.50GHz                  |
| Memory      | 16GB DDR3 RAM                                  |
| Storage     | 2TB drive                                      |
| Graphics    | NVIDIA GeForce GTX 750                         |
| Network     | Qualcomm Atheros Killer E220x Gigabit Ethernet |

# Network Environment

Primary network interface:

```text
enp3s0
```

Current server address:

```text
192.168.1.100
```

Network usage:

* Local network administration
* SSH remote access
* Database development access

# Core Services

The environment currently provides the following services:

## PostgreSQL Database Server

Purpose:

* Database storage
* SQL development
* Server inventory management

Version:

```text
PostgreSQL 16.14
```

Database:

```text
ubuntu_sql_server
```

Primary schema:

```text
server_management
```

## SSH Remote Administration

Purpose:

* Remote server management
* Development access
* VS Code remote administration

Service:

```text
OpenSSH Server
```

# Development Environment

Installed development and administration tools include:

* Git
* Visual Studio Code
* Nano
* Vim
* Tree
* HTOP
* curl
* wget
* PostgreSQL client tools

# Project Environment Structure

Related documentation:

```text
environment/
├── README.md
├── requirements.md
└── software_inventory.md
```

Additional configuration:

```text
environment_variables/
├── .env.example
└── README.md
```

# Current Environment Status

Completed:

* [x] Ubuntu server installation
* [x] Hardware identification
* [x] PostgreSQL deployment
* [x] Database configuration
* [x] SSH configuration
* [x] Dependency installation
* [x] Backup and restore testing
* [x] System maintenance automation

# Future Environment Expansion

Planned additions:

* Application hosting environment
* API deployment environment
* Monitoring services
* Containerization
* Infrastructure automation using Ansible

