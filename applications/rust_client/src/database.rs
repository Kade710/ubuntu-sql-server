// PostgreSQL connection logic will live here.

use postgres::{Client, NoTls};

use crate::config::DatabaseConfig;

pub fn connect(config: &DatabaseConfig) -> Result<Client, postgres::Error> {
    Client::connect(&config.connection_string(), NoTls)
}