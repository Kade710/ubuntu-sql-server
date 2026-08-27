# Monitoring

## Overview

The Ubuntu SQL Server project includes a monitoring system for tracking the health and status of U-Server.

The Go Agent collects system information, evaluates server health, stores health history in PostgreSQL, and sends notifications when important health changes occur.

The monitoring system is designed to provide both current server information and a historical record of server health.

## Monitoring Components

The current monitoring system uses:

- Go Agent
- PostgreSQL
- systemd
- Django Web Dashboard
- ntfy notifications
- Linux system information

The Go Agent performs the main collection and health evaluation work.

---

## System Health Monitoring

The Go Agent monitors several important system measurements.

Current health measurements include:

- 1-minute load average
- Memory usage
- Disk usage
- System uptime

A health check may look similar to:

```text
System Health
---
1-Minute Load Average: 0.01
Memory Usage: 12.13%
Disk Usage: 8.13%
Uptime: 122.42 hours
Overall Status: HEALTHY
```

Values change depending on the current condition of the server.

---

## Health States

The monitoring system classifies U-Server into three health states:

```text
HEALTHY
WARNING
CRITICAL
```

### HEALTHY

The server is operating within the configured health limits.

### WARNING

One or more monitored values have reached a warning threshold.

A warning indicates that the server should be watched, but it does not automatically mean that a service has failed.

### CRITICAL

One or more monitored values have reached a critical threshold.

A critical condition indicates that the server may require immediate attention.

---

## Health Checks

Each completed health check is stored in PostgreSQL.

Health records provide a history of server performance and make it possible to review previous server conditions.

Health checks include:

- Server ID
- Load average
- Memory usage
- Disk usage
- Uptime
- Overall status
- Creation time

Each health check receives its own health check ID.

---

## Health History

Health history allows the project to track server conditions over time instead of only displaying the current state.

Health records are stored in the PostgreSQL database under the `server_management` schema.

This information can be viewed through project applications such as the Django Web Dashboard and database clients.

---

## Status Changes

The Go Agent compares the current health status with the previous recorded health status.

This allows the monitoring system to detect changes such as:

```text
HEALTHY -> WARNING
HEALTHY -> CRITICAL
WARNING -> CRITICAL
WARNING -> HEALTHY
CRITICAL -> HEALTHY
```

A status change can trigger an alert event and notification.

If the health status does not change, another alert does not need to be generated for the same condition.

---

## Alert Events

Important health status changes are stored in:

```text
server_management.alert_events
```

An alert event can contain:

- Server ID
- Previous status
- New status
- Alert title
- Alert message
- Creation time

This provides a permanent history of important server health changes.

For example:

```text
Previous Status: WARNING
New Status: HEALTHY
Title: U-Server Recovered
```

---

## Recovery Monitoring

The monitoring system can detect when U-Server returns to a healthy state after a warning or critical condition.

For example:

```text
WARNING -> HEALTHY
```

or:

```text
CRITICAL -> HEALTHY
```

When a recovery is detected, the Go Agent can create a recovery alert event and send a recovery notification.

This makes it possible to know both when a problem begins and when the server returns to normal.

---

## Notifications

The Go Agent can send health notifications using ntfy.

Notifications may be sent when:

- U-Server enters a WARNING state
- U-Server enters a CRITICAL state
- U-Server recovers to HEALTHY

Notification messages may include:

- Health status
- Load average
- Memory usage
- Disk usage

Private ntfy topic names should be stored in environment configuration and should not be committed to Git.

---

## Go Agent

The Go Agent performs the main monitoring tasks.

During a server refresh, the agent currently performs:

```text
[1/5] Updating server inventory
[2/5] Updating operating system
[3/5] Updating hardware components
[4/5] Updating network interfaces
[5/5] Recording system health
```

The final step records the current health information in PostgreSQL.

---

## Running a Manual Refresh

A manual server refresh can be performed from the Go Agent directory.

Load the environment variables:

```bash
set -a
source .env
set +a
```

Run the refresh:

```bash
./agent --refresh
```

A successful refresh should update server information and create a new health check.

---

## Monitoring Through systemd

The Go Agent can also run through systemd.

Check the service with:

```bash
systemctl status ubuntu-sql-agent.service
```

View recent monitoring activity with:

```bash
journalctl -u ubuntu-sql-agent.service -n 50 --no-pager
```

Follow activity in real time with:

```bash
journalctl -u ubuntu-sql-agent.service -f
```

Press `Ctrl+C` to stop following the journal.

---

## Go Agent Service Behavior

The Go Agent performs its monitoring task and then exits.

Because of this, systemd may report:

```text
ubuntu-sql-agent.service: Deactivated successfully.
Finished ubuntu-sql-agent.service - Ubuntu SQL Server Go Agent.
```

This is normal behavior.

The service does not need to remain continuously active after a successful monitoring run.

---

## Server Inventory Monitoring

The Go Agent also records information about U-Server itself.

Inventory information includes:

- Hostname
- IP address
- Operating system
- CPU
- Memory
- Storage
- GPU
- Motherboard

This information helps identify the system being monitored and provides hardware context for health information.

---

## Operating System Monitoring

The agent records operating system information such as:

- Distribution
- Version
- Kernel
- Architecture

This helps track the software environment running on U-Server.

---

## Hardware Monitoring

Hardware information collected by the agent includes components such as:

- CPU
- RAM
- Motherboard
- GPU

Hardware information is stored separately from health history so the project can maintain organized server records.

---

## Network Monitoring

The Go Agent detects network interfaces available on U-Server.

Collected information may include:

- Interface name
- MAC address
- IP address
- Network type
- Link speed

Interfaces may include physical and virtual networking devices such as:

```text
enp3s0
tailscale0
docker0
```

Docker bridge interfaces may also appear when containers are running.

---

## Database Monitoring Data

Monitoring information is stored in the:

```text
ubuntu_sql_server
```

database.

The main schema is:

```text
server_management
```

Monitoring-related data includes health checks and alert events along with server, hardware, operating system, and network information.

---

## Web Dashboard

The Django Web Dashboard provides a browser-based way to view information collected by the monitoring system.

The dashboard can display server information without requiring direct access to PostgreSQL or the Go Agent command line.

Monitoring features can continue to be added to the dashboard as the project grows.

---

## Monitoring Logs

Go Agent monitoring activity is recorded through the systemd journal.

Recent activity can be viewed with:

```bash
journalctl -u ubuntu-sql-agent.service -n 50 --no-pager
```

Monitoring logs can help identify:

- Collection failures
- Database connection problems
- Health check failures
- Alert failures
- Notification failures
- Successful server refreshes

Additional logging information is documented in:

```text
logs/README.md
```

---

## Database Availability

The monitoring system depends on PostgreSQL for storing collected information.

PostgreSQL status can be checked with:

```bash
sudo systemctl status postgresql
```

The PostgreSQL cluster can be checked with:

```bash
pg_lsclusters
```

If PostgreSQL is unavailable, the Go Agent may be unable to save health checks or other server information.

---

## Security

Monitoring configuration may contain sensitive information.

Do not commit:

- Database passwords
- API keys
- Access tokens
- Private notification topics
- Secret environment variables

Sensitive configuration should be stored in protected environment files.

Example configuration files should only contain safe placeholder values.

---

## Troubleshooting

If monitoring stops working:

1. Check PostgreSQL.
2. Check the Go Agent service.
3. Review the Go Agent journal.
4. Verify environment variables.
5. Check database connectivity.
6. Check network connectivity.
7. Check firewall rules if network access is involved.
8. Run a manual refresh if needed.
9. Confirm that a new health check was created.

Useful commands include:

```bash
sudo systemctl status postgresql
pg_lsclusters
systemctl status ubuntu-sql-agent.service
journalctl -u ubuntu-sql-agent.service -n 50 --no-pager
```

---

## Future Improvements

Possible monitoring improvements include:

- Additional health measurements
- Service availability monitoring
- Docker container monitoring
- Network monitoring improvements
- Alert history in the Web Dashboard
- Additional notification types
- Monitoring multiple servers
- Historical charts and graphs
- Configurable health thresholds
- Additional automated recovery features

Monitoring features will continue to be added as the Ubuntu SQL Server project grows.
