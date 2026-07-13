#!/bin/bash

# ==========================================================
# Ubuntu SQL Server Project
# Dependency Installation Script
# Author: Jonathon Anderson
# ==========================================================

set -e

echo "========================================="
echo "Ubuntu SQL Server Dependency Installer"
echo "========================================="

echo
echo "[1/4] Updating package lists..."
sudo apt update

echo
echo "[2/4] Installing required packages..."

sudo apt install -y \
    postgresql \
    postgresql-client \
    git \
    curl \
    wget \
    unzip \
    nano \
    vim \
    tree \
    htop \
    net-tools \
    openssh-server

echo
echo "[3/4] Verifying installed software..."

echo
echo "Git:"
git --version

echo
echo "PostgreSQL:"
psql --version

echo
echo "SSH:"
systemctl status ssh --no-pager

echo
echo "[4/4] Dependency installation completed."

echo
echo "Installed packages:"
echo " - PostgreSQL"
echo " - PostgreSQL Client"
echo " - Git"
echo " - curl"
echo " - wget"
echo " - unzip"
echo " - nano"
echo " - vim"
echo " - tree"
echo " - htop"
echo " - net-tools"
echo " - OpenSSH Server"

echo
echo "Done."
