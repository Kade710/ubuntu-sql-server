# Backup Documentation

## Backup Overview

Server:
- U-Server

Database:
- ubuntu_sql_server

Backup Tool:
- pg_dump

Database Version:
- PostgreSQL 16.14


# PostgreSQL Backup Process

PostgreSQL backups are created using the pg_dump utility.

Backup location:

```text
~/postgres-backups/


# PostgreSQL Backup Documentation

## Backup Command

```bash
pg_dump -U jonathon_admin -d ubuntu_sql_server -F c -f ~/postgres-backups/ubuntu_sql_server_backup.dump

