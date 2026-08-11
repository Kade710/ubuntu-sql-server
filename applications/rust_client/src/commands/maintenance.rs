// src/commands/maintenance.rs

use std::io::{self, Write};

use crate::config::DatabaseConfig;
use crate::database;

pub fn show_maintenance() {
    println!("Maintenance Logs");
    println!("---")

    print!("Enter Server ID");
    io::stdout()
        .flush()
        .expect("Failed to flush stdout");

    let mut input = String::new();

    io::stdin()
        .read_line(&mut input)
        .expect("Failed to read input");

    let server_id: i32 = match input.trim().parse() {
        Ok(id) => id,
        Err(_) => {
            println!("Invalid server ID");
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
        ORDER BY by created_at DESC
        LIMIT 10
        ".
        &[&server_id],
    ) {
        Ok(rows) => rows,
        Err(error) => {
            println!("Database query failed: {}", error);
            return;
        }
    };

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
            "Desciption: {}",
            desciption.as_deref().unwrap_or("Not available")
        );
        println!(
            "Performed By: {}",
            performed_by.as_deref().unwrap_or("Not available")
        );
        println!("Created At: {:?}", created_at);
    }
}
