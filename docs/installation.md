# Ubuntu Installation Documentation

## Installation Overview

Server Name:
- U-Server

Operating System:
- Ubuntu 24.04.4 LTS Desktop

Installation Type:
- Clean installation

Purpose:
- Linux server development environment
- PostgreSQL database hosting
- SQL development practice
- Linux administration practice
- Future application hosting


# Hardware Environment

The operating system was installed on custom-built hardware.

Hardware Summary:

- Manufacturer: MSI
- Motherboard: MSI Z97 GAMING 5 (MS-7917)
- Processor: Intel Core i5-4690K @ 3.50GHz
- Memory: 16GB DDR3
- Storage: 2TB drive
- Graphics: NVIDIA GeForce GTX 750
- Network: Qualcomm Atheros Killer E220x Gigabit Ethernet


# Operating System Installation

Ubuntu 24.04.4 LTS Desktop was installed to provide a graphical Linux environment while developing and configuring the server.

The desktop environment allows:

- Local administration
- Software development
- VS Code usage
- Hardware troubleshooting
- Database management

The system will transition toward headless server management using SSH.


# Initial System Configuration

After installation, the system was updated to ensure all packages were current.

Commands used:

```bash
sudo apt update
sudo apt upgrade