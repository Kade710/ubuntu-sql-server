# Rust Client

The Rust Client is a command-line application for viewing U-Server information stored in the PostgreSQL database.

It provides a simple menu for accessing server information without using the web dashboard.

## Features

- List registered servers
- View server details
- View hardware information
- View network information
- View health history
- View maintenance logs

## Database

The client reads server management data from the `ubuntu_sql_server` PostgreSQL database.

Most server information is stored in the `server_management` schema.

## Technologies

- Rust
- PostgreSQL

## Status

The Rust Client is working and actively being developed. Additional features may be added as the Ubuntu SQL Server project grows.
