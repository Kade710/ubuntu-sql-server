# Ubuntu Installation Documentation

## Installation Overview

Server Name:

* U-Server

Operating System:

* Ubuntu 24.04.4 LTS Desktop

Installation Type:

* Clean installation

Purpose:

* Linux server development environment
* PostgreSQL database hosting
* SQL development practice
* Linux administration practice
* Future application hosting

# Hardware Environment

The operating system was installed on custom-built hardware.

Hardware Summary:

* Manufacturer: MSI
* Motherboard: MSI Z97 GAMING 5 (MS-7917)
* Processor: Intel Core i5-4690K @ 3.50GHz
* Memory: 16GB DDR3
* Storage: 2TB drive
* Graphics: NVIDIA GeForce GTX 750
* Network: Qualcomm Atheros Killer E220x Gigabit Ethernet

# Operating System Installation

Ubuntu 24.04.4 LTS Desktop was installed to provide a graphical Linux environment while developing and configuring the server.

The desktop environment allows:

* Local administration
* Software development
* VS Code usage
* Hardware troubleshooting
* Database management

The system will transition toward headless server management using SSH.

# Initial System Configuration

After installation, the system was updated to ensure all packages were current.

Commands used:

```bash
sudo apt update
sudo apt upgrade
```

System updates were verified before continuing with server configuration.

# Dependency Installation

Required software packages were installed using an automated dependency script.

Installation script:

```bash
./scripts/install_dependencies.sh
```

The script installs and verifies required server administration and development tools.

Installed software:

* PostgreSQL
* PostgreSQL Client
* Git
* OpenSSH Server
* curl
* wget
* unzip
* nano
* vim
* tree
* htop
* net-tools

Verification commands:

```bash
git --version

psql --version

systemctl status ssh
```

The dependency installation completed successfully.

# PostgreSQL Installation and Configuration

PostgreSQL was installed as the primary database management system.

Installed Version:

```text
PostgreSQL 16.14
```

Service verification:

```bash
systemctl status postgresql@16-main
```

PostgreSQL service status:

* Active
* Running
* Enabled for server operation

# Database Configuration

The primary database was created:

```text
ubuntu_sql_server
```

Database purpose:

* Server inventory management
* Hardware tracking
* Network information storage
* Maintenance logging

The database uses the following schema:

```text
server_management
```

Database tables:

```text
server_inventory
hardware_components
network_interfaces
operating_systems
maintenance_logs
```

# PostgreSQL User Roles

PostgreSQL access was separated using role-based access control.

## postgres

Purpose:

* PostgreSQL system administration

Privileges:

* Superuser
* Create roles
* Create databases

## jonathon_admin

Purpose:

* Database administration
* Development account

Privileges:

* Create roles
* Create databases
* Manage schemas
* Manage tables
* Perform database development tasks

## db_readonly

Purpose:

* Reporting access
* Read-only database usage

Privileges:

* Connect to ubuntu_sql_server
* Access server_management schema
* SELECT permissions on inventory tables

Restrictions:

* Cannot INSERT
* Cannot UPDATE
* Cannot DELETE

# SSH Configuration

SSH was enabled for remote server administration.

Service verification:

```bash
systemctl status ssh
```

Remote administration was tested using:

```bash
ssh jonathon@server_ip
```

Supported administration methods:

* Windows PowerShell SSH
* Visual Studio Code Remote SSH

# Database Backup Configuration

Database backups were automated using:

```bash
./scripts/backup_database.sh
```

Backup format:

```text
PostgreSQL custom format dump
```

Backup location:

```text
~/postgres-backups/ubuntu_sql_server_backup.dump
```

Backup verification was completed using:

```bash
pg_restore -l ~/postgres-backups/ubuntu_sql_server_backup.dump
```

The backup contained:

* Database schema
* Tables
* Data
* Constraints
* Ownership information

# Database Restore Testing

Database restoration was tested using:

```bash
./scripts/restore_database.sh
```

The restore process:

1. Removed the test database
2. Created a clean test database
3. Restored the PostgreSQL backup
4. Verified database tables
5. Verified restored data

Test database:

```text
ubuntu_sql_server_test
```

Validation command:

```sql
SELECT * FROM server_management.server_inventory;
```

Restore testing confirmed:

* Schema restoration
* Table restoration
* Data restoration
* Primary keys
* Foreign keys
* Database ownership

# System Maintenance Automation

Routine system maintenance was automated using:

```bash
./scripts/system_update.sh
```

The script performs:

* Ubuntu package updates
* Installed package upgrades
* Reboot requirement checks
* System uptime verification
* PostgreSQL service verification

Maintenance results are recorded through:

```text
maintenance/maintenance_log.md
```

# Current Installation Status

Completed:

* [x] Ubuntu 24.04.4 LTS installation
* [x] Hardware identification
* [x] System updates
* [x] Dependency installation
* [x] PostgreSQL installation
* [x] Database creation
* [x] Schema deployment
* [x] User and role configuration
* [x] SSH configuration
* [x] Backup automation
* [x] Restore testing
* [x] System maintenance automation

# Future Installation Additions

Planned future configuration:

* Monitoring services
* Ansible automation
* Application hosting
* API deployment
* Web dashboard
* Additional security hardening
