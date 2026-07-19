# =====================================================
# Ubuntu SQL Server
# Hardware Management
# =====================================================

from database import execute_query

def get_hardware(server_id=1):
    """
    Returns hardware components for a server.
    """

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