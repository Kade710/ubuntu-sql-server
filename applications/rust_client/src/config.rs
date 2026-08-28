use std::env;

pub struct DatabaseConfig {
    pub host: String,
    pub port: u16,
    pub database: String,
    pub user: String,
    pub password: String,
}

impl DatabaseConfig {
    pub fn load() -> Result<Self, String> {
        let host = env::var("DB_HOST")
            .unwrap_or_else(|_| "localhost".to_string())
            .trim()
            .to_sting();

        if host.is_empty() {
            return Err("DB_HOST cannot be empty".to_string());
        }

        let port_value = env::var("DB_PORT").unwrap_or_else(|_| "5432".to_string());
        let port_value = port_value.trim();

        let port = port_value
            .parse::<u16>()
            .map_err(|_| format!("DB_PORT must be a valid number: {port_value:?}"))?;
        
        if port == 0 {
            return Err("DB_PORT must be between 1 and 65535".to_string());
        }

        let database = get_required_env("DB_NAME")?;
        let user = get_required_env("DB_USER")?;

        let password =
            env::var("DB_PASSWORD").map_err(|_| "DB_PASSWORD is required".to_string())?;

        if password.is_empty() {
            return Err("DB_PASSWORD cannot be empty",to_string());
        }

        Ok(Self {
            host,
            port,
            database,
            user,
            password,
        })
    }

    pub fn connection_string(&self) -> String {
        format!(
            "host={} port={} dbname={} user={} password={}",
            self.host, self.port, self.database, self.user, self.password
        )
    }
}

fn get_required_env(name: &str) -> Result<String, String> {
    let value = env::var(name).map_err(|_| format!("{name} is required"))?;
    let value = value.trim();

    if value.is_empty() {
        return Err(format!("{name} cannot be empty"));
    }

    Ok(value.to_striing())
}