# IP Addressing

## Overview

U-Server uses several network interfaces for physical networking, remote access, and container networking.

The Go Agent detects these interfaces and stores network information in PostgreSQL.

## Primary Network Interface

The primary physical interface is:

```text
enp3s0
```

Current address:

```text
192.168.1.100
```

Connection type:

```text
Ethernet
```

Link speed:

```text
1000 Mbps
```

This interface provides access to the local network.

---

## Tailscale

Tailscale creates a virtual network interface:

```text
tailscale0
```

Current Tailscale address:

```text
<tailscale-ip>
```

Tailscale provides a separate private network path for remote access between authorized devices.

---

## Docker

Docker creates virtual network interfaces for containers.

The default Docker bridge may appear as:

```text
docker0
```

Current address:

```text
172.17.0.1
```

Docker Compose may also create additional bridge interfaces.

An example detected on U-Server is:

```text
br-86e61bc0a6a5
```

These interfaces allow Docker containers to communicate with the host, other containers, and external networks depending on their configuration.

Docker-generated bridge names and addresses may change when networks are recreated.

---

## Loopback

Linux also provides the loopback interface:

```text
lo
```

The standard IPv4 loopback address is:

```text
127.0.0.1
```

Applications using `127.0.0.1` are connecting to services on U-Server itself.

For example, PostgreSQL applications running directly on U-Server may connect to:

```text
127.0.0.1:5432
```

---

## Address Summary

| Interface | Address | Purpose |
| --- | --- | --- |
| `enp3s0` | `192.168.1.100` | Local network |
| `tailscale0` | `<tailscale-ip>` | Tailscale private network |
| `docker0` | `172.17.0.1` | Docker bridge |
| `lo` | `127.0.0.1` | Local host |

Additional Docker bridge interfaces may be created as container networks are added.

## Checking Addresses

Current network addresses can be viewed with:

```bash
ip addr
```

A shorter summary can be viewed with:

```bash
ip -br addr
```

The primary interface can be checked with:

```bash
ip addr show enp3s0
```

Tailscale can be checked with:

```bash
ip addr show tailscale0
```

## Important Note

IP addresses may change depending on DHCP, network configuration, Docker networks, and other services.

This document should be updated when permanent addressing changes are made.
