# Recovery Plan

## Purpose

This document provides a basic recovery plan for the Ubuntu SQL Server project.

The goal is to restore important services and data as safely as possible after a system failure.

## Recovery Priorities

Recovery should focus on restoring services in a logical order.

1. Restore the operating system and network access.
2. Restore SSH access.
3. Restore PostgreSQL.
4. Restore project files and application code.
5. Restore the Go Agent and monitoring services.
6. Restore the Django Web Dashboard.
7. Restore Docker services.
8. Restore service data from backups when needed.

## Database Recovery

PostgreSQL should be checked before applications that depend on it are started.

Important steps may include:

- Verify the PostgreSQL service is running.
- Verify the `ubuntu_sql_server` database is available.
- Verify the `server_management` schema is present.
- Restore from a database backup if required.
- Test application access after recovery.

## Application Recovery

Applications should be restored from the Git repository when possible.

Environment files and credentials must be restored separately because they should not be stored in Git.

## Docker Recovery

Docker services can be recreated using their configuration files.

Persistent data should be restored from backups if the original data is damaged or lost.

## Minecraft Recovery

Minecraft server data can be restored using the project backup and restore scripts.

A restore should be tested before older backup copies are removed.

## Validation

After recovery, verify:

- PostgreSQL is online.
- The Go Agent is running.
- Health checks are being stored.
- The Web Dashboard is reachable.
- Notifications are working.
- Docker containers are healthy.
- Important application data is available.

## Review

This recovery plan should be updated whenever major services, backup methods, or infrastructure changes are added.
