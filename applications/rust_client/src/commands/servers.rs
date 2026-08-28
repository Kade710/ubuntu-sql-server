// src/commands/servers.rs

use std::io::{self, Write};

use crate::config::DatabaseConfig;
use crate::database;
use crate::models::Server;

pub fn show_servers() {
    println!("\nRegistered Servers");
    println!("---");

    let config = match DatabaseConfig::load() {
        Ok(config) => config,
        Err(error) => {
            println!("Configuration error: {}", error);
            return;
        }
    };

    let mut client = match database::connect(&config) {
        Ok(client) => client,
        Err(error) => {
            println!("Database connection failed: {}", error);
            return;
        }
    };

    let rows = match client.query(
        "
        SELECT
            id,
            hostname,
            ip_address,
            operating_system
        FROM server_management.server_inventory
        ORDER BY id
        ",
        &[],
    ) {
        Ok(rows) => rows,
        Err(error) => {
            println!("Database query failed: {}", error);
            return;
        }
    };

    if rows.is_empty() {
        println!("No registered servers found.");
        return;
    }

    for row in rows {
        let server = Server {
            id: row.get("id"),
            hostname: row.get("hostname"),
            ip_address: row.get("ip_address"),
            operating_system: row.get("operating_system"),
        };

        println!();
        println!("Server ID: {}", server.id);
        println!("Hostname: {}", server.hostname);
        println!(
            "IP Address: {}",
            server.ip_address.as_deref().unwrap_or("Not available")
        );
        println!(
            "Operating System: {}",
            server
                .operating_system
                .as_deref()
                .unwrap_or("Not available")
        );
    }
}

pub fn show_server_details() {
    println!("\nServer Details");
    println!("---");

    print!("Enter Server ID: ");

    if let Err(error) = io::stdout().flush() {
        println!("Failed to flush output: {}", error);
        return;
    }

    let mut input = String::new();

    if let Err(error) = io::stdin().read_line(&mut input) {
        println!("Failed to read input: {}", error);
        return;
    }

    let server_id: i32 = match input.trim().parse() {
        Ok(id) if id > 0 => id,
        _ => {
            println!("Invalid server ID.");
            return;
        }
    };

    let config = match DatabaseConfig::load() {
        Ok(config) => config,
        Err(error) => {
            println!("Configuration error: {}", error);
            return;
        }
    };

    let mut client = match database::connect(&config) {
        Ok(client) => client,
        Err(error) => {
            println!("Database connection failed: {}", error);
            return;
        }
    };

    let row = match client.query_opt(
        "
        SELECT
            id,
            hostname,
            ip_address,
            operating_system,
            cpu,
            ram_gb,
            storage_gb,
            gpu,
            motherboard
        FROM server_management.server_inventory
        WHERE id = $1
        ",
        &[&server_id],
    ) {
        Ok(row) => row,
        Err(error) => {
            println!("Database query failed: {}", error);
            return;
        }
    };

    let Some(row) = row else {
        println!("Server ID {} was not found.", server_id);
        return;
    };

    let id: i32 = row.get("id");
    let hostname: String = row.get("hostname");
    let ip_address: Option<String> = row.get("ip_address");
    let operating_system: Option<String> = row.get("operating_system");
    let cpu: Option<String> = row.get("cpu");
    let ram_gb: Option<i32> = row.get("ram_gb");
    let storage_gb: Option<i32> = row.get("storage_gb");
    let gpu: Option<String> = row.get("gpu");
    let motherboard: Option<String> = row.get("motherboard");

    println!();
    println!("Server ID: {}", id);
    println!("Hostname: {}", hostname);
    println!(
        "IP Address: {}",
        ip_address.as_deref().unwrap_or("Not available")
    );
    println!(
        "Operating System: {}",
        operating_system.as_deref().unwrap_or("Not available")
    );
    println!("CPU: {}", cpu.as_deref().unwrap_or("Not available"));

    match ram_gb {
        Some(value) => println!("RAM: {} GB", value),
        None => println!("RAM: Not available"),
    }

    match storage_gb {
        Some(value) => println!("Storage: {} GB", value),
        None => println!("Storage: Not available"),
    }

    println!("GPU: {}", gpu.as_deref().unwrap_or("Not available"));
    println!(
        "Motherboard: {}",
        motherboard.as_deref().unwrap_or("Not available")
    );
}