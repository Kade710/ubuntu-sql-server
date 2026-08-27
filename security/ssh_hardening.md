# SSH Hardening

## Overview

SSH provides remote command-line access to U-Server.

Because SSH provides administrative access to the system, its configuration should balance security with reliable remote access.

Security changes should be tested carefully to avoid accidentally locking out authorized administrators.

## Service Status

Check the SSH service with:

```bash
systemctl status ssh
```

Check whether SSH is listening with:

```bash
sudo ss -lntp | grep ':22'
```

## Configuration

The main OpenSSH server configuration is normally:

```text
/etc/ssh/sshd_config
```

Additional configuration files may also exist under:

```text
/etc/ssh/sshd_config.d/
```

Before changing SSH configuration, create a backup when appropriate.

## Test Configuration

After editing SSH configuration, test it before restarting the service:

```bash
sudo sshd -t
```

No output normally means the configuration passed the syntax check.

Do not restart SSH after a failed configuration test.

## Restart SSH

After a successful configuration test:

```bash
sudo systemctl restart ssh
```

Check the service afterward:

```bash
systemctl status ssh
```

When working remotely, keep the existing SSH session open while testing a new connection.

This provides a way to fix the configuration if the new connection fails.

## Root Login

Direct root login should normally be disabled:

```text
PermitRootLogin no
```

Administrative work can instead be performed through an authorized user with `sudo`.

## Empty Passwords

Accounts with empty passwords should not be allowed to authenticate through SSH.

Recommended setting:

```text
PermitEmptyPasswords no
```

## Public Key Authentication

SSH supports public key authentication:

```text
PubkeyAuthentication yes
```

SSH keys provide a strong authentication option and remove the need to send an account password during normal key-based authentication.

Private keys must be protected.

## Password Authentication

Password authentication should not be disabled until key-based authentication has been fully configured and tested.

If password authentication is currently required, it may remain enabled while SSH key access is being prepared.

Example:

```text
PasswordAuthentication yes
```

After key authentication is tested from every required administrative device, password authentication can be considered for disabling.

A hardened key-only configuration may use:

```text
PasswordAuthentication no
```

Do not make this change until key-based access has been verified.

## SSH Key Permissions

The SSH directory commonly uses:

```bash
chmod 700 ~/.ssh
```

Private keys commonly use:

```bash
chmod 600 ~/.ssh/<private-key>
```

The authorized keys file commonly uses:

```bash
chmod 600 ~/.ssh/authorized_keys
```

Private keys should never be committed to Git.

## Firewall

SSH normally uses:

```text
22/TCP
```

UFW rules can be checked with:

```bash
sudo ufw status numbered
```

An OpenSSH firewall profile can commonly be allowed with:

```bash
sudo ufw allow OpenSSH
```

Do not remove the working SSH firewall rule while remotely connected unless another tested access path exists.

## Failed Login Attempts

SSH activity can be reviewed through the system journal.

Recent SSH logs can be viewed with:

```bash
journalctl -u ssh -n 50 --no-pager
```

Authentication-related system logs may also provide information about failed login attempts.

Logs should be reviewed when unexpected authentication activity occurs.

## User Access

Only accounts that require remote administration should be able to use SSH.

Unused accounts should not be kept solely for convenience.

Administrative accounts should use strong authentication and should not share credentials.

## Tailscale

Tailscale can provide a private network path for remote access to U-Server.

Using a private network can reduce the need to expose administrative services to broader networks.

SSH authentication and normal Linux permissions should still be maintained even when access occurs through Tailscale.

## Recommended Hardening Process

SSH hardening should be performed gradually:

1. Confirm current SSH access works.
2. Back up the SSH configuration.
3. Configure SSH keys.
4. Test key authentication.
5. Keep the existing session open.
6. Test the SSH configuration with `sshd -t`.
7. Restart SSH.
8. Open a second connection and verify access.
9. Review firewall rules.
10. Consider disabling password authentication only after key access is proven.

## Avoiding Lockouts

Before making major SSH changes:

- Keep an existing SSH session open.
- Make one configuration change at a time.
- Run `sudo sshd -t`.
- Test a second SSH connection.
- Confirm firewall access.
- Make sure local console access is available when possible.

A security setting that prevents authorized administration can create its own recovery problem.

## Future Improvements

Possible future SSH improvements include:

- Full key-based authentication
- Disabling password authentication
- Additional login restrictions
- Automated failed-login protection
- Additional SSH logging and monitoring

Changes should only be documented as active security controls after they have been configured and tested.
