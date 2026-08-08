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
            .unwrap_or_else(|_| "localhost".to_string());

        let port = env::var("DB_PORT")
            .unwrap_or_else(|_| "5432".to_string())
            .parse::<u16>()
            .map_err(|_| "DB_PORT must be a valid number".to_string())?;

        let database =
            env::var("DB_NAME").map_err(|_| "DB_NAME is required".to_string())?;

        let user =
            env::var("DB_USER").map_err(|_| "DB_USER is required".to_string())?;

        let password =
            env::var("DB_PASSWORD").map_err(|_| "DB_PASSWORD is required".to_string())?;

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
            self.host,
            self.port,
            self.database,
            self.user,
            self.password
        )
    }
}
