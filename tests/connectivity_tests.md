# Connectivity Tests

## Overview

This document contains basic connectivity tests for services used by the Ubuntu SQL Server project.

These tests can help determine whether a problem is caused by the application, network, firewall, database, or service configuration.

---

## Network Interface Test

View current network interfaces:

```bash
ip -br addr
```

The primary physical interface should appear in the output.

Additional interfaces may include Tailscale and Docker networks.

---

## Local Network Test

Test the local network gateway or another known device:

```bash
ping -c 4 <ip-address>
```

A successful response confirms basic IP connectivity.

---

## Internet Connectivity Test

Test external network connectivity:

```bash
ping -c 4 1.1.1.1
```

If this works but domain names do not, the problem may involve DNS rather than the network connection itself.

---

## DNS Test

Test name resolution with:

```bash
getent hosts github.com
```

A successful result should return an IP address.

---

## SSH Test

Check the SSH service:

```bash
systemctl status ssh
```

Check whether SSH is listening:

```bash
sudo ss -lntp | grep ':22'
```

From another authorized system, test SSH with:

```bash
ssh <user>@<server-address>
```

A successful login confirms both network connectivity and SSH access.

---

## PostgreSQL Service Test

Check PostgreSQL:

```bash
sudo systemctl status postgresql
```

Check the PostgreSQL cluster:

```bash
pg_lsclusters
```

The main PostgreSQL cluster should report:

```text
online
```

---

## PostgreSQL Port Test

Check whether PostgreSQL is listening:

```bash
sudo ss -lntp | grep ':5432'
```

If PostgreSQL is configured for local connections, the service should be listening on the configured local address.

---

## PostgreSQL Connection Test

Test access to the project database:

```bash
psql -h 127.0.0.1 -p 5432 -U <database-user> -d ubuntu_sql_server
```

A successful connection confirms that PostgreSQL is reachable and the supplied database account can authenticate.

Do not place the database password directly in this document.

---

## Go Agent Connectivity Test

From the Go Agent directory, load the environment configuration:

```bash
set -a
source .env
set +a
```

Run a server refresh:

```bash
./agent --refresh
```

A successful refresh should:

- Connect to PostgreSQL
- Update server inventory
- Update operating system information
- Update hardware information
- Update network interfaces
- Record a health check

The output should end with a successful server refresh message.

---

## Go Agent Service Test

Check the service:

```bash
systemctl status ubuntu-sql-agent.service
```

Review recent runs:

```bash
journalctl -u ubuntu-sql-agent.service -n 50 --no-pager
```

A completed agent run may show:

```text
Deactivated successfully.
Finished ubuntu-sql-agent.service - Ubuntu SQL Server Go Agent.
```

This is expected after the agent completes its task.

---

## Django Connectivity Test

When the Django development server is running, verify that it is listening on the expected port:

```bash
sudo ss -lntp | grep ':8000'
```

If the dashboard works locally but cannot be reached from another device, check:

- Django listening address
- UFW rules
- Server IP address
- Client connectivity

---

## Docker Test

Check running containers:

```bash
docker ps
```

Check Docker networks:

```bash
docker network ls
```

A container should report a healthy or running state when the application is operating normally.

---

## Minecraft Connectivity Test

Check the Minecraft container:

```bash
docker ps
```

Check port 25565:

```bash
sudo ss -lntp | grep ':25565'
```

Review recent server logs:

```bash
docker logs minecraft-server --tail 50
```

If the container is running but players cannot connect, check the firewall and Docker port mapping.

---

## Tailscale Test

Check Tailscale:

```bash
tailscale status
```

View the assigned Tailscale address:

```bash
tailscale ip
```

Connectivity to another Tailscale device can be tested with:

```bash
tailscale ping <device-name-or-address>
```

---

## Firewall Test

View current UFW rules:

```bash
sudo ufw status numbered
```

When a service is running but cannot be reached, compare its listening port with the firewall configuration.

A service being online does not automatically mean its port is allowed through the firewall.

---

## Test Record

Connectivity tests can be recorded using:

```text
Date:
Test:
Source:
Destination:
Port:
Result: PASSED / FAILED
Notes:
```

## Troubleshooting Order

When connectivity fails:

1. Confirm the service is running.
2. Confirm the expected port is listening.
3. Test local connectivity.
4. Verify the server address.
5. Check UFW.
6. Check Docker networking if applicable.
7. Check Tailscale if applicable.
8. Test from the remote system.
9. Review application logs.
