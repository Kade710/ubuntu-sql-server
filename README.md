# Ubuntu SQL Server

Ubuntu SQL Server is a personal server management project built to monitor and manage my Ubuntu server from one place.

The project collects server information and stores it in PostgreSQL. A Django web dashboard makes it easy to view server health, hardware, network information, maintenance history, and other system data.

A Go agent runs health checks and monitors things like CPU load, memory usage, disk usage, and uptime. It can also send notifications when the server enters a warning or critical state and when it recovers.

## Features

- Server inventory tracking
- Hardware and operating system information
- Network interface monitoring
- Health checks and health history
- Maintenance logs
- Warning and critical health alerts
- Recovery notifications
- PostgreSQL database storage
- Django web dashboard
- Go monitoring agent


## Built With

- Ubuntu Server
- PostgreSQL
- Go
- Python
- Django
- Rust
- Docker
- HTML/CSS
- Git

## Project Status

This project is actively being developed as part of my homelab and server management environment. New monitoring, automation, and management features will be added as the project grows.
