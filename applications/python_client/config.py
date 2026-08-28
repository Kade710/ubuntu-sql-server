# =====================================================
# Ubuntu SQL Server
# Python Client Configuration
# =====================================================

import os

from dotenv import load_dotenv

BASE_DIR = os.path.dirname(os.path.abspath(__file__))

load_dotenv(os.path.join(BASE_DIR, ".env"))

def get_required_env(name):
    value = os.getenv(name)

    if value is None or not value.strip():
        raise ValueError(f"Missing required environment variable: {name}")

    return value.strip()

def get_required_password():
    value = os.getenv("DB_PASSWORD")

    if value is None or value == "":
        raise ValueError("Missing requuired environment variable: DB_PASSWORD")

    return value

def get_database_port():
    value = os.getenv("DB_PORT", "5432").strip()

    try:
        port = int(value)
    except ValueError as exc:
        raise ValueError(f"DB_PORT must be a valid integer: {value!r}") from exc

    if not 1 <= 65535:
        raise ValueError(f"DB_PORT must be between 1 and 65535")

    return port

DATABASE_CONFIG = {
    "host": get_required_env("DB_HOST"),
    "database": get_required_env("DB_NAME"),
    "user": get_required_env("DB_USER"),
    "password": get_required_env("DB_PASSWORD"),
    "port": get_database_port(),
}
