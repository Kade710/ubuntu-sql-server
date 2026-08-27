# Linux Users

## Overview

Linux user accounts control access to the operating system, files, applications, services, and administrative commands.

U-Server uses both normal user accounts and system service accounts.

## Current User

The primary development account is:

```text
jonathon
```

This account is used for:

- Project development
- Git operations
- Go development
- Python and Django development
- Server administration
- SSH access
- Application testing

Project files are normally stored under:

```text
/home/jonathon/Projects/
```

The Ubuntu SQL Server repository is located at:

```text
/home/jonathon/Projects/ubuntu-sql-server/
```

## Root

Linux provides the administrative account:

```text
root
```

Direct use of the root account should be limited.

Administrative commands should normally be performed through `sudo` when elevated privileges are required.

Example:

```bash
sudo systemctl status postgresql
```

## PostgreSQL System Account

PostgreSQL uses the Linux service account:

```text
postgres
```

This account is separate from PostgreSQL database roles.

Administrative PostgreSQL commands may be run with:

```bash
sudo -u postgres psql
```

The project database can be opened with:

```bash
sudo -u postgres psql -d ubuntu_sql_server
```

The Linux `postgres` account should not be used as a normal development account.

## Service Accounts

Linux services may run under dedicated system accounts.

This limits the amount of access a service has to the rest of the operating system.

Service account configuration can be inspected through systemd.

For example:

```bash
systemctl cat <service-name>
```

The `User=` and `Group=` settings can identify which account runs a service when those settings are defined.

## Viewing Users

All known accounts can be viewed with:

```bash
getent passwd
```

The current account can be displayed with:

```bash
whoami
```

Current login information can be viewed with:

```bash
id
```

Example:

```bash
id jonathon
```

This displays the user's UID, primary group, and additional group memberships.

## Home Directories

Normal user accounts commonly have home directories under:

```text
/home/
```

For example:

```text
/home/jonathon/
```

Home directories may contain:

- SSH configuration
- Git configuration
- Project repositories
- Application configuration
- Shell configuration
- Local development tools

Permissions should prevent unnecessary access by other users.

## sudo Access

Users with authorized administrative access may use `sudo`.

Current sudo permissions can be checked with:

```bash
sudo -l
```

Administrative privileges should only be granted when required.

## File Ownership

Files created during normal development should generally be owned by the development user.

Check ownership with:

```bash
ls -l
```

Ownership can be changed when necessary with:

```bash
sudo chown <user>:<group> <file>
```

Avoid running normal development commands with `sudo` because this may create root-owned files inside the project.

## SSH Access

Linux accounts may also be used for SSH authentication.

Only accounts that require remote access should be configured for SSH.

SSH configuration and hardening are documented under:

```text
security/ssh_hardening.md
```

## Removing Users

Accounts should not be removed without first checking:

- File ownership
- Running services
- Scheduled jobs
- SSH access
- Application dependencies
- Group memberships

A service or application may stop working if it depends on the removed account.

## Security

User accounts should follow these guidelines:

- Use strong authentication.
- Do not share account credentials.
- Limit administrative privileges.
- Remove unused accounts when appropriate.
- Protect home directories.
- Review SSH access.
- Use dedicated service accounts when appropriate.
- Avoid running applications as root without a specific reason.
