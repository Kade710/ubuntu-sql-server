// src/menu.rs

use std::io::{self, Write};

pub fn show_menu() -> io::Result<String> {
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

    print!("\nSelect an option: ");
    io::stdout().flush()?;

    let mut choice = String::new();
    io::stdin().read_line(&mut choice)?;

    Ok(choice.trim().to_string())
}
