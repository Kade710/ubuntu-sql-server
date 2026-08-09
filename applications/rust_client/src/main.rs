mod commands;
mod config;
mod database;
mod menu;
mod models;

fn main() {
    loop {
        let choice = menu::show_menu();

        match choice.as_str() {
            "1" => commands::servers::show_servers(),
            "2" => commands::servers::show_server_details(),
            "3" => commands::hardware::show_hardware(),
            "4" => commands::network::show_network(),
            "5" => commands::health::show_health(),
            "6" => commands::maintenance::show_maintenance(),

            "0" => {
                println!("\nSee ya!");
                break;
            }

            _ => {
                println!("\nInvalid selection. Try again.");
            }
        }

        println!();
    }
}
