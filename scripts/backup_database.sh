#!/bin/bash

# PostgreSQL Database Backup Script
# Server: U-Server
# Database: ubuntu_sql_server

DATABASE="ubuntu_sql_server"
USER="jonathon_admin"
BACKUP_DIR="$HOME/postgres-backups"
BACKUP_FILE="$BACKUP_DIR/ubuntu_sql_server_backup.dump"

echo "Starting PostgreSQL backup..."

mkdir -p "$BACKUP_DIR"

pg_dump -U "$USER" -d "$DATABASE" -F c -f "$BACKUP_FILE"

if [ $? -eq 0 ]; then
    echo "Backup completed successfully."
    echo "Backup location: $Backup_FILE"
else
    echo "Backup failed."
    exit 1
fi