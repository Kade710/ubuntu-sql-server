# Users and Permissions

## Overview

This directory documents user accounts, groups, database roles, and access controls used by the Ubuntu SQL Server project.

Linux and PostgreSQL use separate permission systems. Linux users and groups control access to the operating system, files, services, and applications. PostgreSQL roles control access to databases and database objects.

## Documentation

- `linux_users.md` - Linux user accounts and system accounts
- `groups.md` - Linux groups and group permissions
- `database_roles.md` - PostgreSQL roles and database permissions

## Access Control

Access should follow the principle of least privilege.

Users, applications, and services should only receive the permissions required to perform their intended tasks.

Administrative privileges should not be used when normal user permissions are enough.

## Main Permission Areas

Access control may apply to:

- Linux user accounts
- Linux groups
- Files and directories
- systemd services
- Docker
- SSH
- PostgreSQL
- Application configuration
- Environment files
- Backup files

## Security

Passwords and other credentials should never be stored in this documentation.

Do not commit:

- Linux passwords
- Database passwords
- SSH private keys
- API keys
- Access tokens
- Secret environment variables

Example commands should use placeholders when credentials are required.

## Auditing Access

Linux account information can be reviewed with:

```bash
getent passwd
```

Groups can be reviewed with:

```bash
getent group
```

The current user's groups can be viewed with:

```bash
groups
```

PostgreSQL roles can be viewed inside `psql` with:

```sql
\du
```

Permissions should be reviewed when new users, services, applications, or databases are added.

## Related Documentation

Additional security information is available under:

```text
security/
```

Environment variable documentation is available under:

```text
environment_variables/
```

## Status

This documentation should be updated whenever user accounts, groups, database roles, or important permissions change.
