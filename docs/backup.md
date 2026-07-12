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
```

# Recovery Testing

## Restore Test Environment

A temporary database was created to verify backup recovery.

Test Database:

```text
ubuntu_sql_server_test
```

# Automated Database Backup

## Backup Script

A shell script was created to automate PostgreSQL database backups.

Script:

```text
scripts/backup_database.sh
```

# Database Restore Testing

## Restore Script

A restore script was created to automate PostgreSQL database recovery testing.

Script:

```text
scripts/restore_database.sh
```

Purpose:

* Restore PostgreSQL database backups
* Validate backup integrity
* Provide a repeatable recovery procedure

## Restore Process

The restore script performs the following actions:

1. Checks that the backup file exists.
2. Removes the existing test database.
3. Creates a new test database.
4. Restores the PostgreSQL dump file.
5. Confirms successful completion.

## Running the Restore Test

Command:

```bash
./scripts/restore_database.sh
```

## Verification

Database restoration was verified by connecting to the restored database:

```bash
psql -U jonathon_admin -d ubuntu_sql_server_test
```

Tables verified:

```sql
\dt server_management.*
```

## Recovery Validation

The following data was successfully restored:

* Server inventory records
* Hardware components
* Operating system information
* Network interfaces
* Maintenance logs

Result:

* Backup successfully restored
* Database structure preserved
* Data integrity confirmed
