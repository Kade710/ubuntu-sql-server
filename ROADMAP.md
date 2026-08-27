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

## Go Application

- [ ] Create Go API
- [x] Connect Go application to PostgreSQL
- [ ] Create REST endpoints
- [ ] Add authentication

### Go Agent Hardening


- [x] Add command-line argument validation
- [x] Add proper exit codes for command failures
- [x] Harden full server refresh error handling
- [x] Stop refresh when a required stage fails
- [x] Validate database configuration
- [x] Validate PostgreSQL port configuration
- [x] Prevent database password output
- [x] Add ntfy request timeout
- [x] Validate notification input
- [x] Limit notification error response size
- [x] Test health warning notifications
- [x] Test recovery notifications
- [x] Verify alert event database logging

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

# Future Ideas

- [x] Add Docker and Docker Compose support
- [ ] Add CI/CD pipeline
- [ ] Add automated testing
- [x] Add API documentation
- [ ] Add cloud backup
- [x] Add disaster recovery plan