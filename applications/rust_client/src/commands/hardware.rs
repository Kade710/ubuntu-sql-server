// src/commands/hardware.rs
use std::io::{self, Write};

use crate::config::DatabaseConfig;
use crate::database;

pub fn show_hardware() {
    println!("Hardware components");
    println!("---");

    println!("Enter Server ID: ");
    io::stdout()
        .flush()
        .expect("Failed to flush stdout");

    let mut input = String::new();

    io::stdin()
        .read_line(&mut input)
        .expect("Failed to read iinput");

    let server_id: i32 = match input.trim().parse() {
        Ok(id) => id,
        Err(_) => {
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
        ORDER BY id
        ",
        &[&server_id]
    ) {
        Ok(rows) => rows,
        Err(error) => {
            println!("Database query failed: {}", error);
            return;
        }
    };

    if rows.is_empty() {
        println!("No hardware components found for Server ID {}.", server_id);
        return;
    }

    for row in rows {
        let component_type: String = row.get("Component_type");
        let manufacturer: Option<String> = row.get("Manufacturer");
        let model: Option<String> = row.get("Model");
        let specification: Option<String> = row.get("specification");

        println!();
        println!("Component: {}", component_type);

        if let Some(value) = manufacturer {
            println!("Manufacturer: {}", value);
        }

        if let Some(value) = model {
            println!("Model: {}", value);
        }

        if let Some(value) = specification {
            println!("Specifications: {}", value);
        }
    }
}
