// src/commands/hardware.rs

use std::io::{self, Write};

use crate::config::DatabaseConfig;
use crate::database;

pub fn show_hardware() {
    println!("\nHardware Components");
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

    let rows = match client.query(
        "
        SELECT
            component_type,
            manufacturer,
            model,
            specification
        FROM server_management.hardware_components
        WHERE server_id = $1
        ORDER BY component_type
        ",
        &[&server_id],
    ) {
        Ok(rows) => rows,
        Err(error) => {
            println!("Database query failed: {}", error);
            return;
        }
    };

    if rows.is_empty() {
        println!(
            "No hardware components found for Server ID {}.",
            server_id
        );
        return;
    }

    for row in rows {
        let component_type: String = row.get("component_type");
        let manufacturer: Option<String> = row.get("manufacturer");
        let model: Option<String> = row.get("model");
        let specification: Option<String> = row.get("specification");

        println!();
        println!("Component: {}", component_type);
        println!(
            "Manufacturer: {}",
            manufacturer.as_deref().unwrap_or("Not available")
        );
        println!(
            "Model: {}",
            model.as_deref().unwrap_or("Not available")
        );
        println!(
            "Specifications: {}",
            specification.as_deref().unwrap_or("Not available")
        );
    }
}