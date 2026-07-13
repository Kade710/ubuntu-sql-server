-- =====================================================
-- Ubuntu SQL Server
-- Database Deployment Script
-- =====================================================
--
-- Executes all SQL setup files in the correct order.
--
-- Order:
-- 1. Schema creation
-- 2. Table creation
-- 3. Sample data insertion
--
-- Author:
-- Jonathon Anderson
--
-- =====================================================

\echo 'Creating database schema...'
\i sql/schema.sql

\echo 'Creating database tables...'
\i sql/tables.sql

\echo 'Loading sample data...'
\i sql/sample_data.sql

\echo 'Database deployment completed.'
