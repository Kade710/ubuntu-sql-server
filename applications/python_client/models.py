# =====================================================
# Ubuntu SQL Server
# Data Models
# =====================================================

from dataclasses import dataclass


@dataclass
class Server:
    id: int
    hostname: str
    ip_address: str
    operating_system: str
    cpu: str
    ram_gb: int
    storage_gb: int
    gpu: str
    motherboard: str