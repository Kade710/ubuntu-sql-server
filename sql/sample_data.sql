-- =====================================================
-- Ubuntu SQL Server
-- Sample Data
-- =====================================================
--
-- Initial server inventory data
--
-- Author:
-- Jonathon Anderson
--
-- =====================================================


-- =====================================================
-- Server Inventory
-- =====================================================

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
VALUES (
    'U-Server',
    '192.168.1.100',
    'Ubuntu 24.04.4 LTS',
    'Intel Core i5-4690K 3.50GHz',
    16,
    2000,
    'NVIDIA GeForce GTX 750',
    'MSI Z97 GAMING 5 (MS-7917)'
);


-- =====================================================
-- Hardware Components
-- =====================================================

INSERT INTO server_management.hardware_components (
    server_id,
    component_type,
    manufacturer,
    model,
    specification
)
VALUES
(
    1,
    'CPU',
    'Intel',
    'Core i5-4690K',
    '4 Cores, 4 Threads, 3.50 GHz Base, 4.00 GHz Turbo'
),
(
    1,
    'RAM',
    'G.Skill',
    'F3-12800CL10-8GBXL',
    '16 GB DDR3 1600 MT/s (2 x 8 GB)'
),
(
    1,
    'Motherboard',
    'MSI',
    'Z97 GAMING 5 (MS-7917)',
    'Intel Z97 Chipset'
),
(
    1,
    'GPU',
    'NVIDIA',
    'GeForce GTX 750',
    'GM107'
);


-- =====================================================
-- Operating System
-- =====================================================

INSERT INTO server_management.operating_systems (
    server_id,
    distribution,
    version,
    kernel,
    architecture
)
VALUES
(
    1,
    'Ubuntu',
    '24.04.4 LTS',
    'Linux 6.17.0-35-generic',
    'x86_64'
);


-- =====================================================
-- Network Interface
-- =====================================================

INSERT INTO server_management.network_interfaces (
    server_id,
    interface_name,
    mac_address,
    ip_address,
    network_type,
    speed_mbps
)
VALUES
(
    1,
    'enp3s0',
    '02:00:00:00:00:01',
    '192.168.1.100',
    'Ethernet',
    1000
);


-- =====================================================
-- Maintenance History
-- =====================================================

INSERT INTO server_management.maintenance_logs (
    server_id,
    action,
    description,
    performed_by
)
VALUES
(
    1,
    'PostgreSQL Installation',
    'Installed PostgreSQL 16.14 on Ubuntu 24.04.4 LTS and verified database service operation.',
    'jonathon_admin'
),
(
    1,
    'Database Configuration',
    'Created ubuntu_sql_server database, server_management schema, and initial inventory tables.',
    'jonathon_admin'
),
(
    1,
    'SSH Configuration',
    'Configured SSH access for remote administration from Windows PowerShell and Visual Studio Code.',
    'jonathon_admin'
),
(
    1,
    'System Update',
    'Updated Ubuntu packages and verified system services after maintenance.',
    'jonathon_admin'
),
(
    1,
    'System Maintenance Automation',
    'Created system_update.sh script to automate Ubuntu package updates, reboot checks, uptime verification, and PostgreSQL service monitoring.',
    'jonathon_admin'
),
(
    1,
    'Dependency Installation Script',
    'Created install_dependencies.sh to install and verify required software packages for the Ubuntu SQL Server project.',
    'jonathon_admin'
);
