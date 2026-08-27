# PostgreSQL Database Roles

## Overview

PostgreSQL uses roles to control database authentication, ownership, and permissions.

PostgreSQL roles are separate from Linux user accounts.

A Linux user does not automatically receive PostgreSQL access simply because the account exists on U-Server.

## Project Database

The main project database is:

```text
ubuntu_sql_server
```

The main project schema is:

```text
server_management
```

## Current Project Roles

The project currently uses PostgreSQL roles including:

```text
postgres
jonathon_admin
```

These roles have different purposes.

---

## postgres Role

The PostgreSQL installation includes the administrative role:

```text
postgres
```

This role is used for PostgreSQL administration.

Administrative access can be opened with:

```bash
sudo -u postgres psql
```

The project database can be opened with:

```bash
sudo -u postgres psql -d ubuntu_sql_server
```

The `postgres` role should generally be reserved for database administration rather than normal application access.

---

## jonathon_admin Role

The project uses:

```text
jonathon_admin
```

for project database access and ownership of several database objects.

Applications can use a dedicated PostgreSQL role instead of connecting as the PostgreSQL superuser.

The role's password should never be stored in documentation or committed to Git.

---

## Viewing Roles

Inside `psql`, list PostgreSQL roles with:

```sql
\du
```

More information can be queried with:

```sql
SELECT rolname,
       rolsuper,
       rolcreatedb,
       rolcreaterole,
       rolcanlogin
FROM pg_roles
ORDER BY rolname;
```

This helps identify which roles have administrative privileges and which roles can log in.

---

## Database Ownership

Database ownership can be checked inside `psql` with:

```sql
\l
```

A specific database can also be inspected through PostgreSQL system information.

Ownership should be assigned intentionally.

Applications normally do not require PostgreSQL superuser privileges.

---

## Schema Ownership

Schema information can be viewed with:

```sql
\dn+
```

The project schema is:

```text
server_management
```

Schema ownership affects who can create or modify objects inside the schema.

---

## Table Ownership

Project tables can be listed with:

```sql
\dt server_management.*
```

Detailed information about a table can be viewed with:

```sql
\d server_management.<table-name>
```

Table ownership can also be queried with:

```sql
SELECT schemaname,
       tablename,
       tableowner
FROM pg_tables
WHERE schemaname = 'server_management'
ORDER BY tablename;
```

This is useful for finding ownership differences between project tables.

---

## Permissions

PostgreSQL permissions may include:

```text
SELECT
INSERT
UPDATE
DELETE
TRUNCATE
REFERENCES
TRIGGER
CONNECT
CREATE
USAGE
EXECUTE
```

A role should only receive the permissions required for its purpose.

---

## GRANT

Permissions can be provided using `GRANT`.

Example:

```sql
GRANT SELECT ON server_management.<table-name> TO <role>;
```

Multiple permissions can be granted together:

```sql
GRANT SELECT, INSERT, UPDATE
ON server_management.<table-name>
TO <role>;
```

Only grant permissions that the role actually requires.

---

## REVOKE

Permissions can be removed with `REVOKE`.

Example:

```sql
REVOKE INSERT
ON server_management.<table-name>
FROM <role>;
```

Permissions should be reviewed after major application or role changes.

---

## Schema Access

A role may require access to the project schema.

Example:

```sql
GRANT USAGE
ON SCHEMA server_management
TO <role>;
```

Schema access and table permissions are separate.

Giving a role access to a schema does not automatically give it full access to every table.

---

## Sequence Permissions

Tables using generated IDs may depend on PostgreSQL sequences.

Applications inserting records may require appropriate sequence permissions.

Sequences can be listed with:

```sql
\ds server_management.*
```

Permissions should be granted only when required by the application's database operations.

---

## Application Access

Applications such as the Go Agent and Django Web Dashboard should use appropriate application credentials rather than relying on PostgreSQL superuser access.

Database configuration may include:

```text
DB_HOST
DB_PORT
DB_NAME
DB_USER
DB_PASSWORD
```

Real credentials should be stored in protected environment configuration.

Example environment files committed to Git should contain placeholders only.

---

## Ownership Problems

Database ownership can affect whether an application is able to modify tables, sequences, or other objects.

When a permission problem occurs, check:

1. Which PostgreSQL role the application uses.
2. Who owns the database.
3. Who owns the schema.
4. Who owns the affected table.
5. Who owns related sequences.
6. Which privileges have been granted.

Do not automatically give the application superuser privileges to solve a permission problem.

Fix the specific ownership or permission issue instead.

---

## Checking Current User

Inside PostgreSQL:

```sql
SELECT current_user;
```

The current database can be checked with:

```sql
SELECT current_database();
```

Connection information can be viewed with:

```sql
\conninfo
```

These commands are useful when troubleshooting application permissions.

---

## Security

Database roles should follow the principle of least privilege.

Recommended practices include:

- Do not use `postgres` for normal applications.
- Do not give applications unnecessary superuser privileges.
- Store passwords outside source code.
- Review database ownership.
- Review schema and table permissions.
- Remove unused roles when appropriate.
- Use separate roles when applications require different access levels.

## Future Improvements

As the project grows, additional database roles may be created for:

- Read-only access
- Web Dashboard access
- Monitoring
- Backup operations
- Reporting
- Additional applications

New roles should be documented after they are created and tested.
