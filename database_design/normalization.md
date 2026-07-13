# Database Normalization

## Overview

The `ubuntu_sql_server` database follows relational database normalization principles to reduce duplicate data, improve data integrity, and maintain clear relationships between records.

The database design primarily follows Third Normal Form (3NF).

---

# First Normal Form (1NF)

A table satisfies First Normal Form when:

- Each column contains a single value
- Each record is uniquely identifiable
- There are no repeating groups of data

The database follows 1NF by storing each server as an individual record.

Example:


server_inventory


Each server has a unique identifier:


id


Information such as:

- hostname
- IP address
- CPU
- RAM
- storage

is stored as individual fields.

---

# Second Normal Form (2NF)

A table satisfies Second Normal Form when:

- It is already in 1NF
- All non-key attributes depend on the primary key

The database separates related information into dedicated tables.

Example:

Instead of storing multiple hardware components directly inside:


server_inventory


hardware details are stored separately:


hardware_components


Relationship:


server_inventory.id
|
v
hardware_components.server_id


This prevents repeating hardware information.

---

# Third Normal Form (3NF)

A table satisfies Third Normal Form when:

- It is already in 2NF
- Non-key fields do not depend on other non-key fields

The database separates different categories of information.

Examples:

Operating system information is stored in:


operating_systems


Network information is stored in:


network_interfaces


Maintenance records are stored in:


maintenance_logs


This prevents unrelated information from being stored together.

---

# Normalized Database Structure


server_inventory

Primary Entity
|
|
+-- hardware_components

Hardware Information
|
|
+-- operating_systems

Operating System Information
|
|
+-- network_interfaces

Network Information
|
|
+-- maintenance_logs

Maintenance History


---

# Benefits of Normalization

The normalized design provides:

- Reduced duplicate information
- Easier database maintenance
- Improved data accuracy
- Better scalability
- Cleaner API integration
- Easier reporting

---

# Future Expansion

Future tables can be added without redesigning the existing structure.

Potential additions:

- users
- authentication
- applications
- monitoring_metrics
- alerts
- performance_history

The current structure provides a foundation for expanding the Ubuntu SQL Server management platform.
