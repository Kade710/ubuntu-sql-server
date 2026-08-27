# Virtualization and Containers

## Overview

This directory contains documentation for virtualization and container technologies used on U-Server.

U-Server currently uses Docker and Docker Compose to run containerized services.

Containers allow applications to run in isolated environments while still sharing the Linux host system.

## Docker Installation

Docker is installed and running on U-Server.

The current Docker root directory is:

```text
/var/lib/docker
```

Docker can be checked with:

```bash
docker --version
```

Running containers can be viewed with:

```bash
docker ps
```

All containers, including stopped containers, can be viewed with:

```bash
docker ps -a
```

Detailed Docker information can be viewed with:

```bash
docker info
```

## Docker Compose

Docker Compose is installed on U-Server.

The installed version can be checked with:

```bash
docker compose version
```

Docker Compose uses YAML configuration files to define and manage containerized applications.

Common Compose filenames include:

```text
compose.yaml
compose.yml
docker-compose.yaml
docker-compose.yml
```

A Compose project should normally be managed from the directory containing its Compose file.

For example:

```bash
cd <compose-project-directory>
docker compose up -d
```

Running `docker compose up -d` from a directory without a Compose configuration will return an error stating that no configuration file was found.

## Current Containers

The currently active Docker container is:

```text
bittensor-node-dev
```

It uses the image:

```text
docker-bittensor-node:latest
```

Current container status can be checked with:

```bash
docker ps
```

Detailed information about the container can be viewed with:

```bash
docker inspect bittensor-node-dev
```

## Current Docker Images

Docker images currently stored on U-Server include:

```text
docker-bittensor-node:latest
itzg/minecraft-server:latest
nvidia/cuda:12.0.0-base-ubuntu22.04
```

Images can be viewed with:

```bash
docker images
```

An image being stored on U-Server does not mean a container using that image is currently running.

Unused images may remain available so containers can be recreated later without downloading or rebuilding the image again.

## Compose Projects

Two Docker Compose configurations are currently stored under the Projects directory.

```text
/home/jonathon/Projects/ubuntu-sql-server/minecraft/compose.yaml
/home/jonathon/Projects/bittensor-node/docker/docker-compose.yaml
```

These configurations manage separate Docker projects.

## Bittensor Development Container

The Bittensor Docker project is located at:

```text
~/Projects/bittensor-node/docker/
```

Its Compose configuration is:

```text
~/Projects/bittensor-node/docker/docker-compose.yaml
```

The current development container is:

```text
bittensor-node-dev
```

The container currently uses:

```text
docker-bittensor-node:latest
```

as its Docker image.

The Bittensor Docker environment also uses an NVIDIA CUDA base image during its container build process.

The CUDA image currently stored on U-Server is:

```text
nvidia/cuda:12.0.0-base-ubuntu22.04
```

The Bittensor container can be checked with:

```bash
docker ps
```

Its logs can be viewed with:

```bash
docker logs bittensor-node-dev
```

Recent logs can be viewed with:

```bash
docker logs --tail 50 bittensor-node-dev
```

Live logs can be followed with:

```bash
docker logs -f bittensor-node-dev
```

Press `Ctrl+C` to stop following the logs.

## Minecraft Docker Project

The Minecraft Docker project is stored at:

```text
~/Projects/ubuntu-sql-server/minecraft/
```

Its Compose configuration is:

```text
~/Projects/ubuntu-sql-server/minecraft/compose.yaml
```

The Minecraft Docker image currently stored on U-Server is:

```text
itzg/minecraft-server:latest
```

The Minecraft image being present does not mean the Minecraft server is currently running.

The current Docker container list should always be checked before assuming the Minecraft server is online:

```bash
docker ps
```

The Minecraft Compose configuration can be checked with:

```bash
cd ~/Projects/ubuntu-sql-server/minecraft
docker compose config
```

The project can be started with:

```bash
cd ~/Projects/ubuntu-sql-server/minecraft
docker compose up -d
```

After starting it, verify the container with:

```bash
docker ps
```

## Docker Networks

Docker creates virtual networks for container communication.

Current Docker networks include:

```text
bridge
docker_default
host
none
```

Networks can be viewed with:

```bash
docker network ls
```

Detailed information about a network can be viewed with:

```bash
docker network inspect <network-name>
```

### bridge

The default Docker bridge network provides networking for containers that use Docker's standard bridge configuration.

### docker_default

The `docker_default` network is a Docker bridge network created for a Compose project.

Docker Compose normally creates networks for applications when they are started.

### host

The `host` network allows a container to use the host's network stack directly when configured to do so.

### none

The `none` network provides a container with no normal external network connectivity.

## Docker Network Interfaces

Docker networking may create virtual Linux interfaces on U-Server.

Examples include:

```text
docker0
```

and interfaces beginning with:

```text
br-
```

These bridge interfaces may also be detected by the Ubuntu SQL Server Go Agent when network inventory is collected.

Docker-generated bridge names can change when networks are removed and recreated.

## Docker Volumes

There are currently no named Docker volumes listed on U-Server.

Named volumes can be checked with:

```bash
docker volume ls
```

An empty volume list does not mean container data is temporary.

Applications can also use bind mounts that map directories or files from U-Server directly into a container.

## Persistent Data

Important application data should not depend only on the writable filesystem inside a container.

Containers may be removed and recreated during updates, troubleshooting, or configuration changes.

Persistent data should use one of the following methods:

- Host directories
- Bind mounts
- Docker volumes

The correct storage method depends on the application.

Important persistent data should also be included in an appropriate backup plan.

## Container Logs

Docker captures output generated by containers.

Logs for a container can be viewed with:

```bash
docker logs <container-name>
```

Recent logs can be viewed with:

```bash
docker logs --tail 50 <container-name>
```

Live logs can be followed with:

```bash
docker logs -f <container-name>
```

Press `Ctrl+C` to stop following live logs.

Container logs are useful for troubleshooting:

- Application startup problems
- Container crashes
- Configuration problems
- Dependency failures
- Network problems
- Application errors
- Unexpected shutdowns

## Container Management

### List Running Containers

```bash
docker ps
```

### List All Containers

```bash
docker ps -a
```

### Start a Container

```bash
docker start <container-name>
```

### Stop a Container

```bash
docker stop <container-name>
```

### Restart a Container

```bash
docker restart <container-name>
```

### Inspect a Container

```bash
docker inspect <container-name>
```

### View Container Logs

```bash
docker logs <container-name>
```

## Docker Compose Management

A Compose project should normally be managed from its project directory.

### Validate Configuration

```bash
docker compose config
```

### Start the Project

```bash
docker compose up -d
```

### View Project Containers

```bash
docker compose ps
```

### Stop the Project

```bash
docker compose stop
```

### Restart the Project

```bash
docker compose restart
```

### Stop and Remove Project Containers

```bash
docker compose down
```

`docker compose down` removes containers and project networks created by Compose.

Persistent data should be understood before removing volumes or other storage.

## Port Publishing

Docker containers may publish application ports through U-Server.

Published ports can be viewed with:

```bash
docker ps
```

A published port may appear similar to:

```text
0.0.0.0:25565->25565/tcp
```

This means traffic arriving on port `25565` on U-Server is forwarded to port `25565` inside the container.

A published Docker port and a firewall rule are separate parts of network access.

When troubleshooting a containerized service, check:

1. Whether the container is running.
2. Whether the application is running inside the container.
3. Whether the required port is published.
4. Whether UFW allows the connection.
5. Whether the client can reach U-Server.

## Firewall

UFW is used on U-Server to control incoming network connections.

Current firewall rules can be viewed with:

```bash
sudo ufw status numbered
```

Containerized services should only expose ports that are actually required.

Firewall configuration should be reviewed whenever a new containerized network service is added.

## Docker Resource Monitoring

Docker can report resource usage for running containers.

Use:

```bash
docker stats
```

This can display information such as:

- CPU usage
- Memory usage
- Network activity
- Block I/O
- Process count

Press `Ctrl+C` to exit.

A single container should not be allowed to consume unnecessary host resources.

Resource limits may be added when monitoring shows that they are needed.

## Docker Disk Usage

Docker storage usage can be checked with:

```bash
docker system df
```

Detailed usage can be viewed with:

```bash
docker system df -v
```

This can help identify disk space being used by:

- Images
- Containers
- Build cache
- Volumes

Unused Docker data should only be removed after confirming it is no longer required.

## Docker Permissions

Docker access should be treated as privileged system access.

Users with permission to control the Docker daemon can perform operations that have significant control over U-Server.

Current group membership can be checked with:

```bash
groups
```

Docker group membership can be checked with:

```bash
getent group docker
```

Only trusted users should receive Docker access.

## Docker Service

The Docker service can be checked with:

```bash
systemctl status docker
```

If Docker is not running, it can be started with:

```bash
sudo systemctl start docker
```

Docker can be restarted with:

```bash
sudo systemctl restart docker
```

Docker service logs can be viewed with:

```bash
journalctl -u docker
```

Recent Docker service logs can be viewed with:

```bash
journalctl -u docker -n 50 --no-pager
```

## Troubleshooting

If a containerized application stops working, begin by checking Docker itself:

```bash
systemctl status docker
```

Then check running containers:

```bash
docker ps
```

Check all containers:

```bash
docker ps -a
```

Review the affected container:

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

If Docker Compose is involved, move into the correct project directory and run:

```bash
docker compose config
docker compose ps
```

This helps separate Docker problems from application, networking, or firewall problems.

## Recovery

Containerized applications can often be recreated using their Compose configuration and persistent data.

A basic recovery process is:

1. Confirm Docker is running.
2. Locate the application's Compose file.
3. Confirm persistent application data is available.
4. Restore data from backup if necessary.
5. Validate the Compose configuration.
6. Start the project.
7. Check the container status.
8. Review application logs.
9. Verify network access.
10. Confirm the application is working normally.

Example:

```bash
docker compose config
docker compose up -d
docker compose ps
```

## Security

Docker security should follow the same basic security practices as the rest of U-Server.

- Only trusted users should control Docker.
- Only required ports should be published.
- Environment files should be protected.
- Passwords and tokens should not be committed to Git.
- Secrets should not be placed directly in public Compose files.
- Container images should be kept updated.
- Container logs should be reviewed after unexpected behavior.
- Important persistent data should be backed up.
- Firewall rules should be reviewed when services are added.
- Containers should not receive unnecessary privileges.

## Virtual Machines

U-Server currently uses containers for the documented virtualized application environments.

The current environment does not depend on a full virtual machine platform for these services.

If a hypervisor or virtual machine platform is added later, such as:

- KVM
- QEMU
- VirtualBox
- VMware

its configuration and management can be documented in this directory.

## Current Environment Summary

The current container environment includes:

```text
Container Platform: Docker
Container Management: Docker Compose
Docker Root: /var/lib/docker

Running Container:
- bittensor-node-dev

Stored Images:
- docker-bittensor-node:latest
- itzg/minecraft-server:latest
- nvidia/cuda:12.0.0-base-ubuntu22.04

Docker Networks:
- bridge
- docker_default
- host
- none

Named Docker Volumes:
- None currently listed

Compose Projects:
- ubuntu-sql-server/minecraft
- bittensor-node/docker
```

This section represents the current documented state and should be updated when containers, images, networks, volumes, or Compose projects change.

## Future Improvements

Possible future improvements include:

- Additional container monitoring
- Docker health monitoring
- Resource usage history
- Container backup automation
- Improved container recovery procedures
- Additional Compose projects
- Image update management
- Automated image cleanup
- Improved Docker network documentation
- Dashboard container status
- Container health alerts
- Virtual machine support if needed

This document should be updated whenever major Docker services or virtualization technologies are added, removed, or changed.
