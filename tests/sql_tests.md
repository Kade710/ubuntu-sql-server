# SQL Tests

## Overview

This document contains PostgreSQL tests for the Ubuntu SQL Server project.

SQL testing helps verify that the database, schema, tables, relationships, and application data are working correctly.

## Database

Main database:

```text
ubuntu_sql_server
```

Main schema:

```text
server_management
```

---

## Connect to PostgreSQL

Connect using the PostgreSQL administrative account:

```bash
sudo -u postgres psql -d ubuntu_sql_server
```

Application database access can also be tested with:

```bash
psql -h 127.0.0.1 -U <database-user> -d ubuntu_sql_server
```

---

## Check Database

Inside `psql`, display the current connection:

```sql
\conninfo
```

The connection should show the expected database and PostgreSQL server.

---

## Check Schema

List schemas:

```sql
\dn
```

The following schema should exist:

```text
server_management
```

---

## Check Tables

List tables in the project schema:

```sql
\dt server_management.*
```

Current project tables include:

```text
server_inventory
operating_systems
hardware_components
network_interfaces
health_checks
maintenance_logs
alert_events
```

---

## Server Inventory Test

Check server inventory:

```sql
SELECT *
FROM server_management.server_inventory;
```

At least one server record should exist after the Go Agent has registered U-Server.

---

## Operating System Test

```sql
SELECT *
FROM server_management.operating_systems;
```

The results should contain operating system information associated with a server record.

---

## Hardware Test

```sql
SELECT *
FROM server_management.hardware_components
ORDER BY id;
```

Results may include components such as:

```text
CPU
RAM
Motherboard
GPU
```

---

## Network Interface Test

```sql
SELECT *
FROM server_management.network_interfaces
ORDER BY id;
```

The results should contain detected physical or virtual network interfaces.

---

## Health Check Test

View recent health checks:

```sql
SELECT *
FROM server_management.health_checks
ORDER BY created_at DESC
LIMIT 10;
```

Recent health checks should contain measurements for:

- Load average
- Memory usage
- Disk usage
- Uptime
- Overall status

---

## Health Status Test

View recent health states:

```sql
SELECT id,
       server_id,
       overall_status,
       created_at
FROM server_management.health_checks
ORDER BY created_at DESC
LIMIT 10;
```

Expected status values may include:

```text
HEALTHY
WARNING
CRITICAL
```

---

## Alert Event Test

View alert history:

```sql
SELECT id,
       server_id,
       previous_status,
       new_status,
       title,
       created_at
FROM server_management.alert_events
ORDER BY created_at DESC;
```

Alert records should appear after qualifying health status changes occur.

---

## Maintenance Log Test

```sql
SELECT *
FROM server_management.maintenance_logs
ORDER BY id DESC;
```

This can be used to verify that maintenance records are being stored correctly.

---

## Foreign Key Test

Inspect a table with:

```sql
\d server_management.alert_events
```

The `alert_events` table should contain a foreign key connecting:

```text
server_id
```

to:

```text
server_management.server_inventory(id)
```

Other related tables can be inspected in the same way.

---

## Row Count Test

Table counts can be checked with:

```sql
SELECT COUNT(*)
FROM server_management.health_checks;
```

For example:

```sql
SELECT COUNT(*)
FROM server_management.alert_events;
```

This provides a quick way to confirm whether applications are creating records.

---

## Recent Data Test

To confirm that the Go Agent is currently writing data, run:

```bash
./agent --refresh
```

Then query:

```sql
SELECT *
FROM server_management.health_checks
ORDER BY created_at DESC
LIMIT 1;
```

The newest record should match the recent agent run.

---

## Read-Only Testing

When possible, troubleshooting queries should begin with `SELECT` statements.

Avoid using commands such as:

```sql
DELETE
DROP
TRUNCATE
UPDATE
```

during basic testing unless the change is intentional.

Testing should not modify production or important data without a clear reason and a backup when appropriate.

---

## Test Record

SQL tests can be documented using:

```text
Date:
Database:
Test:
Query:
Expected Result:
Actual Result:
Result: PASSED / FAILED
Notes:
```

## Future SQL Testing

Future database testing may include:

- Automated SQL tests
- Constraint testing
- Migration testing
- Backup and restore testing
- Performance testing
- Application integration testing
- Invalid data testing
- Transaction testing

Automated tests can be added as the database layer becomes more formalized.
