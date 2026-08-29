# Prometheus

Prometheus provides metrics collection and time-series storage for the Ubuntu SQL Server monitoring stack.

## Overview

Prometheus runs as a Docker container and collects metrics from the U-Server host through Node Exporter.

The monitoring flow is:

U-Server → Node Exporter → Prometheus → Grafana

Prometheus also monitors its own metrics endpoint.

## Access

Prometheus is available on the local network at:

`http://192.168.1.100:9090`

The Prometheus readiness endpoint is:

`http://192.168.1.100:9090/-/ready`

## Configuration

The Prometheus configuration file is stored at:

`monitoring/prometheus/prometheus.yml`

The configuration is mounted read-only into the Prometheus container at:

`/etc/prometheus/prometheus.yml`

The Docker deployment is defined in:

`docker/compose/monitoring.yaml`

## Scrape Configuration

Prometheus uses a 15-second scrape interval and evaluation interval.

The monitoring stack currently contains two Prometheus scrape targets:

### Prometheus

Prometheus monitors its own metrics endpoint at:

`localhost:9090`

Job name:

`prometheus`

### Node Exporter

Node Exporter provides operating system and hardware metrics from U-Server.

Prometheus connects to Node Exporter through the internal Docker network at:

`node-exporter:9100`

Job name:

`node-exporter`

Both targets have been verified as healthy with a Prometheus target status of `UP`.

## Node Exporter

Node Exporter runs as part of the Docker monitoring stack and exposes host metrics on port `9100`.

The host root filesystem is mounted read-only into the Node Exporter container so that it can collect metrics from U-Server.

Node Exporter provides metrics for:

- CPU utilization
- Memory utilization
- Filesystem capacity
- Available disk space
- Disk read and write activity
- Network receive and transmit activity
- System load
- System uptime
- Hardware temperature sensors

## Metrics

Metrics currently used by the U-Server Grafana dashboard include:

- `node_cpu_seconds_total`
- `node_memory_MemAvailable_bytes`
- `node_memory_MemTotal_bytes`
- `node_filesystem_avail_bytes`
- `node_filesystem_size_bytes`
- `node_boot_time_seconds`
- `node_load1`
- `node_network_receive_bytes_total`
- `node_network_transmit_bytes_total`
- `node_disk_read_bytes_total`
- `node_disk_written_bytes_total`
- `node_hwmon_temp_celsius`

## Validate Configuration

The active Prometheus configuration can be validated with:

```bash
docker exec prometheus promtool check config /etc/prometheus/prometheus.yml
