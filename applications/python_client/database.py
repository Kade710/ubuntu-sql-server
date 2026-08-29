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

    try:
        return psycopg.connect(
            host=DATABASE_CONFIG["host"],
            dbname=DATABASE_CONFIG["database"],
            user=DATABASE_CONFIG["user"],
            password=DATABASE_CONFIG["password"],
             port=DATABASE_CONFIG["port"]
        )

    except psycopg.Error as exc:
        raise RuntimeError(
            f"Failed to connect to PostgreSQL: {exc}"
        ) from exc


def execute_query(query, params=None):
    """
    Executes a SELECT query and returns results.
    """

    if not isinstance(query, str) or not query.strip():
        raise ValueError("Query must be non-empty string")

    connection = get_connection()

    try:
        with connection.cursor() as cursor:
            cursor.execute(query, params)
            return cursor.fetchall()

    except psycopg.Error as exc:
        raise RuntimeError(
            f"PostgreSQL query failed: {exc}"
        ) from exc
        
    finally:
        connection.close()