# Build Journal

## Overview

This journal tracks major development work completed on the Ubuntu SQL Server project.

It is meant to provide a simple history of how the project has grown, what was tested, and what still needs work.

---

## Server Environment

U-Server was configured as the main development and hosting system for the project.

The environment includes:

- Ubuntu
- PostgreSQL
- Git
- Go
- Python
- Django
- Rust
- Docker
- SSH
- Tailscale

The project is stored under:

```text
~/Projects/ubuntu-sql-server
```

---

## PostgreSQL Database

The `ubuntu_sql_server` PostgreSQL database was created to store server management information.

The main schema is:

```text
server_management
```

Current database areas include:

- Server inventory
- Operating system information
- Hardware components
- Network interfaces
- Health checks
- Maintenance logs
- Alert events

Foreign keys are used to connect related records back to the main server inventory.

---

## Go Agent

A Go-based monitoring agent was developed to collect information directly from U-Server.

The agent can collect:

- Server inventory
- Operating system information
- CPU information
- Memory information
- Motherboard information
- GPU information
- Network interfaces
- System health

A full refresh can be performed with:

```bash
./agent --refresh
```

The refresh process updates the main server information and records a new health check.

---

## Health Monitoring

System health monitoring was added to the Go Agent.

Current measurements include:

- 1-minute load average
- Memory usage
- Disk usage
- Uptime

The monitoring system classifies server health as:

```text
HEALTHY
WARNING
CRITICAL
```

Health checks are stored in PostgreSQL so previous server conditions can be reviewed.

---

## Health Alerts

Status-change detection was added to the Go Agent.

The agent compares the newest health status with the previous health status.

This allows it to detect changes such as:

```text
HEALTHY -> WARNING
WARNING -> CRITICAL
WARNING -> HEALTHY
CRITICAL -> HEALTHY
```

The system can also recognize when U-Server recovers from a warning or critical condition.

---

## Alert Event History

The following table was added to PostgreSQL:

```text
server_management.alert_events
```

Alert events store:

- Server ID
- Previous status
- New status
- Alert title
- Alert message
- Creation time

This provides a database history of important health changes.

Alert event storage was tested successfully with a recovery event.

---

## Notifications

ntfy notification support was added to the Go Agent.

The monitoring system can send notifications for:

- Warning conditions
- Critical conditions
- Server recovery

Notification configuration is stored outside of the source code using environment variables.

---

## systemd Integration

The Go Agent was integrated with systemd.

Service:

```text
ubuntu-sql-agent.service
```

Agent activity can be reviewed with:

```bash
journalctl -u ubuntu-sql-agent.service
```

The agent performs its monitoring work and then exits successfully.

A completed run may appear as:

```text
ubuntu-sql-agent.service: Deactivated successfully.
Finished ubuntu-sql-agent.service - Ubuntu SQL Server Go Agent.
```

This is expected behavior.

---

## Django Web Dashboard

A Django Web Dashboard was added to provide browser-based access to server information stored in PostgreSQL.

The dashboard is being developed to provide easier access to:

- Server information
- Hardware information
- Network information
- Health history
- Maintenance information

Database and network connectivity have also been tested during dashboard development.

---

## Client Applications

Additional clients have been developed for interacting with project data.

Current client areas include:

- Python client
- Rust client

These applications provide additional ways to access server management information without working directly in PostgreSQL.

---

## Docker

Docker was added to U-Server to support containerized services.

Docker networking creates virtual interfaces that may also be detected by the Go Agent.

Container status can be checked with:

```bash
docker ps
```

---

## Minecraft Server

A Minecraft server was deployed using Docker.

The server uses:

```text
minecraft-server
```

as its container name.

The Minecraft service is exposed through port:

```text
25565
```

RCON is used for server administration and backup operations.

---

## Minecraft Backup System

Automated Minecraft backups were developed and tested.

The backup process:

1. Temporarily disables automatic saving.
2. Flushes world data to disk.
3. Creates a compressed backup archive.
4. Re-enables automatic saving.
5. Reports whether the backup completed successfully.

Backups are scheduled using systemd.

The backup timer is:

```text
minecraft-backup.timer
```

The backup service is:

```text
minecraft-backup.service
```

---

## Minecraft Restore System

A restore script was created and tested.

Before replacing the active Minecraft data, the restore process:

1. Requires confirmation.
2. Checks the selected backup.
3. Stops the Minecraft container.
4. Preserves the current data.
5. Extracts the selected backup.
6. Restarts the Minecraft server.

A full restore test was completed successfully and the Minecraft server returned online afterward.

---

## Networking

U-Server currently uses several physical and virtual network interfaces.

These include:

```text
enp3s0
tailscale0
docker0
```

Additional Docker bridge interfaces may appear while containers are running.

Networking is used for:

- SSH
- PostgreSQL
- Web Dashboard access
- Docker
- Minecraft
- Tailscale

---

## Firewall

UFW is used to control incoming connections.

Firewall configuration should be checked when a service is running but cannot be reached.

Current rules can be viewed with:

```bash
sudo ufw status numbered
```

This became an important troubleshooting step while testing project services.

---

## Logging

systemd, PostgreSQL, Docker, and application logs are used for troubleshooting and monitoring.

Go Agent logs can be viewed with:

```bash
journalctl -u ubuntu-sql-agent.service
```

Minecraft logs can be viewed with:

```bash
docker logs minecraft-server
```

Additional logging information is documented in:

```text
logs/README.md
```

---

## Documentation

Project documentation has been expanded to cover areas including:

- API
- Applications
- Architecture
- Automation
- Backups
- Compliance
- Configuration
- Database
- Database design
- Disaster recovery
- Environment variables
- Licensing and attribution
- Logging
- Monitoring
- Networking
- Troubleshooting

Documentation will continue to be updated as features are completed.

---

## Current Development

The project is still actively being developed.

Current work includes improving:

- Health monitoring
- Alert history
- Web Dashboard features
- Application integration
- Automation
- Documentation

New entries should be added to this journal when major features are built, tested, or changed.
