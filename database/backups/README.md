# Database Backups

This directory is used for PostgreSQL backup and recovery resources.

Database backups help protect server management data from accidental deletion, database problems, or system failure.

## Backup Goals

Database backups should:

- Protect important PostgreSQL data
- Be created before major database changes when needed
- Be stored safely
- Avoid exposing database credentials
- Be tested to make sure they can be restored

## Backup Files

Database backup files should not be committed to Git unless they contain safe test data.

Common PostgreSQL backup files may include:

- `.dump`
- `.backup`
- `.sql.backup`

## Restore Testing

Backups should be tested when possible to make sure the database can be recovered successfully.

## Status

Database backup procedures will continue to be improved as the project grows.
