# PostgreSQL Setup Documentation

## Database Overview

Server:
- U-Server

Database System:
- PostgreSQL 16.14

Database Purpose:
- SQL development environment
- Server inventory management
- Database administration practice
- Future application backend support

# PostgreSQL Installation

PostgreSQL was installed on Ubuntu 24.04.4 LTS to provide a relational database environment.

Installation was verified after setup.

Command:

```bash
psql --version
```

# Database Structure

Database:
- ubuntu_sql_server

Database User:
- jonathon_admin

Schema:
- server_management

Current Tables:
- hardware_components
- server_inventory
- operating_systems

## Current Database Purpose

The database currently stores server hardware and system inventory information.

Current tracked information:
- Hostname
- IP address
- Operating system
- CPU
- RAM
- Storage
- GPU
- Motherboard

# Hardware Components Table

The hardware_components table stores individual hardware components associated with a server.

Table Relationship:

server_inventory:
- Stores server identity and network information

hardware_components:
- Stores physical hardware installed in the server


Current Hardware Records:

Server:
- U-Server

Components:
- CPU
  - Intel Core i5-4690K
  - 4 Cores / 4 Threads
  - 3.50 GHz Base Frequency

- RAM
  - G.Skill F3-12800CL10-8GBXL
  - 16 GB DDR3
  - 1600 MT/s

- Motherboard
  - MSI Z97 GAMING 5 (MS-7917)
  - Intel Z97 Chipset

- GPU
  - NVIDIA GeForce GTX 750
  - GM107

# Operating Systems Table

The operating_systems table stores operating system information separately from the main server inventory.

Table Relationship:

server_inventory:
- Stores the server identity

operating_systems:
- Stores operating system details associated with the server


Current Operating System Record:

Server:
- U-Server

Operating System:
- Ubuntu 24.04.4 LTS

Kernel:
- Linux 6.17.0-35-generic

Architecture:
- x86_64

## Future Database Expansion

Planned tables:

- network_interfaces
- maintenance_logs
