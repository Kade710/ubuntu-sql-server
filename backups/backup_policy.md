# Backup Policy

## Purpose

The purpose of this policy is to provide a basic plan for protecting important Ubuntu SQL Server data.

## Backup Goals

Backups should:

- Protect important project and server data
- Be created on a regular schedule when needed
- Be stored separately from active application data
- Avoid storing passwords or other secrets when possible
- Be tested to make sure they can be restored

## Backup Types

Depending on the service, backups may include:

- PostgreSQL database backups
- Application data backups
- Configuration backups
- Minecraft world and server data

## Retention

Old backups should be removed when they are no longer needed to prevent unnecessary storage use.

Retention periods may be different depending on the service and importance of the data.

## Restore Testing

Backups should be tested from time to time to make sure the data can be restored successfully.

A backup should not be considered reliable until the restore process has been tested.

## Security

Sensitive information such as passwords, API keys, and environment files should not be included in backups unless they are required for recovery and properly protected.

Backups containing sensitive information should not be committed to Git.
