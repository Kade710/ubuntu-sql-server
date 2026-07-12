# Monitoring Alerts

## Overview

This document defines alert thresholds used to identify potential issues affecting U-Server.

Server:

* U-Server

Operating System:

* Ubuntu 24.04.4 LTS Desktop

Monitoring Objectives:

* Detect hardware resource issues
* Verify network connectivity
* Monitor PostgreSQL availability
* Identify storage capacity problems
* Ensure backup operations remain functional

# CPU Alerts

## High CPU Utilization

Condition:

* CPU utilization remains above 90%

Action:

* Identify high CPU processes
* Review running services
* Investigate abnormal workloads

# Memory Alerts

## Low Available Memory

Condition:

* Available memory falls below 2 GB

Action:

* Check memory usage
* Identify high memory processes
* Restart services if necessary

# Storage Alerts

## Low Disk Space

Condition:

* Disk utilization exceeds 80%

Action:

* Review disk usage
* Remove unnecessary files
* Expand storage if required

# Network Alerts

## Connectivity Failure

Condition:

* Unable to reach external network

Verification Command:

```bash
ping -c 4 8.8.8.8
```

Action:

* Verify Ethernet connection
* Check IP configuration
* Verify gateway connectivity

# PostgreSQL Alerts

## Database Service Stopped

Condition:

* PostgreSQL service is inactive

Verification Command:

```bash
systemctl status postgresql@16-main
```

Action:

* Restart PostgreSQL service
* Review service logs
* Verify database availability

## Database Connection Failure

Condition:

* Unable to connect to ubuntu_sql_server

Verification Command:

```bash
psql -U jonathon_admin -d ubuntu_sql_server
```

Action:

* Verify PostgreSQL service
* Verify user credentials
* Review PostgreSQL logs

# Backup Alerts

## Backup Failure

Condition:

* pg_dump backup does not complete successfully

Action:

* Verify disk space
* Verify database connectivity
* Re-run backup
* Review backup logs

# SSH Alerts

## Remote Access Failure

Condition:

* SSH connection cannot be established

Verification Command:

```bash
ssh jonathon@U-Server
```

Action:

* Verify SSH service
* Confirm network connectivity
* Review firewall configuration

# Maintenance Response

When an alert occurs:

1. Identify the affected service.
2. Verify the issue using system commands.
3. Apply corrective action.
4. Confirm the issue has been resolved.
5. Record the maintenance activity in the PostgreSQL maintenance_logs table.
