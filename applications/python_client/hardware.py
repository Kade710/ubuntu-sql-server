# =====================================================
# Ubuntu SQL Server
# Hardware Management
# =====================================================

from database import execute_query


def get_hardware(server_id=1):
    """
    Returns hardware components for a server.
    """

    if not isinstance(server_id, int) or isinstance(server_id, bool):
        raise TypeError("server_id must be an integer")

    if server_id <= 0:
        raise ValueError("server_id must be greater than zero")

    query = """
        SELECT
            component_type,
            manufacturer,
            model,
            specification
        FROM server_management.hardware_components
        WHERE server_id = %s
        ORDER BY component_type;
    """

    return execute_query(query, (server_id,))