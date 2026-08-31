# Database Design Documentation

## Overview

The Ubuntu SQL Server project uses PostgreSQL as the primary database engine.

The database was designed to manage server infrastructure information including:

- Server inventory
- Hardware components
- Operating system information
- Network interfaces
- Maintenance history

Database Name:
ubuntu_sql_server

Schema:
server_management

---

# Database Architecture

The database uses a relational model with a central server inventory table.

The primary table is:
server_inventory


Additional information is separated into related tables:
server_inventory
|
|
+-- hardware_components

    +-- operating_systems

    +-- network_interfaces

    +-- maintenance_logs

---

# Design Goals

The database design focuses on:

- Avoiding duplicate information
- Maintaining relationships between records
- Supporting future application development
- Allowing API and dashboard integration
- Maintaining accurate infrastructure records

---

# Tables

## server_inventory

Stores the primary information about managed servers.

Contains:

- Hostname
- IP address
- Operating system
- CPU
- RAM
- Storage
- GPU
- Motherboard

---

## hardware_components

Stores additional hardware details associated with a server.

Relationship:


hardware_components.server_id
|
v
server_inventory.id

---

## operating_systems

Stores detailed operating system information.

Includes:

- Distribution
- Version
- Kernel
- Architecture

---

## network_interfaces

Stores network adapter information.

Includes:

- Interface name
- MAC address
- IP address
- Network type
- Speed

---

## maintenance_logs

Stores maintenance and administrative actions.

Includes:

- Action performed
- Description
- Administrator
- Timestamp

---

# Future Development

The database will support:

- Python database clients
- API development
- Web dashboard reporting
- Infrastructure automation
- Monitoring integrations
