# Changelog

All notable changes to the Ubuntu SQL Server project are documented in this file.

The project is under active development, so features and documentation may change as the server environment grows.

## Unreleased

### Added

#### Go Agent

- Added Go-based server monitoring agent.
- Added full server refresh command using:

```bash
./agent --refresh
```

- Added server inventory collection.
- Added operating system detection.
- Added hardware component detection.
- Added network interface detection.
- Added system health monitoring.
- Added PostgreSQL integration.
- Added maintenance log support.
- Added systemd integration.
- Added ntfy health notifications.
- Added health status transition detection.
- Added recovery notifications.
- Added alert event database logging.

#### Health Monitoring

- Added monitoring for:
  - 1-minute load average
  - Memory usage
  - Disk usage
  - System uptime

- Added health classifications:

```text
HEALTHY
WARNING
CRITICAL
```

- Added health history storage in PostgreSQL.
- Added comparison between current and previous health states.
- Added detection of warning, critical, and recovery conditions.

#### Alert Events

- Added:

```text
server_management.alert_events
```

- Alert events can record:
  - Server ID
  - Previous status
  - New status
  - Alert title
  - Alert message
  - Creation time

- Added foreign key relationship between alert events and server inventory.
- Successfully tested storage of a recovery event.

#### PostgreSQL

- Added the `ubuntu_sql_server` project database.
- Added the `server_management` schema.
- Added database storage for:
  - Server inventory
  - Operating systems
  - Hardware components
  - Network interfaces
  - Health checks
  - Maintenance logs
  - Alert events

- Added PostgreSQL application access.
- Added database role and permission documentation.

#### Django Web Dashboard

- Added Django-based Web Dashboard.
- Added browser-based access to server information.
- Added database integration with PostgreSQL.
- Added dashboard support for project monitoring information.
- Added Django migration structure.

#### Client Applications

- Added Python client.
- Added Rust client.
- Added Go Agent.

- Added client access to server management information stored in PostgreSQL.

#### API

- Added API project structure.
- Added API documentation structure.
- Added endpoint documentation.
- Added authentication documentation.

#### Docker

- Added Docker support to U-Server.
- Added Docker Compose support.
- Added container networking.
- Added Docker logging and troubleshooting documentation.
- Added container resource monitoring documentation.

#### Bittensor Development Environment

- Added Docker-based Bittensor development environment.
- Added Bittensor Compose configuration.
- Added NVIDIA CUDA base environment.
- Added `bittensor-node-dev` development container.

#### Minecraft

- Added Docker Compose configuration for Minecraft.
- Added Minecraft server image.
- Added persistent Minecraft project structure.
- Added backup scripts.
- Added restore scripts.
- Added systemd backup service.
- Added systemd backup timer.
- Added backup and restore documentation.

#### Networking

- Added documentation for:
  - IP addressing
  - Network interfaces
  - Network services
  - Docker networking
  - Tailscale
  - Firewall troubleshooting

- Added monitoring of physical and virtual network interfaces.

#### Security

- Added UFW firewall documentation.
- Added Linux permission documentation.
- Added SSH hardening documentation.
- Added database role documentation.
- Added secret-management guidelines.
- Added environment variable documentation.

#### Logging

- Added centralized project logging documentation.
- Added systemd journal documentation.
- Added Go Agent logging documentation.
- Added PostgreSQL logging documentation.
- Added Docker logging documentation.
- Added health and alert event logging documentation.

#### Testing

- Added connectivity testing documentation.
- Added SQL testing documentation.
- Added database verification procedures.
- Added network troubleshooting tests.
- Added Go Agent connectivity testing procedures.

#### Backups and Recovery

- Added backup policy documentation.
- Added database backup documentation.
- Added disaster recovery documentation.
- Added recovery planning.
- Added Minecraft backup and restore procedures.

#### Performance

- Added performance documentation structure.
- Added benchmark documentation.
- Added performance tuning documentation.
- Added guidelines for measuring performance before tuning.

#### Automation

- Added automation documentation.
- Added cron job documentation.
- Added Ansible project structure.
- Added systemd-based automation for monitoring and backups.

#### Environment

- Added environment documentation.
- Added software inventory documentation.
- Added requirements documentation.
- Added environment variable examples.
- Added `.env.example` support.

#### Architecture

- Added architecture documentation.
- Added system design documentation.
- Added future design documentation.
- Added diagram directory structure.

#### Database Design

- Added database design documentation.
- Added entity relationship documentation.
- Added normalization documentation.
- Added migration documentation.

#### Project Documentation

Expanded project documentation to include:

```text
api/
applications/
architecture/
automation/
backups/
compliance/
configs/
database/
database_design/
diagrams/
disaster_recovery/
docs/
environment/
environment_variables/
license_and_attribution/
logs/
monitoring/
networking/
notes/
performance/
security/
tests/
user_and_permissions/
virtualization/
```

### Changed

- Expanded the Go Agent from basic system information collection into a server monitoring and alerting application.
- Expanded PostgreSQL from basic inventory storage into a central server-management database.
- Added additional network interface detection for Docker and Tailscale interfaces.
- Added health status transition logic to prevent unnecessary repeated alerts.
- Added recovery detection after WARNING or CRITICAL conditions.
- Improved project documentation across the repository.
- Updated documentation to reflect the current U-Server environment.
- Expanded troubleshooting procedures for PostgreSQL, Docker, networking, UFW, and systemd.
- Improved handling of environment variables and sensitive configuration.

### Fixed

- Fixed Go Agent database code related to alert event storage.
- Fixed alert event ownership and database access issues.
- Fixed PostgreSQL connectivity after server startup.
- Fixed Web Dashboard connectivity affected by firewall configuration.
- Fixed Go Agent build errors while adding alert event support.
- Fixed health alert history so qualifying status changes can be stored in PostgreSQL.
- Corrected documentation placeholders as project components were completed.

### Verified

The following functionality has been manually tested:

- PostgreSQL service startup
- PostgreSQL cluster availability
- Go Agent compilation
- Go Agent full server refresh
- Server inventory collection
- Operating system collection
- Hardware collection
- Network interface collection
- Health check creation
- Health history storage
- Health status transition detection
- Alert event creation
- Recovery event creation
- ntfy notification delivery
- systemd Go Agent execution
- Go Agent journal logging
- Docker installation
- Docker Compose installation
- Docker container operation
- Docker networking
- Bittensor development container operation
- Django database connectivity
- UFW troubleshooting

## Project Status

Ubuntu SQL Server is currently under active development.

The project has moved beyond its original database setup and now includes server monitoring, health history, alerting, application clients, a Web Dashboard, container support, automation, networking, backups, security documentation, and disaster recovery planning.

Future releases will be assigned version numbers as larger development milestones are completed.

## Changelog Format

Future entries should use the following sections when appropriate:

```text
Added
Changed
Fixed
Removed
Security
Verified
```

New changes should be added to the `Unreleased` section first.

When a project version is released, those changes can be moved into a version section such as:

```text
## 0.1.0 - YYYY-MM-DD
```
