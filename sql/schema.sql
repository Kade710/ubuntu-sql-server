-- =====================================================
-- Ubuntu SQL Server
-- Schema Definition
-- =====================================================
--
-- Project:
-- Ubuntu SQL Server
--
-- Description:
-- Creates the database schema used by the project.
--
-- Author:
-- Jonathon Anderson
--
-- =====================================================

CREATE SCHEMA IF NOT EXISTS server_management;

COMMENT ON SCHEMA server_management IS
'Primary schema for server inventory, hardware, networking, operating system, and maintenance records.';