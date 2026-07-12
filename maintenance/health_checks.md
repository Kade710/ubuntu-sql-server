# Health Checks

## Server

Hostname:
- U-Server

Operating System:
- Ubuntu 24.04.4 LTS


## Hardware Status

### Storage

Device:
- /dev/sda

Capacity:
- 1.8 TB

Current Usage:
- 14 GB used

Available:
- 1.7 TB


### Memory

Installed RAM:
- 16 GB

Current Usage:
- ~1.3 GB

Swap:
- 4 GB available


## System Status

Uptime:
- 1 day

Load Average:
- 0.00, 0.00, 0.00


## Network Status

Primary Interface:
- enp3s0

IP Address:
- 192.168.1.100


## Database Status

Database:
- ubuntu_sql_server

PostgreSQL Version:
- 16.14

Service:
postgresql@16-main


Status:
- Active (running)


## Verification Commands

Storage:
```bash
lsblk
df -h
```

Memory:
```  bash
free -h
```

System:
``` bash
uptime
```

Network:
``` bash
ip addr show
```

Database:
``` bash
systemctl status postgresql@16-main
```