# Rust Client

The Rust Client is a command-line application for viewing U-Server information stored in the PostgreSQL database.

It provides a lightweight terminal interface for accessing server management information without using the web dashboard.

## Features

- List registered servers
- View detailed server information
- View hardware components
- View network interfaces
- View health history
- View maintenance logs
- Validate server ID input
- Handle database and configuration errors gracefully
- Display health and maintenance timestamps in a human-readable format

## Technologies

- Rust
- PostgreSQL
- Chrono

## Project Structure

```text
rust_client/
├── src/
│   ├── commands/
│   │   ├── hardware.rs
│   │   ├── health.rs
│   │   ├── maintenance.rs
│   │   ├── mod.rs
│   │   ├── network.rs
│   │   └── servers.rs
│   ├── config.rs
│   ├── database.rs
│   ├── main.rs
│   ├── menu.rs
│   └── models.rs
├── Cargo.lock
├── Cargo.toml
└── README.md
