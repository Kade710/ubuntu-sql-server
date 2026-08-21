#!/bin/bash

set -e

MINECRAFT_DIR="$HOME/Projects/ubuntu-sql-server/minecraft"
DATA_DIR="$MINECRAFT_DIR/data"
BACKUP_DIR="$MINECRAFT_DIR/backups"

MAX_BACKUPS=7

TIMESTAMP=$(date +"%Y-%m-%d_%H-%M-%S")
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
    --exclude='./.rcon-cli.env' \
    --exclude='./.rcon-cli.yaml' \
    -C "$DATA_DIR" \
    .

docker exec minecraft-server rcon-cli save-on

trap - EXIT

echo "Backup completed successfully."
echo "$BACKUP_FILE"

echo "Cleaning old backups...."

BACKUP_COUNT=$(find "$BACKUP_DIR" -maxdepth 1 -type f -name 'minecraft_*.tar.gz' | wc -l)

if [ "$BACKUP_COUNT" -gt "$MAX_BACKUPS" ]; then
    find "$BACKUP_DIR" -maxdepth 1 -type f -name 'minecraft_*.tar.gz' \
        -printf '%T@ %p\n' \
        | sort -n \
        | head -n "$((BACKUP_COUNT - MAX-BACKUPS)) \
        | cut -d' ' -f2- \
        | xargs -r rm --
fi
