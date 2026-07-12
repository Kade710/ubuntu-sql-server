# Automation Scripts

## Overview

This directory contains automation scripts used to maintain and operate the Ubuntu SQL Server environment.

The scripts automate common administrative tasks including:

- Installing required software dependencies
- Performing system maintenance
- Creating PostgreSQL database backups
- Restoring PostgreSQL backups

Server:

- U-Server

Operating System:

- Ubuntu 24.04.4 LTS


---

# Available Scripts


## install_dependencies.sh

Purpose:

Installs required software packages needed for the server environment.


Installed Components:

- PostgreSQL
- PostgreSQL Client
- Git
- curl
- wget
- unzip
- nano
- vim
- tree
- htop
- net-tools
- OpenSSH Server


Usage:

```bash
./scripts/install_dependencies.sh