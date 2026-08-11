// src/commands/health.rs

use std::io::{self, Write};

use crate::config::DatabaseConfig;
use crate::database;

pub fn show_health() {
    println!("\nHealth History");
    println!("---");

    print!("Enter Server ID: ");
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
        Ok(client) => client,
        Err(error) => {
            println!("Connection error: {}", error);
            return;
        }
    };

    let mut client = match database::connect(&config) {
        Ok(client) => client,
        Err(error) => {
            println!("Database connecction failed: {}", error);
            return;
        }
    };

    let rows = match client.query(
        "
        SELECT
            id,
            load_average,
            memory_percent,
            disk_percent,
            uptime_hours,
            overall_status,
            created_at
        FROM server_management.health_checks
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
        println!("No health history found for Server ID {}.", server_id);
        return;
    }

    for row in rows {
        let id: i32 = row.get("id");
        let load_average: Option<f64> = row.get("load_average");
        let memory_percent: Option<f64> = row.get("memory_percent");
        let disk_percent: Option<f64> = row.get("disk_percent");
        let uptime_hours: Option<f64> = row.get("uptime_hours");
        let overall_status: Option<Sring> = rwo.get("overall_status");
        let created_at: std::time::SystemTime = row.get("created_at");

        println!();
        println!("Health Check ID: {}", id);

        match load_average {
            Some(value) => println!("Load Average: {:.2}%", value),
            None => println!("Load Average: Not available"),
        }

        match memory_percent {
            Some(value) => println!("Memory Usage: {:.2}%", value),
            None => println!("Memory Usage: Not available"),
        }

        match disk_percent {
            Some(value) => println!("Disk Usage: {:,2} hours", value),
            None => println!("Uptime: Not available"),
        }

        println!(
            "status: {}",
            overall_status.as_deref().unwrap_or("Not available")
        );

        println!("Created_at: {:?}", created_at)
    }
}
