-- =====================================================
-- Ubuntu SQL Server
-- Example SQL Queries
-- =====================================================
--
-- Project:
-- Ubuntu SQL Server
--
-- Description:
-- Common queries used to view and manage the
-- Ubuntu SQL Server inventory database.
--
-- Author:
-- Jonathon Anderson
--
-- =====================================================


-- =====================================================
-- View All Servers
-- =====================================================

SELECT *
FROM server_management.server_inventory;


-- =====================================================
-- View Hardware Components
-- =====================================================

SELECT *
FROM server_management.hardware_components
ORDER BY component_type;


-- =====================================================
-- View Operating Systems
-- =====================================================

SELECT *
FROM server_management.operating_systems;


-- =====================================================
-- View Network Interfaces
-- =====================================================

SELECT *
FROM server_management.network_interfaces;


-- =====================================================
-- View Maintenance History
-- =====================================================

SELECT *
FROM server_management.maintenance_logs
ORDER BY created_at DESC;


-- =====================================================
-- Server Hardware Report
-- =====================================================

SELECT
    s.hostname,
    h.component_type,
    h.manufacturer,
    h.model,
    h.specification
FROM server_management.server_inventory s
JOIN server_management.hardware_components h
ON s.id = h.server_id
ORDER BY h.component_type;


-- =====================================================
-- Operating System Report
-- =====================================================

SELECT
    s.hostname,
    o.distribution,
    o.version,
    o.kernel,
    o.architecture
FROM server_management.server_inventory s
JOIN server_management.operating_systems o
ON s.id = o.server_id;


-- =====================================================
-- Network Report
-- =====================================================

SELECT
    s.hostname,
    n.interface_name,
    n.ip_address,
    n.mac_address,
    n.network_type,
    n.speed_mbps
FROM server_management.server_inventory s
JOIN server_management.network_interfaces n
ON s.id = n.server_id;


-- =====================================================
-- Maintenance Report
-- =====================================================

SELECT
    s.hostname,
    m.action,
    m.performed_by,
    m.created_at
FROM server_management.server_inventory s
JOIN server_management.maintenance_logs m
ON s.id = m.server_id
ORDER BY m.created_at DESC;


-- =====================================================
-- Hardware Count
-- =====================================================

SELECT
    component_type,
    COUNT(*) AS total_components
FROM server_management.hardware_components
GROUP BY component_type
ORDER BY component_type;


-- =====================================================
-- Server Summary
-- =====================================================

SELECT
    hostname,
    operating_system,
    cpu,
    ram_gb,
    storage_gb,
    gpu
FROM server_management.server_inventory;
