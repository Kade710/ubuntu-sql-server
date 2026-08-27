# Troubleshooting

This document contains common problems and troubleshooting steps for the Ubuntu SQL Server project.

## PostgreSQL Connection Problems

If an application cannot connect to PostgreSQL:

- Check that PostgreSQL is running.
- Verify that the PostgreSQL cluster is online.
- Confirm the database host and port.
- Check database credentials.
- Review PostgreSQL access permissions.
- Check firewall rules if the connection is coming from another system.

Useful commands:

```bash
sudo systemctl status postgresql
pg_lsclusters