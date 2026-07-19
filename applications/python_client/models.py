# =====================================================
# Ubuntu SQL Server
# Data Models
# =====================================================

class Server:

    def __init__(
        self,
        id,
        hostname,
        ip_address,
        operating_system,
        cpu,
        ram_gb,
        storage_gb,
        gpu,
        motherboard
    ):
        self.id = id
        self.hostname = hostname
        self.ip_address = ip_address
        self.operating_system = operating_system
        self.cpu = cpu
        self.ram_gb = ram_gb
        self.storage_gb = storage_gb
        self.gpu = gpu
        self.motherboard = motherboard