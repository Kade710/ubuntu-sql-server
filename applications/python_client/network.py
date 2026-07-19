# =====================================================
# Ubuntu SQL Server
# Network Management
# =====================================================

from database import execute_query

def get_network_interfaces(server_id=1):
    """
    Returns all network interfaces for a server
    """

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