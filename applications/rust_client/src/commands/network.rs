// src/commands/netnork.rs

use std::io::{self, Write};

use crate::config::DatabaseConfig;
use crate::database;

pub fn show_network() {
    println!("\nNetwork Interfaces")
    println!("---");

    println!("Enter Server ID: ");
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
            interface_name,
            mac_address,
            network_type,
            speed_mbps
        FROM server_management.network_interfaces
        WHERE server_id = $1
        ORDER BY id
        ",
        &[&server_id],
    ) {
        Ok(rows) => rows,
        Err(error) => {
            println!(
                "No network interfaces found for Server ID {}.",
                server_id
            );
            return;
        }

        for row in rows {
            let interface_name: String = row.get("interface_name");
            let mac_address: Option<String> = row.get("mac_address");
            let ip_address: Option<String> = row.get("ip_address");
            let network_type: Option<String> = row.get("network_type");
            let speed_mbps: Option<i32> = row.get("speed_mbps");

        println!();
        println!("Interface: {}", interface_name);
        println!(
            "MAC Address: {}",
            mac_address.as_deref().unwrap_or("Not available")
        );
        println!(
            "IP Address: {}",
            ip_address.as_deref().unwrap_or("Not available")
        );
        println!(
            "Network Type: {}",
            network_type.as_deref().unwrap_or("Not available")
        );

        match speed_mbps {
            Some(speed) => println!("Speed: {} Mbps", speed),
            None => println!("Speed: Not available"),
        }
    }
}
