// src/commands/serever.rs

use crate::config::DatabaseConfig;
use crate::database;
use crate::models::Server;

pub fn show_servers() {
    println!("\n Registered Servers");
    println!("---")

    let config = match DatabaseConfig::load() {
        Ok(client) => client,
        Err(error) => {
            println!("Configuration error: {}, error");
            return;
        }
    };

    let mut client = match database::connect(&config) {
        Ok(client) => client,
        Err(error) => {
            println!("Database connection failed: {}, error")
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
            println!("Database query failed:", error);
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
            "IP Address:{}",
            server.ip_address.as_deref().unwrap_or("Not available")
        );

        println!(
            "Operating System: {}",
            server
                .operating_system
                .as_deref()
                .unwrap_or("Not availabe")
        );
    }
}
