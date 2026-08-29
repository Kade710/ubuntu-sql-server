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


def create_server(
    hostname,
    ip_address=None,
    operating_system=None,
    cpu=None,
    ram_gb=None,
    storage_gb=None,
    gpu=None,
    motherboard=None,
):
    """
    Creates a new server inventory record.
    """

    query = """
        INSERT INTO server_management.server_inventory (
            hostname,
            ip_address,
            operating_system,
            cpu,
            ram_gb,
            storage_gb,
            gpu,
            motherboard
        )
        VALUES (%s, %s, %s, %s, %s, %s, %s, %s)
        RETURNING
            id,
            hostname,
            ip_address,
            operating_system,
            cpu,
            ram_gb,
            storage_gb,
            gpu,
            motherboard,
            created_at;
    """

    params = (
        hostname,
        ip_address,
        operating_system,
        cpu,
        ram_gb,
        storage_gb,
        gpu,
        motherboard,
    )

    return execute_query(query, params)


def update_server(
    server_id,
    hostname,
    ip_address=None,
    operating_system=None,
    cpu=None,
    ram_gb=None,
    storage_gb=None,
    gpu=None,
    motherboard=None,
):
    """
    Updates an existing server inventory record.
    """

    query = """
        UPDATE server_management.server_inventory
        SET
            hostname = %s,
            ip_address = %s,
            operating_system = %s,
            cpu = %s,
            ram_gb = %s,
            storage_gb = %s,
            gpu = %s,
            motherboard = %s
        WHERE id = %s
        RETURNING
            id,
            hostname,
            ip_address,
            operating_system,
            cpu,
            ram_gb,
            storage_gb,
            gpu,
            motherboard,
            created_at;
    """

    params = (
        hostname,
        ip_address,
        operating_system,
        cpu,
        ram_gb,
        storage_gb,
        gpu,
        motherboard,
        server_id,
    )

    return execute_query(query, params)