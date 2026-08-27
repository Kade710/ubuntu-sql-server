# Firewall Security

## Overview

U-Server uses UFW to control incoming network connections.

The firewall helps prevent unnecessary access to services running on the server.

A service should only be exposed through the firewall when another trusted system needs to connect to it.

## Check Firewall Status

Current firewall rules can be viewed with:

```bash
sudo ufw status numbered
```

A more detailed status can be viewed with:

```bash
sudo ufw status verbose
```

## Basic Management

Enable UFW:

```bash
sudo ufw enable
```

Disable UFW:

```bash
sudo ufw disable
```

Disabling the firewall should only be done when necessary for troubleshooting or maintenance.

## Allowing Services

A service can be allowed by application profile when one exists.

For example:

```bash
sudo ufw allow OpenSSH
```

A specific TCP port can also be allowed:

```bash
sudo ufw allow <port>/tcp
```

Rules should only be added when the service requires incoming network access.

## Removing Rules

Numbered rules can be viewed with:

```bash
sudo ufw status numbered
```

A rule can then be removed with:

```bash
sudo ufw delete <rule-number>
```

Always review the numbered rules again after deleting one because the rule numbers may change.

## SSH

SSH normally uses:

```text
22/TCP
```

SSH access should be confirmed before changing firewall rules remotely.

Removing SSH access while connected remotely could prevent further administration of U-Server.

## PostgreSQL

PostgreSQL normally uses:

```text
5432/TCP
```

Applications running directly on U-Server can connect through localhost without requiring PostgreSQL to be exposed to the local network.

PostgreSQL should only be opened through the firewall when remote database access is actually required.

Database access should also be controlled through PostgreSQL configuration and authentication.

## Django Web Dashboard

The Django development server commonly uses:

```text
8000/TCP
```

Firewall access may be required when testing the dashboard from another device.

Opening port 8000 does not automatically make Django available. The application must also be listening on an address that accepts remote connections.

The Django development server should not be treated as a production web server.

## Minecraft

The Minecraft server uses:

```text
25565/TCP
```

Because Minecraft runs inside Docker, both Docker port publishing and host firewall configuration should be considered when troubleshooting access.

## Tailscale

Tailscale provides another private network path to U-Server.

Firewall rules and application listening addresses should still be considered when deciding which services should be reachable.

Tailscale should not be treated as a reason to expose unnecessary services publicly.

## Checking Listening Ports

Firewall rules should be compared with services that are actually listening.

View listening TCP ports with:

```bash
sudo ss -lntp
```

Check a specific port with:

```bash
sudo ss -lntp | grep <port>
```

A firewall rule does not start a service.

Likewise, a running service can still be unreachable when its firewall rule or listening configuration is incorrect.

## Troubleshooting

If a service is running but cannot be reached:

1. Confirm the service is running.
2. Check its listening port.
3. Check the listening address.
4. Review UFW rules.
5. Check Docker port mappings if containers are involved.
6. Check the server IP address.
7. Test the connection locally.
8. Test the connection from the remote system.
9. Review service logs.

## Security Guidelines

- Keep UFW enabled when possible.
- Only allow required services.
- Remove rules for services that are no longer used.
- Avoid exposing PostgreSQL unnecessarily.
- Protect administrative services.
- Review firewall rules after installing new applications.
- Document important firewall changes.
- Verify remote access before removing SSH-related rules.

## Review

Firewall rules should be reviewed whenever network services are added, removed, or changed.
