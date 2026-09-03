# User Access Management

## Overview

The U-Server web dashboard provides centralized user access management using Django authentication, role-based access control (RBAC), SSH public key management, and access audit logging.

The access-management system is designed to separate dashboard permissions from operating-system privileges and to avoid storing passwords or private SSH keys in audit records.

---

## Roles and Permissions

U-Server uses Django groups as application roles.

### Administrator

Permissions:

- `manage_users`
- `manage_ssh_access`
- `manage_servers`

Administrators can manage dashboard users, SSH access, and servers.

### Operator

Permissions:

- `manage_ssh_access`
- `manage_servers`

Operators can manage SSH access and servers but cannot manage dashboard users.

### Viewer

Viewer accounts have no management permissions.

They are intended for access that does not require administrative changes.

---

## User Management

Authorized administrators can manage users through the web dashboard or user-management API.

Supported operations include:

- Create a user
- Disable a user
- Assign a managed role
- View user status

The managed roles are:

- Administrator
- Operator
- Viewer

Role assignment replaces the user's previous managed role.

Passwords are processed through Django's authentication system and are stored as password hashes rather than plaintext passwords.

---

## User Onboarding

A standard onboarding sequence is:

1. Create the dashboard user.
2. Assign the appropriate role.
3. Verify the user's permissions.
4. Register an SSH public key if SSH access is required.
5. Verify that the expected audit events were created.

User creation and role assignment are recorded in the access audit log.

---

## User Revocation

A standard revocation sequence is:

1. Disable the dashboard user.
2. Revoke registered SSH keys when applicable.
3. Verify that the SSH authorized-keys file was synchronized.
4. Verify the corresponding access audit events.

Disabling a Django account prevents that account from authenticating through the normal Django authentication system.

SSH access is managed separately and must be revoked separately when a user has registered SSH access.

---

## SSH Access Architecture

Dashboard-managed SSH access uses a dedicated Linux account:

`uaccess`

The Django application does not manage the personal or administrative SSH keys belonging to the `jonathon` Linux account.

The `uaccess` account:

- Has no sudo privileges
- Has its password locked
- Uses public-key authentication
- Cannot modify its own `authorized_keys` file

The server SSH configuration retains:

- `PermitRootLogin no`
- `PasswordAuthentication no`

These settings must not be weakened to support dashboard-managed access.

---

## SSH Public Key Registration

The dashboard accepts supported OpenSSH public keys.

Supported key types include:

- `ssh-ed25519`
- `ssh-rsa`
- `ecdsa-sha2-nistp256`
- `ecdsa-sha2-nistp384`
- `ecdsa-sha2-nistp521`

The dashboard calculates and stores a SHA-256 fingerprint for each registered key.

The public key is stored because it is required to construct the `authorized_keys` file.

Private SSH keys are never required and must never be submitted to or stored by the dashboard.

Duplicate SSH keys are rejected using the stored fingerprint.

---

## SSH Key Synchronization

Active public keys are written to the staging file:

`/var/lib/ubuntu-sql-server/ssh/uaccess.keys`

The Django application invokes the fixed synchronization helper:

`/usr/local/sbin/u-server-sync-uaccess-keys`

The helper installs the authorized keys for the dedicated `uaccess` account.

The live file is:

`/home/uaccess/.ssh/authorized_keys`

The live SSH directory and authorized-keys file are controlled by root so that `uaccess` cannot modify its own authentication configuration.

The current live file permissions allow `uaccess` to read the authorized keys while preventing it from changing or replacing them.

---

## Privilege Separation

Django does not run as root.

The application uses a narrowly scoped sudo rule allowing the application host user to execute only the SSH synchronization helper without a password.

The sudoers configuration is stored in:

`/etc/sudoers.d/ubuntu-sql-server-ssh`

The permitted command is:

`/usr/local/sbin/u-server-sync-uaccess-keys`

General unrestricted sudo access is not granted to the Django application.

---

## SSH Key Revocation

When a key is revoked:

1. The database record is marked inactive.
2. A revocation timestamp is recorded.
3. The active-key set is synchronized with the operating system.
4. The revoked key is removed from the live `authorized_keys` file.
5. An audit event is created after successful synchronization.

If synchronization fails, the database change is rolled back and the key remains active.

A failed synchronization does not create a successful revocation audit event.

---

## Access Audit Logging

Security-sensitive access-management operations are recorded in the `AccessAuditLog` model.

Recorded actions include:

- `USER_CREATED`
- `USER_DISABLED`
- `USER_ROLE_CHANGED`
- `SSH_KEY_REGISTERED`
- `SSH_KEY_REVOKED`

Audit records contain:

- Actor
- Action
- Target type
- Target identifier
- Descriptive details
- Timestamp

Passwords and private SSH key material must never be written to the audit log.

SSH audit records identify keys using their fingerprint rather than copying the complete public key into the audit record.

For SSH registration and revocation, successful audit events are written only after operating-system key synchronization succeeds.

---

## Verification

User onboarding and revocation were verified through the application workflow.

The verified user lifecycle was:

1. Create user
2. Assign Operator role
3. Disable user

The resulting audit sequence was:

- `USER_CREATED`
- `USER_ROLE_CHANGED`
- `USER_DISABLED`

SSH registration and revocation were previously verified against the dedicated `uaccess` account.

Automated tests mock the operating-system SSH synchronization function so the Django test suite cannot modify the production `uaccess` authorized-keys file or invoke the live synchronization helper.

The Version 4.0 access-management test suite verifies:

- User creation
- Role assignment
- User disabling
- User audit events
- Browser-based user-management audit events
- Password exclusion from audit records
- SSH key registration auditing
- SSH key revocation auditing
- SSH synchronization failure rollback
- Prevention of false audit events after synchronization failures
- SSH permission enforcement

At final verification, the dashboard test suite contained 25 tests and completed successfully.

---

## Security Notes

Do not:

- Store plaintext passwords
- Store private SSH keys
- Give `uaccess` sudo privileges
- Allow `uaccess` to modify its own `authorized_keys`
- Run Django as root
- Broaden the SSH synchronization sudo rule unnecessarily
- Disable public-key-only SSH protections

Future security hardening may further restrict role assignment, protect critical administrator accounts, and strengthen application configuration.
