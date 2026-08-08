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
            "2" => println!("Server details cooming soon"),
            "3" => commands::hardware::show_hardware(),
            "4" => commands::network::show_network(),
            "5" => commands::health::show_health(),
            "6" => commands::maintenance::show_maintenance(),

            "0" => {
                printfn!("\nSee ya!");
                break;
            }

            _ => {
                println!("\nInvalid selection. Try again.");
            }
        }

        println!();
    }
}