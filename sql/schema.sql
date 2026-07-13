-- =====================================================
-- Ubuntu SQL Server
-- Schema Definition
-- =====================================================
--
-- Project: Ubuntu SQL Server
-- Repository: github.com/Kade710/ubuntu-sql-server
--
-- Description:
-- Creates the primary schema used by the Ubuntu SQL
-- Server project.
--
-- Database:
-- ubuntu_sql_server
--
-- Author:
-- Jonathon Anderson
--
-- =====================================================

CREATE SCHEMA IF NOT EXISTS server_management;

COMMENT ON SCHEMA server_management IS
'Primary schema for server inventory, hardware, networking, operating system, and maintenance records.';
