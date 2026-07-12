#!/bin/bash

# PostgreSQL Database Restore Script
# Server: U-Server
# Database: ubuntu_sql_server_test

DATABASE="ubuntu_sql_server_test"
USER="jonathon_admin"
BACKUP_FILE="$HOME/postgres-backups/ubuntu_sql_server_backup.dump"

echo "Starting PostgreSQL restore..."

if [ ! -f "$BACKUP_FILE" ]; then
    echo "Backup file not found:"
    echo "$BACKUP_FILE"
    exit 1
fi

echo "Dropping existing test database..."

psql -U "$USER" -d postgres -c "DROP DATABASE IF EXISTS $DATABASE;"

echo "Creating test database..."

psql -U "$USER" -d postgres -c "CREATE DATABASE $DATABASE;"

echo "Restoring backup..."

pg_restore -U "$USER" -d "$DATABASE" "$BACKUP_FILE"

if [ $? -eq 0 ]; then
    echo "Restore completed successfully."
else
    echo "Restore failed."
    exit 1
fi