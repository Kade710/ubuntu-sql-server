# Security Baseline

## Purpose

This document defines the basic security practices used by the Ubuntu SQL Server project.

The goal is to reduce security risks while keeping the server and its services manageable.

## Access Control

- Limit server access to authorized users.
- Use SSH for remote administration.
- Use separate database roles when appropriate.
- Avoid giving applications more permissions than they need.

## Secrets

Passwords, API keys, tokens, and other secrets should not be stored directly in source code.

Environment files containing secrets should be excluded from Git.

Example environment files should contain placeholders instead of real credentials.

## Network Security

The server firewall should only allow ports that are required for active services.

Unused services and ports should remain blocked.

Remote access should be limited whenever possible.

## Database Security

Database users should only receive the permissions required for their role.

Database credentials should be stored outside of source code.

## Updates

The operating system, applications, dependencies, and security packages should be kept up to date.

Updates should be reviewed and tested when necessary before being used in important environments.

## Backups

Important data should be backed up when appropriate.

Backups containing sensitive information should be protected and should not be committed to Git.

Restore procedures should be tested to make sure backups can be used successfully.

## Monitoring

Server health and important system events should be monitored.

Warnings, critical conditions, and recovery events should be recorded when supported by the monitoring system.

## Review

This security baseline should be updated as new services, applications, and security controls are added to the project.
