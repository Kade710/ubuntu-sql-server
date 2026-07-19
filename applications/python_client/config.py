# =====================================================
# Ubuntu SQL Server
# Python Client Configuration
# =====================================================

import os
from dotenv import load_dotenv

load_dotenv()

DATABASE_CONFIG = {
    "host": os.getenv("DB_HOST"),
    "database": os.getenv("DB_NAME"),
    "user": os.getenv("DB_USER"),
    "password": os.getenv("DB_PASSORD"),
    "port": int(os.getenv("DB_PORT", 5432))
}
