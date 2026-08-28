// src/commands/health.rs

use chrono::{DateTime, Local};
use std::io::{self, Write};

use crate::config::DatabaseConfig;
use crate::database;

pub fn show_health() {
    println!("\nHealth History");
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
            load_average::DOUBLE PRECISION AS load_average,
            memory_percent::DOUBLE PRECISION AS memory_percent,
            disk_percent::DOUBLE PRECISION AS disk_percent,
            uptime_hours::DOUBLE PRECISION AS uptime_hours,
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
        let overall_status: Option<String> = row.get("overall_status");
        let created_at: std::time::SystemTime = row.get("created_at");
        let created_at: DateTime<Local> = created_at.into();

        println!();
        println!("Health Check ID: {}", id);

        match load_average {
            Some(value) => println!("Load Average: {:.2}", value),
            None => println!("Load Average: Not available"),
        }

        match memory_percent {
            Some(value) => println!("Memory Usage: {:.2}%", value),
            None => println!("Memory Usage: Not available"),
        }

        match disk_percent {
            Some(value) => println!("Disk Usage: {:.2}%", value),
            None => println!("Disk Usage: Not available"),
        }

        match uptime_hours {
            Some(value) => println!("Uptime: {:.2} hours", value),
            None => println!("Uptime: Not available"),
        }

        println!(
            "Status: {}",
            overall_status.as_deref().unwrap_or("Not available")
        );

        println!(
            "Created At: {}",
            created_at.format("%Y-%m-%d %H:%M:%S")
        );
    }
}
