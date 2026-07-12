#!/bin/bash

# Ubuntu System Update Script
#Server: U-Server

echo "Starting system update..."

echo "Updating package lists..."

sudo apt update

echo "Upgrading Installed packages..."

sudo apt upgrade -y

echo "Checking reboot requirement..."

if command -v needs-restarting >/dev/null 2>&1; then
    needs-restarting -r
else
    echo "needs-restarting command not installed."
fi

echo "Checking system uptime..."

uptime

echo "Checking PostgreSQL service..."

systemctl status postgresql@16-main --no-pager

echo "System maintenance check completed."