// src/commands/maintenance.rs

use std::io::{self, Write};

use crate::config::DatabaseConfig;
use crate::database;

pub fn show_maintenance() {
    println!("\nMaintenance Logs");
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
            id,
            action,
            description,
            performed_by,
            created_at
        FROM server_management.maintenance_logs
        WHERE server_id = $1
        ORDER BY created_at DESC
        LIMIT 10
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
        println!("No maintenance logs found for Server ID {}.", server_id);
        return;
    }

    for row in rows {
        let id: i32 = row.get("id");
        let action: String = row.get("action");
        let description: Option<String> = row.get("description");
        let performed_by: Option<String> = row.get("performed_by");
        let created_at: std::time::SystemTime = row.get("created_at");

        println!();
        println!("Maintenance Log ID: {}", id);
        println!("Action: {}", action);
        println!(
            "Description: {}",
            description.as_deref().unwrap_or("Not available")
        );
        println!(
            "Performed By: {}",
            performed_by.as_deref().unwrap_or("Not available")
        );
        println!("Created At: {:?}", created_at);
    }
}
