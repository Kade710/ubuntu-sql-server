# Security Documentation

## Security Overview

Server:
- U-Server

Operating System:
- Ubuntu 24.04.4 LTS Desktop

Security Goals:

- Maintain controlled user access
- Protect database credentials
- Enable secure remote administration
- Follow basic Linux security practices


# User Management

Linux and PostgreSQL access are separated using different user accounts.

## Linux User

Primary Linux user:

```text
jonathon
```

# PostgreSQL Security Configuration

## Role-Based Access Control

PostgreSQL was configured using separate accounts with different privilege levels.

## postgres

Purpose:
- PostgreSQL system administration

Privileges:
- Superuser
- Create roles
- Create databases


## jonathon_admin

Purpose:
- Database administration and development

Privileges:
- Create roles
- Create databases
- Manage database schemas and tables


## db_readonly

Purpose:
- Reporting and read-only database access

Privileges:
- Connect to ubuntu_sql_server
- Access server_management schema
- SELECT access on inventory tables

Restrictions:
- Cannot INSERT
- Cannot UPDATE
- Cannot DELETE


# Permission Verification

Read Test:

Command:

```sql
SELECT * FROM server_management.server_inventory;
```