# =====================================================
# Ubuntu SQL Server
# Network Management
# =====================================================

from database import execute_query

def get_network_interfaces(server_id=1):
    """
    Returns all network interfaces for a server
    """

    if not isinstance(server_id, int) or isinstance(server_id, bool):
        raise TypeError("server_id must be an integer")

    if server_id <= 0:
        raise ValueError("server_id must be greater than zero")

    query = """
        SELECT
            interface_name,
            ip_address,
            mac_address,
            network_type,
            speed_mbps
        FROM server_management.network_interfaces
        Where server_id = %s
        ORDER BY interface_name;
    """

    return execute_query(query, (server_id,))