# Go Agent

The Go Agent collects system information from U-Server and stores it in the PostgreSQL database.

It is used to keep server information up to date and monitor the health of the system.

## Features

- Collects server inventory information
- Records operating system information
- Detects hardware components
- Records network interfaces
- Runs system health checks
- Tracks CPU load, memory usage, disk usage, and uptime
- Stores health history in PostgreSQL
- Detects changes in server health
- Sends warning, critical, and recovery notifications
- Records alert events in the database

## Health Monitoring

The agent classifies the server as:

- `HEALTHY`
- `WARNING`
- `CRITICAL`

When the health status changes, the agent can send a notification and record the event in the database.

## Database

The agent stores collected information in the `server_management` schema of the `ubuntu_sql_server` PostgreSQL database.

## Status

The Go Agent is working and actively being developed. More monitoring and automation features may be added as the project grows.