# Software Inventory

## Overview

This document lists the software currently installed and configured on U-Server.

The inventory is maintained to provide visibility into the server environment and assist with troubleshooting, maintenance, and future rebuilds.

# Operating System

## Ubuntu

Distribution:

```text
Ubuntu 24.04.4 LTS Desktop
```

Purpose:

* Linux server development environment
* PostgreSQL hosting
* Automation development
* Future application hosting

# Database Software

## PostgreSQL

Software:

```text
PostgreSQL 16.14
```

Purpose:

* Relational database management
* SQL development
* Server inventory storage

Database:

```text
ubuntu_sql_server
```

Primary schema:

```text
server_management
```

Installed PostgreSQL components:

* PostgreSQL Server
* PostgreSQL Client Tools

# Remote Administration

## OpenSSH Server

Service:

```text
OpenSSH Server
```

Purpose:

* Remote administration
* Secure shell access
* VS Code remote development

Service status:

```text
Active
```

# Version Control

## Git

Version:

```text
Git 2.43.0
```

Purpose:

* Source control
* Repository management
* Project version tracking

# Development Tools

## Visual Studio Code

Purpose:

* Code editing
* Repository management
* Remote development

## Text Editors

Installed:

* Nano
* Vim

# System Administration Tools

## Monitoring and Diagnostics

Installed:

* htop
* tree
* net-tools

Purpose:

### htop

* Process monitoring
* Resource usage monitoring

### tree

* Directory structure visualization

### net-tools

* Network troubleshooting tools

# Network Utilities

Installed:

* curl
* wget

Purpose:

* Downloading resources
* API testing
* Network troubleshooting

# Archive Utilities

Installed:

* unzip

Purpose:

* Extracting compressed files
* Managing downloaded archives

# Database Management Tools

Installed:

* psql
* pg_dump
* pg_restore

Purpose:

* Database administration
* Backup creation
* Database recovery testing

# Automation Tools

## Custom Scripts

Location:

```text
scripts/
```

Current automation scripts:

| Script                  | Purpose                                        |
| ----------------------- | ---------------------------------------------- |
| install_dependencies.sh | Installs required system dependencies          |
| backup_database.sh      | Creates PostgreSQL database backups            |
| restore_database.sh     | Restores PostgreSQL backups                    |
| system_update.sh        | Performs system maintenance checks and updates |

# Installed Software Summary

| Category         | Software                           |
| ---------------- | ---------------------------------- |
| Operating System | Ubuntu 24.04.4 LTS                 |
| Database         | PostgreSQL 16.14                   |
| Version Control  | Git 2.43.0                         |
| Remote Access    | OpenSSH Server                     |
| Editors          | Nano, Vim, VS Code                 |
| Monitoring       | htop                               |
| Utilities        | curl, wget, unzip, tree, net-tools |
| Database Tools   | psql, pg_dump, pg_restore          |

# Inventory Maintenance

This document should be updated when:

* New software is installed
* Software versions change
* Services are added or removed
* Server responsibilities change

