# Network Services

## Overview

U-Server hosts several services that use network connections.

This document provides a reference for the main network services used by the Ubuntu SQL Server project and related services running on U-Server.

## Service Summary

| Service | Default Port | Protocol | Purpose |
| --- | ---: | --- | --- |
| SSH | 22 | TCP | Remote server administration |
| PostgreSQL | 5432 | TCP | Database connections |
| Django | 8000 | TCP | Development Web Dashboard |
| Minecraft | 25565 | TCP | Minecraft server |

Ports listed here represent the standard or currently used project ports. Actual availability depends on service and firewall configuration.

---

## SSH

Default port:

```text
22/TCP
```

SSH provides remote command-line administration of U-Server.

Service status can be checked with:

```bash
systemctl status ssh
```

Listening status can be checked with:

```bash
sudo ss -lntp | grep ':22'
```

SSH should only be available to authorized users.

---

## PostgreSQL

Default port:

```text
5432/TCP
```

PostgreSQL stores information used by the Ubuntu SQL Server project.

The PostgreSQL cluster can be checked with:

```bash
pg_lsclusters
```

Listening status can be checked with:

```bash
sudo ss -lntp | grep ':5432'
```

Applications running directly on U-Server may connect through:

```text
127.0.0.1:5432
```

PostgreSQL should not be exposed to untrusted networks.

---

## Django Web Dashboard

Development port:

```text
8000/TCP
```

The Django Web Dashboard provides browser-based access to server information.

When using Django's development server, it may be started with:

```bash
python manage.py runserver
```

To accept connections from other devices during development, Django may be started with an appropriate listening address.

Firewall access must also be configured if the dashboard needs to be reached from another system.

The Django development server is intended for development and testing rather than production deployment.

---

## Minecraft Server

Port:

```text
25565/TCP
```

The Minecraft server runs inside Docker.

The container can be checked with:

```bash
docker ps
```

The listening port can be checked with:

```bash
sudo ss -lntp | grep ':25565'
```

Container logs can be viewed with:

```bash
docker logs minecraft-server --tail 50
```

Docker publishes the Minecraft port from the container to U-Server.

---

## Tailscale

Tailscale provides private remote networking through the:

```text
tailscale0
```

interface.

Tailscale status can be checked with:

```bash
tailscale status
```

The assigned Tailscale address can be checked with:

```bash
tailscale ip
```

Services reached through Tailscale are still subject to their own listening configuration and other security controls.

---

## Docker Networking

Docker creates virtual networks for containers.

Current Docker networks can be viewed with:

```bash
docker network ls
```

Detailed information about a network can be viewed with:

```bash
docker network inspect <network-name>
```

Published container ports can be viewed with:

```bash
docker ps
```

Docker Compose may create additional bridge networks when applications are started.

---

## Firewall

UFW controls incoming network access to U-Server.

Current rules can be viewed with:

```bash
sudo ufw status numbered
```

Firewall rules should only allow services that require incoming connections.

Opening a firewall port does not automatically start a service. Likewise, a running service may still be unreachable if the required firewall rule is missing.

---

## Troubleshooting

If a network service cannot be reached:

1. Confirm the service is running.
2. Check whether the expected port is listening.
3. Verify the service is listening on the correct interface.
4. Check UFW rules.
5. Check Docker port mappings when containers are involved.
6. Verify the server IP address.
7. Test local access from U-Server.
8. Test access from the remote client.
9. Review the service logs.

Useful commands include:

```bash
ip -br addr
sudo ss -lntp
sudo ufw status numbered
docker ps
tailscale status
```

## Security

Only required network services should be exposed.

Administrative and database services should be limited to trusted systems whenever possible.

Firewall rules should be reviewed whenever a new network service is added.
