# =====================================================
# Ubuntu SQL Server
# Python Client Test
# =====================================================

from database import get_connection


def main():

    connection = get_connection()

    print("Connected to PostgreSQL successfully!")

    connection.close()

    servers = execute_query("""
        SELECT
            hostname,
            ip_addess,
            operating_system,
            cpu,
            ram_gb,
            storage_gb,
            gpu
        FROM server_managment.server_inventory;
    """)

    print("\nServer Inventory:")

    for server in servers:
        print(server)


if __name__ == "__main__":
    main()
