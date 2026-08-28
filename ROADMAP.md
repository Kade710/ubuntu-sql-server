# Ubuntu SQL Server Project Roadmap

## Version 0.1 - Project Initialization

- [x] Create GitHub repository
- [x] Create project folder structure
- [x] Add README.md
- [x] Add LICENSE
- [x] Add .gitignore
- [x] Create documentation placeholders
- [x] Create SQL placeholders
- [x] Create application placeholders
- [x] Create scripts placeholders
- [x] Create security placeholders
- [x] Create networking placeholders
- [x] Create monitoring placeholders
- [x] Create backup placeholders

---

# Version 1.0 - Ubuntu Server Foundation

## Hardware Setup

- [x] Verify server hardware specifications
- [x] Document CPU model
- [x] Document RAM capacity
- [x] Document storage devices
- [x] Document GPU model
- [ ] Document power supply specifications
- [ ] Install 250GB SSD
- [ ] Install 1TB HDD
- [ ] Verify BIOS settings
- [ ] Enable virtualization support (if available)

## Ubuntu Installation

- [x] Download Ubuntu Desktop ISO
- [x] Create bootable USB installer
- [x] Install Ubuntu on SSD
- [x] Configure user account
- [x] Set hostname
- [ ] Configure timezone
- [x] Apply system updates
- [x] Document installation process

---

# Version 1.1 - Remote Administration

## SSH Configuration

- [x] Install OpenSSH server
- [x] Verify SSH service is running
- [x] Connect from Windows workstation
- [ ] Configure static IP address
- [x] Document SSH setup
- [ ] Create SSH key authentication
- [ ] Disable root SSH login
- [ ] Disable password authentication (after testing)

---

# Version 1.2 - PostgreSQL Database Server

## Database Installation

- [x] Install PostgreSQL
- [x] Verify PostgreSQL service
- [x] Configure PostgreSQL startup
- [x] Create database administrator account
- [x] Create first database
- [x] Create database users
- [x] Configure permissions

## Database Development

- [x] Create database schema
- [x] Create tables
- [x] Define primary keys
- [x] Define foreign keys
- [x] Add sample data
- [x] Create SQL queries
- [x] Document database design
- [x] Create ER diagram

---

# Version 1.3 - Backup and Recovery

## Backup System

- [x] Create database backup script
- [x] Test database backup
- [x] Create restore script
- [x] Test database restoration
- [x] Document backup process
- [x] Create backup schedule

---

# Version 1.4 - Security Hardening

## System Security

- [x] Configure UFW firewall
- [ ] Configure SSH security
- [ ] Create least-privilege users
- [x] Review file permissions
- [ ] Enable automatic security updates
- [x] Document security configuration

## Database Security

- [x] Secure PostgreSQL authentication
- [x] Create database roles
- [x] Limit database permissions
- [x] Review database exposure

---

# Version 1.5 - Monitoring

## System Monitoring

- [x] Monitor CPU usage
- [x] Monitor RAM usage
- [x] Monitor disk usage
- [x] Monitor system uptime
- [x] Document system health checks

## Future Monitoring Tools

- [ ] Install Prometheus
- [ ] Install Grafana
- [ ] Create monitoring dashboard

---

# Version 2.0 - Application Development

## Python Application

- [x] Create Python PostgreSQL client
- [x] Connect application to database
- [ ] Create CRUD operations
- [x] Document setup

### Python Client Hardening

#### Configuration

- [x] Validate required database environment variables
- [x] Validate PostgreSQL port
- [x] Preserve database password exactly
- [x] Protect environment credentials from Git

#### Database

- [x] Add PostgreSQL connection timeout
- [x] Validate queries before execution
- [x] Add contextual database errors
- [x] Ensure database connections are closed
- [x] Verify live PostgreSQL connectivity

#### Inventory

- [x] Harden server inventory query
- [x] Add deterministic inventory ordering
- [x] Verify inventory against live database

#### Hardware

- [x] Validate server ID
- [x] Reject invalid server IDs
- [x] Verify hardware query against live database

#### Network

- [x] Validate server ID
- [x] Reject invalid server IDs
- [x] Verify network query against live database

#### Maintenance

- [x] Validate server ID
- [x] Reject invalid server IDs
- [x] Verify maintenance query against live database

#### Data Models

- [x] Refactor Server model to dataclass
- [x] Add model type annotations
- [x] Verify model compilation and compatibility

#### Report Generation

- [x] Use UTF-8 report output
- [x] Support configurable output directory
- [x] Prevent incomplete final reports
- [x] Add collision-resistant report filenames
- [x] Return generated report path
- [x] Verify generated report contents

#### Administrative Client

- [x] Harden menu input handling
- [x] Handle invalid menu options
- [x] Handle Ctrl+C cleanly
- [x] Handle EOF cleanly
- [x] Handle expected runtime errors
- [x] Verify all menu options

#### Dependencies

- [x] Pin verified Python dependencies
- [x] Verify requirements installation

#### Python Client Final Verification

- [x] Complete audit of Python Client source files
- [x] Compile all active Python source files
- [x] Verify all Python modules import together
- [x] Test all interactive menu options
- [x] Test invalid menu input
- [x] Test normal exit
- [x] Test Ctrl+C handling
- [x] Test EOF handling
- [x] Verify live PostgreSQL queries
- [x] Verify report generation
- [x] Inspect generated report contents
- [x] Verify dependency installation
- [x] Review Git working tree
- [x] Verify repository synchronized with origin

**Status:** HARDENED / FINAL VERIFICATION COMPLETE

---

## Go Application

- [ ] Create Go API
- [x] Connect Go application to PostgreSQL
- [ ] Create REST endpoints
- [ ] Add authentication

### Go Agent Hardening

#### Layer: Command / Agent Entry Point

`applications/go_agent/cmd/agent/main.go`

- [x] Add command-line argument validation
- [x] Add proper exit codes for command failures
- [x] Harden full server refresh error handling
- [x] Stop refresh when a required stage fails
- [x] Verify interactive agent operation
- [x] Verify non-interactive refresh operation

**Status:** HARDENED / VERIFIED

---

#### Layer: Configuration

`applications/go_agent/internal/config/config.go`

- [x] Validate required database configuration
- [x] Validate PostgreSQL port configuration
- [x] Prevent database password output
- [x] Trim appropriate configuration values
- [x] Preserve database password exactly
- [x] Verify invalid configuration failure handling

**Status:** HARDENED / VERIFIED

---

#### Layer: Alerting

`applications/go_agent/internal/alerts/ntfy.go`

- [x] Add ntfy request timeout
- [x] Validate notification input
- [x] Safely construct notification topic URL
- [x] Limit notification error response size
- [x] Handle notification response errors
- [x] Test health warning notifications
- [x] Test recovery notifications
- [x] Verify alert event database logging

**Status:** HARDENED / END-TO-END VERIFIED

---

#### Layer: Database

`applications/go_agent/internal/database/database.go`

##### Connection Management

- [x] Add PostgreSQL connection pool limits
- [x] Add PostgreSQL connection lifetime limits
- [x] Add PostgreSQL idle connection lifetime
- [x] Add database connection timeout

##### Query and Transaction Safety

- [x] Add database query timeouts
- [x] Add database transaction timeouts
- [x] Validate health history query limits
- [x] Improve database transaction error reporting

##### Database Operations

- [x] Verify server inventory database operations
- [x] Verify network interface database operations
- [x] Verify maintenance log database operations
- [x] Verify operating system database operations
- [x] Verify hardware database operations
- [x] Verify health check database operations
- [x] Verify alert event database operations

**Status:** HARDENED / END-TO-END VERIFIED

---

#### Layer: Hardware Collection

`applications/go_agent/internal/hardware/hardware.go`

- [x] Harden hardware command execution
- [x] Add external command timeouts
- [x] Collect CPU information
- [x] Collect storage information
- [x] Collect GPU information
- [x] Collect motherboard information
- [x] Add optional PSU inventory support
- [x] Validate PSU wattage input
- [ ] Populate PSU specifications after hardware upgrade

**Status:** HARDENED / VERIFIED  
**PSU inventory:** IMPLEMENTED / HARDWARE DATA PENDING

---

#### Layer: Hardware Components

`applications/go_agent/internal/hardware/components.go`

- [x] Harden hardware component generation
- [x] Normalize component strings
- [x] Protect against invalid RAM values
- [x] Improve CPU manufacturer detection
- [x] Improve GPU manufacturer detection
- [x] Improve motherboard manufacturer detection
- [x] Add PSU component support
- [x] Prevent empty PSU records
- [x] Remove test PSU database record
- [x] Verify hardware component registration

**Status:** HARDENED / VERIFIED

---

#### Layer: System Health

`applications/go_agent/internal/health/health.go`

- [x] Harden system health metric parsing
- [x] Validate health metric values
- [x] Check memory scanner errors
- [x] Prevent invalid memory calculations
- [x] Prevent invalid disk calculations
- [x] Validate load average
- [x] Validate uptime
- [x] Add named health thresholds
- [x] Separate warning and critical load thresholds
- [x] Verify hardened health collection
- [x] Verify health check database storage

**Status:** HARDENED / VERIFIED

---

#### Layer: Server Inventory

`applications/go_agent/internal/inventory/inventory.go`

- [x] Harden server inventory collection
- [x] Normalize hostname
- [x] Harden operating system name parsing
- [x] Prefer physical LAN IPv4 over virtual interfaces
- [x] Exclude Docker interfaces from preferred IP selection
- [x] Exclude Tailscale interfaces from preferred IP selection
- [x] Retain fallback IPv4 selection
- [x] Verify server inventory database update
- [x] Verify physical LAN address selection

**Status:** HARDENED / VERIFIED

---

#### Layer: Maintenance

`applications/go_agent/internal/maintenance/`

- [x] Audit `maintenance.go`
- [x] Review maintenance log data structure
- [x] Verify maintenance model fields
- [x] Verify `maintenance.go` formatting and compilation
- [x] Audit remaining maintenance package files
- [x] Harden maintenance collection
- [x] Validate maintenance input
- [x] Verify maintenance database integration

**Status:** HARDENED / END-TO-END VERIFIED

---

#### Layer: Networking

`applications/go_agent/internal/network/`

- [x] Audit network package files
- [x] Harden network interface collection
- [x] Validate interface data
- [x] Review virtual interface handling
- [x] Verify network database integration

**Status:** HARDENED / END-TO-END VERIFIED

---

#### Layer: Operating System Information

`applications/go_agent/internal/osinfo/`

- [x] Audit operating system information files
- [x] Harden OS information collection
- [x] Validate parsed OS information
- [x] Verify OS database integration

**Status:** HARDENED / END-TO-END VERIFIED

---

#### Layer: System Information

`applications/go_agent/internal/system/`

- [x] Audit system package files
- [x] Harden hostname collection
- [x] Harden memory collection
- [x] Validate system metric parsing
- [x] Verify system information integration

**Status:** HARDENED / END-TO-END VERIFIED

---

#### Go Agent Final Verification

- [x] Complete audit of all Go Agent source files
- [x] Run `gofmt` across Go Agent
- [x] Run clean Go Agent build
- [x] Test all interactive menu options
- [x] Test `--refresh`
- [x] Test systemd execution
- [x] Verify database records after final test
- [x] Verify warning notification path
- [x] Verify recovery notification path
- [x] Review Git working tree
- [x] Push final hardened Go Agent

**Status:** IN PROGRESS

---

## Rust Application

- [x] Create Rust database client
- [x] Test PostgreSQL integration
- [x] Document implementation

## Web Dashboard

- [x] Create dashboard interface
- [x] Display database information
- [x] Display server statistics

---

# Version 3.0 - Advanced Homelab Expansion

## Storage Expansion

- [x] Add additional storage drives
- [ ] Research RAID configurations
- [ ] Configure redundant storage
- [ ] Document storage architecture

## NAS Development

- [ ] Install file sharing services
- [ ] Configure network shares
- [ ] Create user permissions
- [ ] Configure backup storage

## Virtualization

- [ ] Install virtual machine platform
- [ ] Create test virtual machines
- [ ] Document VM management
- [x] Install Docker and Docker Compose
- [x] Document container management

---

# Completed Features

- [x] Ubuntu server environment
- [x] PostgreSQL database server
- [x] Server management database
- [x] Go monitoring agent
- [x] Python PostgreSQL client
- [x] Rust PostgreSQL client
- [x] Django Web Dashboard
- [x] Server inventory collection
- [x] Hardware inventory collection
- [x] Operating system inventory
- [x] Network interface inventory
- [x] Automated health checks
- [x] Health history
- [x] Health status alerts
- [x] Recovery notifications
- [x] Alert event history
- [x] systemd Go Agent automation
- [x] Database backup and recovery
- [x] Docker and Docker Compose
- [x] Project security documentation
- [x] Disaster recovery documentation
- [x] Go Agent command-line hardening
- [x] Go Agent refresh failure handling
- [x] Configuration validation
- [x] ntfy notification hardening
- [x] Health alert end-to-end verification
- [x] Go Agent database layer hardening
- [x] Go Agent hardware layer hardening
- [x] Go Agent health layer hardening
- [x] Go Agent inventory layer hardening
- [x] Go Agent maintenance layer hardening

---

# Future Ideas

## Development & Automation

- [ ] Add CI/CD pipeline
- [ ] Add automated testing

## Python Application

- [ ] Implement full Python CRUD operations

## Infrastructure

- [x] Add Docker and Docker Compose support
- [ ] Add cloud backup

## Database & Data Maintenance

- [ ] Remove stale network interface records during server refresh

## Documentation & Recovery

- [x] Add API documentation
- [x] Add disaster recovery plan