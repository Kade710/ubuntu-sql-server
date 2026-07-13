-- =====================================================
-- Ubuntu SQL Server
-- Database Tables
-- =====================================================
--
-- Creates tables for:
-- Server inventory
-- Hardware components
-- Operating systems
-- Network interfaces
-- Maintenance records
--
-- Author:
-- Jonathon Anderson
--
-- =====================================================

CREATE TABLE IF NOT EXISTS server_management.server_inventory (
    id SERIAL PRIMARY KEY,
    hostname VARCHAR(100) NOT NULL,
    ip_address VARCHAR(45),
    operating_system VARCHAR(100),
    cpu VARCHAR(100),
    ram_gb INTEGER,
    storage_gb INTEGER,
    gpu VARCHAR(100),
    motherboard VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS server_management.hardware_components (
    id SERIAL PRIMARY KEY,
    server_id INTEGER NOT NULL,
    component_type VARCHAR(50) NOT NULL,
    manufacturer VARCHAR(100),
    model VARCHAR(100),
    specification TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT hardware_components_server_id_fkey
        FOREIGN KEY (server_id)
        REFERENCES server_management.server_inventory(id)
);


CREATE TABLE IF NOT EXISTS server_management.operating_systems (
    id SERIAL PRIMARY KEY,
    server_id INTEGER NOT NULL,
    distribution VARCHAR(100) NOT NULL,
    version VARCHAR(100),
    kernel VARCHAR(100),
    architecture VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT operating_systems_server_id_fkey
        FOREIGN KEY (server_id)
        REFERENCES server_management.server_inventory(id)
);


CREATE TABLE IF NOT EXISTS server_management.network_interfaces (
    id SERIAL PRIMARY KEY,
    server_id INTEGER NOT NULL,
    interface_name VARCHAR(50) NOT NULL,
    mac_address VARCHAR(17),
    ip_address VARCHAR(45),
    network_type VARCHAR(50),
    speed_mbps INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT network_interfaces_server_id_fkey
        FOREIGN KEY (server_id)
        REFERENCES server_management.server_inventory(id)
);


CREATE TABLE IF NOT EXISTS server_management.maintenance_logs (
    id SERIAL PRIMARY KEY,
    server_id INTEGER NOT NULL,
    action VARCHAR(100) NOT NULL,
    description TEXT,
    performed_by VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT maintenance_logs_server_id_fkey
        FOREIGN KEY (server_id)
        REFERENCES server_management.server_inventory(id)
);