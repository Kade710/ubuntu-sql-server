# Ansible Automation

## Overview

This directory contains planned Ansible automation for the Ubuntu SQL Server project.

The goal of this automation is to transition repeatable server configuration tasks from individual Bash scripts into a structured configuration management system.

Ansible will allow the server environment to be rebuilt consistently across multiple systems.

# Current Automation Status

The project currently uses Bash-based automation scripts located in:

```text
scripts/
```

Completed automation:

* `install_dependencies.sh`

  * Installs required software packages
  * Verifies required services

* `backup_database.sh`

  * Creates PostgreSQL database backups
  * Stores backup files in the configured backup directory

* `restore_database.sh`

  * Tests database restoration procedures
  * Validates backup recovery

* `system_update.sh`

  * Updates Ubuntu packages
  * Checks reboot requirements
  * Verifies system uptime
  * Checks PostgreSQL service status

# Future Ansible Migration

Future Ansible playbooks will replace manual and Bash-based configuration tasks.

Planned automation areas:

## System Configuration

* Ubuntu package installation
* System updates
* User configuration
* Directory creation

## Database Deployment

* PostgreSQL installation
* Database creation
* Database roles
* Schema deployment
* Permission configuration

## Security Configuration

* SSH configuration
* Firewall configuration
* User permissions
* Security baseline configuration

## Monitoring Configuration

* System monitoring tools
* Database health checks
* Service monitoring
* Alert configuration

# Planned Directory Structure

Future structure:

```text
ansible/

├── inventory/
│
├── playbooks/
│
├── roles/
│
└── README.md
```

# Current Status

Ansible automation has not been implemented yet.

Current phase:

* Bash automation completed
* Documentation in progress
* Ansible migration planned for future development

# Purpose

The purpose of this directory is to document the future transition toward Infrastructure as Code (IaC) practices.

The completed Bash automation provides the foundation that will later be converted into reusable Ansible playbooks and roles.

