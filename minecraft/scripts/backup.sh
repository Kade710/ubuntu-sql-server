#!/bin/bash

set -e

MINECRAFT_DIR="$HOME/Projects/ubuntu-sql-server/minecraft"
DATA_DIR="$MINECRAFT_DIR/data"
BACKUP_DIR="$MINECRAFT_DIR/backups"

TIMESTAMP=$(date +"=%Y-%m-%d_%H-%M-%S")
BACKUP_FILE="$BACKUP_DIR/minecraft_$TIMESTAMP.tar.gz"

mkdir -p "$BACKUP_DIR"

echo "Starting Minecraft backup...."

# Make sure world saving is enabled again even if the script fails.
cleanup() {
    docker exec minecraft-server rcon-cli save-on >/dev/null 2>&1 || true
}

trap cleanup EXIT

echo "Flushing world data...."
docker exec minecraft-server rcon-cli save-off
docker exec minecraft-server rcon-cli save-all flush

echo "Creating backup:"
echo "$BACKUP_FILE"

tar -czf "$BACKUP_FILE" \
    -C "$DATA_DIR" \
    .

docker exec minecraft-server rcon-cli save-on

trap - EXIT

echo "Backup completed successfully."
echo "$BACKUP_FILE"
