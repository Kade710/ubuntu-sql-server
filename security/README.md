# Security

## Overview

This directory contains security documentation for the Ubuntu SQL Server project.

Security controls are used to protect U-Server, PostgreSQL, applications, network services, credentials, and project data.

The goal is to use practical security controls without making services harder to manage than necessary.

## Security Areas

The project focuses on:

- Firewall protection
- SSH security
- File and directory permissions
- Database access
- Secret management
- Network access
- Software updates
- Logging and monitoring
- Backup protection

## Documentation

- `firewalls.md` - Firewall configuration and management
- `permissions.md` - Linux file, directory, and service permissions
- `ssh_hardening.md` - SSH security and hardening practices

## Security Principles

Security changes should follow a few basic rules:

1. Only expose services that need network access.
2. Give users and applications only the permissions they need.
3. Keep passwords, tokens, and other secrets out of Git.
4. Review logs when unexpected activity occurs.
5. Keep software and dependencies updated.
6. Back up important data before major configuration changes.
7. Test security changes before relying on them.
8. Avoid locking out administrative access while changing remote access settings.

## Secrets

Sensitive information should not be stored directly in source code.

Examples include:

- Database passwords
- API keys
- Access tokens
- Private keys
- SSH private keys
- Notification topics
- Authentication credentials

Environment variables or protected configuration files should be used when sensitive values are required.

Example files committed to Git should contain placeholders instead of real credentials.

## Monitoring

Security also depends on knowing when something is not working normally.

U-Server uses logging and monitoring to help identify:

- Service failures
- Database connection problems
- Network problems
- Health warnings
- Critical health conditions
- Notification failures

Related documentation is available under:

```text
logs/
monitoring/
```

## Security Changes

Major security changes should be documented when they affect:

- Firewall rules
- SSH access
- Database access
- Application permissions
- Network exposure
- Authentication
- Service accounts

## Status

Security documentation should continue to be updated as new services and applications are added to U-Server.
