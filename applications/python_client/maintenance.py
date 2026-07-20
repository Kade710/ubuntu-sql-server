# =====================================================
# Ubuntu SQL Server
# Maintenance Management
# =====================================================

from database import execute_query

def get_maintenance_logs(server_id=1):
    """
    Returns maintenance history for a server.
    """

    query = """
        SELECT
            action,
            performed_by,
            created_at
        FROM server_management.maintenance_logs
        WHERE server_id = %s
        ORDER BY created_at DESC;
    """

    return execute_query(query, (server_id,))