# Database

This directory contains database-related files and documentation for the Ubuntu SQL Server project.

PostgreSQL is the main database system used to store server information collected by the project.

## Database Structure

Main database:

`ubuntu_sql_server`

Main schema:

`server_management`

The database stores information such as:

- Server inventory
- Operating system information
- Hardware components
- Network interfaces
- Health checks
- Maintenance logs
- Alert events

## Access

Applications such as the Go Agent, Django Web Dashboard, and project clients use PostgreSQL to store or retrieve server management data.

Database credentials and other secrets should not be committed to Git.

## Directories

- `backups/` - Database backup documentation and files
- `migrations/` - Database migration documentation and scripts

## Status

The database structure continues to grow as new server management features are added.
