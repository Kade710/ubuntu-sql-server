# Container Management

## Overview

U-Server uses Docker and Docker Compose for containerized applications.

Docker provides isolated application environments while sharing the Linux host kernel.

Docker stores its data under:

```text
/var/lib/docker
```

## Docker Installation

Check the installed Docker version:

```bash
docker --version
```

Check Docker Compose:

```bash
docker compose version
```

View detailed Docker information:

```bash
docker info
```

## Container Management

List running containers:

```bash
docker ps
```

List all containers, including stopped containers:

```bash
docker ps -a
```

Start a container:

```bash
docker start <container-name>
```

Stop a container:

```bash
docker stop <container-name>
```

Restart a container:

```bash
docker restart <container-name>
```

Inspect a container:

```bash
docker inspect <container-name>
```

## Container Logs

View container logs:

```bash
docker logs <container-name>
```

View the most recent 50 lines:

```bash
docker logs --tail 50 <container-name>
```

Follow logs in real time:

```bash
docker logs -f <container-name>
```

Press `Ctrl+C` to stop following logs.

## Docker Compose

Docker Compose uses YAML configuration files to define and manage containerized applications.

Common filenames include:

```text
compose.yaml
compose.yml
docker-compose.yaml
docker-compose.yml
```

A Compose project should normally be managed from the directory containing its configuration file.

Validate the configuration:

```bash
docker compose config
```

Start the project:

```bash
docker compose up -d
```

View project containers:

```bash
docker compose ps
```

Stop the project:

```bash
docker compose stop
```

Restart the project:

```bash
docker compose restart
```

Stop and remove project containers and networks:

```bash
docker compose down
```

Persistent storage should be understood before removing volumes or application data.

## Bittensor Development Container

The Bittensor Docker project is located at:

```text
~/Projects/bittensor-node/docker/
```

Its Compose configuration is:

```text
~/Projects/bittensor-node/docker/docker-compose.yaml
```

The development container is:

```text
bittensor-node-dev
```

The project uses:

```text
docker-bittensor-node:latest
```

The NVIDIA CUDA image currently stored on U-Server is:

```text
nvidia/cuda:12.0.0-base-ubuntu22.04
```

Check the Bittensor container:

```bash
docker ps
```

Inspect it:

```bash
docker inspect bittensor-node-dev
```

View recent logs:

```bash
docker logs --tail 50 bittensor-node-dev
```

## Docker Images

List locally stored images:

```bash
docker images
```

An image being stored locally does not mean that a container using it is currently running.

Unused images should only be removed after confirming they are no longer required.

## Docker Networks

List Docker networks:

```bash
docker network ls
```

Inspect a network:

```bash
docker network inspect <network-name>
```

Common Docker network types include:

```text
bridge
host
none
```

Docker Compose may create additional project-specific bridge networks.

## Docker Network Interfaces

Docker networking creates virtual Linux interfaces on U-Server.

Examples include:

```text
docker0
br-*
```

These interfaces may also be detected by the Ubuntu SQL Server monitoring system when network inventory is collected.

Docker-generated bridge names may change when networks are removed and recreated.

## Docker Volumes

List named volumes:

```bash
docker volume ls
```

An empty named-volume list does not necessarily mean that container data is temporary.

Applications may also use bind mounts that map host directories or files directly into containers.

## Persistent Data

Important application data should not depend only on a container's writable filesystem.

Persistent data can use:

- Host directories
- Bind mounts
- Docker volumes

Important persistent data should also be included in the server backup strategy.

## Docker Resource Monitoring

Monitor running containers:

```bash
docker stats
```

This reports information such as:

- CPU usage
- Memory usage
- Network activity
- Block I/O
- Process count

Press `Ctrl+C` to exit.

## Docker Disk Usage

Check Docker storage usage:

```bash
docker system df
```

View detailed usage:

```bash
docker system df -v
```

This can identify space used by:

- Images
- Containers
- Build cache
- Volumes

Unused Docker data should only be removed after confirming it is no longer required.

## Docker Service

Check Docker:

```bash
systemctl status docker
```

Start Docker:

```bash
sudo systemctl start docker
```

Restart Docker:

```bash
sudo systemctl restart docker
```

View recent Docker service logs:

```bash
journalctl -u docker -n 50 --no-pager
```

## Docker Permissions

Docker access should be treated as privileged system access.

Check current group membership:

```bash
groups
```

Check Docker group membership:

```bash
getent group docker
```

Only trusted users should receive Docker access.

## Firewall and Ports

Check listening host ports:

```bash
sudo ss -lntp
```

Check UFW:

```bash
sudo ufw status numbered
```

Containerized services should expose only the ports they require.

Publishing a Docker port and allowing that port through the host firewall are separate configuration steps.

## Security

Docker security should follow the same security practices as the rest of U-Server.

- Only trusted users should control Docker.
- Only required ports should be published.
- Environment files containing secrets should be protected.
- Passwords and tokens should not be committed to Git.
- Secrets should not be stored directly in public Compose files.
- Container images should be kept updated.
- Containers should not receive unnecessary privileges.
- Important persistent data should be backed up.
- Firewall rules should be reviewed when services are added.

## Troubleshooting

Check Docker:

```bash
systemctl status docker
```

Check all containers:

```bash
docker ps -a
```

Review container logs:

```bash
docker logs --tail 50 <container-name>
```

Check Docker networks:

```bash
docker network ls
```

Check listening ports:

```bash
sudo ss -lntp
```

Check the firewall:

```bash
sudo ufw status numbered
```

For a Compose project, move into its project directory and run:

```bash
docker compose config
docker compose ps
```

These checks help distinguish Docker problems from application, networking, and firewall problems.
