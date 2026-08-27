# Database Migrations

This directory is used to document and organize changes to the PostgreSQL database structure.

Migrations help track changes such as:

- Creating tables
- Adding or removing columns
- Creating indexes
- Adding constraints
- Updating relationships between tables
- Adding new database features

## Purpose

Database changes should be documented so the structure of the `ubuntu_sql_server` database can be recreated and maintained over time.

Migration files should be added when database changes need to be saved as part of the project.

## Guidelines

- Keep migrations organized and easy to understand.
- Avoid storing passwords or other secrets in migration files.
- Test important database changes before relying on them.
- Back up important data before major changes.
- Do not modify old migrations without a clear reason.

## Status

Migration documentation and scripts will be added as database changes are formalized.
