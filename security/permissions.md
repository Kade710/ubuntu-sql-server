# Permissions

## Overview

Linux permissions control which users and groups can read, modify, or execute files and directories.

Correct permissions help protect application code, configuration files, credentials, backups, and server data.

## Viewing Permissions

File permissions can be viewed with:

```bash
ls -l
```

Directory contents and permissions can be viewed with:

```bash
ls -la
```

Example:

```text
-rw-r--r-- 1 user group file.txt
```

The permission groups represent:

```text
owner | group | others
```

## Permission Types

Linux uses three basic permission types:

```text
r = read
w = write
x = execute
```

Directories require execute permission to allow users to enter or access files inside them.

## chmod

Permissions can be changed using `chmod`.

Example:

```bash
chmod +x script.sh
```

This makes a script executable.

Numeric permissions may also be used:

```bash
chmod 644 file.txt
chmod 755 script.sh
```

Permissions should not be made more open than necessary.

Avoid using:

```bash
chmod 777
```

unless there is a specific and understood reason.

## Ownership

File ownership can be viewed with:

```bash
ls -l
```

Ownership can be changed using:

```bash
sudo chown <user>:<group> <file>
```

Recursive ownership changes can be performed with:

```bash
sudo chown -R <user>:<group> <directory>
```

Recursive changes should be used carefully.

## Project Files

Project source files should normally be owned by the development user rather than `root`.

Running development commands with `sudo` can accidentally create root-owned files inside the repository.

If permission problems appear after using `sudo`, check ownership before changing permissions.

## Executable Scripts

Project scripts that need to run directly should have execute permission.

Example:

```bash
chmod +x scripts/backup.sh
chmod +x scripts/restore.sh
```

Scripts that do not need direct execution should not receive unnecessary execute permissions.

## Environment Files

Files containing passwords, tokens, or other secrets should have restricted permissions.

A common private-file permission is:

```bash
chmod 600 .env
```

This allows the owner to read and write the file while preventing access by other users.

Actual permissions should match the users and services that need access.

## SSH Keys

SSH private keys should have restrictive permissions.

A common private key permission is:

```bash
chmod 600 ~/.ssh/<private-key>
```

The `.ssh` directory commonly uses:

```bash
chmod 700 ~/.ssh
```

Public keys do not require the same secrecy as private keys.

## PostgreSQL

PostgreSQL runs under its own service account and maintains protected database files.

Database storage directories should not be manually opened to normal users simply to work around an application permission problem.

Applications should access PostgreSQL through database accounts and permissions rather than direct access to PostgreSQL data files.

## Docker

Docker access should be treated as privileged access.

Users who can control the Docker daemon can perform powerful operations on the host.

Membership in the `docker` group should therefore only be given to trusted users.

## systemd

System-level service files are normally stored under:

```text
/etc/systemd/system/
```

Changing system service definitions normally requires administrative privileges.

After modifying or installing a service file, reload systemd with:

```bash
sudo systemctl daemon-reload
```

## Backups

Backup files should only be readable by users or services that require access.

Backups may contain application data or configuration that should not be publicly accessible.

Backups containing sensitive data should never be committed to Git.

## Git

Before committing files, check that the repository does not contain:

- `.env` files
- Passwords
- Private keys
- Tokens
- Database dumps containing sensitive data
- Runtime backups
- Sensitive logs

`.gitignore` should be used to prevent local and sensitive files from being accidentally committed.

## Troubleshooting

When an application reports a permission error:

1. Identify the file or directory involved.
2. Check its owner.
3. Check its group.
4. Check its permissions.
5. Determine which user runs the application.
6. Change only the permission or ownership that is required.
7. Test the application again.

Avoid immediately using `chmod 777` as a solution.

## Security Principle

Permissions should follow the principle of least privilege.

Users and applications should receive only the access required to perform their intended work.
