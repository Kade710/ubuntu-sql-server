# =====================================================
# Ubuntu SQL Server
# PostgreSQL Database Connector
# =====================================================

import psycopg
from config import DATABASE_CONFIG


def get_connection():
    """
    Creates and returns a PostgreSQL database connection.
    """

    connection = psycopg.connect(
        host=DATABASE_CONFIG["host"],
        dbname=DATABASE_CONFIG["database"],
        user=DATABASE_CONFIG["user"],
        password=DATABASE_CONFIG["password"],
        port=DATABASE_CONFIG["port"]
    )

    return connection

def execute_query(query, params=None):
    """
    Executes a SELECT query and returns results.
    """

    connection = get_connection()

    try:
        with connection.cursor() as cursor:
            cursor.execute(query, params)

            return cursor.fetchall
        
    finally:
        connection.close()