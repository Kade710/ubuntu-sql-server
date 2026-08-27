# Linux Groups

## Overview

Linux groups allow permissions to be shared between multiple users.

Groups can control access to files, directories, devices, administrative commands, and services.

## Viewing Groups

Display all known groups:

```bash
getent group
```

Display groups for the current user:

```bash
groups
```

Display detailed user information:

```bash
id
```

A specific user can be checked with:

```bash
id <username>
```

## Primary and Additional Groups

Every Linux user has a primary group.

Users may also belong to additional groups that provide access to other system resources.

Group membership should only be added when required.

## sudo Group

Ubuntu commonly uses the:

```text
sudo
```

group to provide administrative access.

Membership can be checked with:

```bash
getent group sudo
```

Users in this group may be able to execute commands with elevated privileges.

Administrative group membership should be limited to trusted accounts.

## Docker Group

Docker may use the:

```text
docker
```

group to allow users to control Docker without running every command through `sudo`.

Membership can be checked with:

```bash
getent group docker
```

The current user's membership can be checked with:

```bash
groups
```

Docker group access should be treated as privileged access because control of the Docker daemon can provide significant control over the host system.

Only trusted users should belong to this group.

## Adding a User to a Group

A user can be added to an additional group with:

```bash
sudo usermod -aG <group> <username>
```

The `-a` option is important because it appends the group instead of replacing existing additional group memberships.

A new login session may be required before the new group membership becomes active.

## Removing Group Membership

Group membership can be removed with:

```bash
sudo gpasswd -d <username> <group>
```

Access should be tested after changing important group memberships.

## File Groups

Files and directories have both an owner and a group.

View them with:

```bash
ls -l
```

A file's group can be changed with:

```bash
sudo chgrp <group> <file>
```

Ownership and group can be changed together with:

```bash
sudo chown <user>:<group> <file>
```

## Project Permissions

Shared project directories may use group permissions when more than one authorized account needs access.

Group access should only be added when there is a real need for shared access.

The project should not use overly broad permissions simply to avoid ownership problems.

## Service Groups

Some applications create dedicated groups during installation.

These groups may be used to control access to:

- Service files
- Application sockets
- Devices
- Logs
- Configuration files

Service group memberships should not be changed without understanding what access the group provides.

## Security

Group permissions should follow the principle of least privilege.

Before adding a user to a group:

1. Determine what access the group provides.
2. Confirm that the user requires that access.
3. Add the membership.
4. Verify that the required task works.
5. Avoid adding unrelated permissions.

Privileged groups such as `sudo` and `docker` should receive additional review.
