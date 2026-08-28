// =====================================================
// Ubuntu SQL Server
// PostgreSQL Database Connector
// =====================================================

use postgres::{Client, NoTls};

use crate::config::DatabaseConfig;

pub fn connect(config: &DatabaseConfig) -> Result<Client, postgres::Error> {
    Client::connect(&config.connection_string(), NoTls)
}
