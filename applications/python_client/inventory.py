# =====================================================
# Ubuntu SQL Server
# Inventory Management
# =====================================================

from database import execute_query


def get_servers():
    """
    Returns all servers in inventory.
    """

    query = """
        SELECT
            id,
            hostname,
            ip_address,
            operating_system,
            cpu,
            ram_gb,
            storage_gb,
            gpu,
            motherboard
        FROM server_management.server_inventory
        ORDER BY hostname, id;
    """

    return execute_query(query)