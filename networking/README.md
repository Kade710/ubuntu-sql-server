# Networking

## Overview

This directory contains networking documentation for the Ubuntu SQL Server project.

U-Server uses physical and virtual network interfaces to provide local network access, remote administration, Docker networking, and access to hosted services.

## Current Networking

The primary physical network interface is:

```text
enp3s0
```

U-Server also uses virtual interfaces created by services such as:

- Tailscale
- Docker
- Docker Compose networks

## Network Services

Network services may include:

- SSH
- PostgreSQL
- Django Web Dashboard
- Minecraft Server
- Tailscale
- Docker services

Access to these services depends on the service configuration, listening address, and firewall rules.

## Firewall

UFW is used to control incoming network connections.

Current firewall rules can be viewed with:

```bash
sudo ufw status numbered
```

Only ports required by active services should be opened.

## Listening Ports

Listening TCP ports can be checked with:

```bash
sudo ss -lntp
```

A specific port can be checked with:

```bash
sudo ss -lntp | grep <port>
```

## Documentation

- `ip_addressing.md` - Documents IP addressing and network interfaces.
- `services.md` - Documents network services and ports.

## Security

Network services should only be exposed when required.

Database access, administrative services, and application ports should be limited to trusted networks whenever possible.

## Status

Networking documentation should be updated whenever interfaces, addresses, firewall rules, or hosted services change.
