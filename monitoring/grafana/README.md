# Grafana

Grafana provides visualization and dashboarding for the Ubuntu SQL Server monitoring stack.

## Overview

Grafana runs as a Docker container and uses Prometheus as its primary data source. Prometheus collects host metrics from Node Exporter running as part of the monitoring stack.

## Architecture

U-Server → Node Exporter → Prometheus → Grafana

## Access

Grafana is available on the local network at:

http://192.168.1.100:3000

## Prometheus Data Source

Grafana connects to Prometheus through the internal Docker network using:

http://prometheus:9090

The Prometheus data source has been tested successfully from Grafana.

## U-Server Monitoring Dashboard

The U-Server Monitoring dashboard provides visibility into server health and resource utilization.

Current dashboard panels include:

- CPU usage
- RAM usage
- Disk usage
- Available disk space
- System uptime
- 1-minute system load
- Network inbound and outbound traffic
- Disk read and write throughput
- CPU temperature

## Metrics

Host metrics are provided by Node Exporter and collected by Prometheus every 15 seconds.

The dashboard uses Prometheus metrics including:

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

## Docker Deployment

Grafana, Prometheus, and Node Exporter are managed through:

`docker/compose/monitoring.yaml`

Grafana persistent data is stored in the Docker-managed `grafana-data` volume.

## Dashboard Files

Dashboard definitions belong in:

`monitoring/grafana/dashboards/`

This directory is reserved for exported or provisioned Grafana dashboard JSON files.

## Security

Grafana administrative credentials must not be committed to the repository.

Secrets and environment-specific credentials should be stored outside version control or supplied through protected environment configuration.
