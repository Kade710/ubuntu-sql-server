#!/bin/bash

set -e

MINECRAFT_DIR="$HOME/Projects/ubuntu-sql-server/minecraft"
DATA_DIR="$MINECRAFT_DIR/data"
BACKUP_DIR="$MINECRAFT_DIR/backups"

if [ -z "$1" ]; then
    echo "Usage:"
    echo "./scripts/restore.sh <backup-file>"
    echo
    echo "Available backups:"
    ls -lh "$BACKUP_DIR"/*.tar.gz 2>/dev/null || echo "No backups found."
    exit 1
fi

BACKUP_FILE="$1"

# Allow just the filename to be supplied.
if [[ "$BACKUP_FILE" != /* ]]; then
    BACKUP_FILE="$BACKUP_DIR/$BACKUP_FILE"
fi

if [ ! -f "$BACKUP_FILE" ]; then
    echo "Backup not found:"
    echo "$BACKUP_FILE"
    exit 1
fi

echo "Backup selected:"
echo "$BACKUP_FILE"
echo
echo "WARNING: This will replace the current Minecraft data."
read -r -p "Type RESTORE to continue: " CONFIRM

if [ "$CONFIRM" != "RESTORE" ]; then
    echo "Restore cancelled."
    exit 1
fi

echo "Checking backup archive..."

tar -tzf "$BACKUP_FILE" >/dev/null

echo "Stopping Minecraft server..."

cd "$MINECRAFT_DIR"
docker compose down

SAFETY_DIR="$MINECRAFT_DIR/pre-restore-data-$(date +"%Y-%m-%d_%H-%M-%S")"

echo "Saving current data to:"
echo "$SAFETY_DIR"

mv "$DATA_DIR" "$SAFETY_DIR"
mkdir -p "$DATA_DIR"

restore_failed() {
    echo "Restore failed."

    rm -rf "$DATA_DIR"

    if [ -d "$SAFETY_DIR" ]; then
        echo "Restoring previous Minecraft data..."
        mv "$SAFETY_DIR" "$DATA_DIR"
    fi

    docker compose up -d || true
}

trap restore_failed ERR

echo "Extracting backup..."

tar -xzf "$BACKUP_FILE" -C "$DATA_DIR"

echo "Starting Minecraft server..."

docker compose up -d

trap - ERR

echo
echo "Restore completed."
echo "Previous data remains at:"
echo "$SAFETY_DIR"
echo
echo "Check server health with:"
echo "docker ps"