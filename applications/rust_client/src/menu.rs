// Interactive menu logic will live here.

use std::io::{self, Write};

pub fn show_menu() -> String {
    println!("------------------------------");
    println!(" Ubuntu SQL Server Rust Client");
    println!("------------------------------");
    println!("1. List of servers");
    println!("2. View Server Details");
    println!("3. View Hardware");
    println!("4. View Network");
    println!("5. View Health History");
    println!("6. View Maintenance Logs");
    println!("0. Exit");
    println!("\nSelect an option: ");

    io::stout()
        .flush()
        .expect("Failed to flush stout");

    choice.trim().to_string()
}