# Entity Relationship Diagram (ERD)

## Overview

The `ubuntu_sql_server` database uses a relational database structure centered around the `server_inventory` table.

The `server_inventory` table acts as the primary entity. Supporting information is stored in related tables using foreign key relationships.

Database Schema:


server_management


---

# Entity Relationship Model

                     server_inventory
                           |
                           |
    ------------------------------------------------
    |                  |              |             |
    |                  |              |             |
    v                  v              v             v

hardware_components operating_systems network_interfaces maintenance_logs


---

# Table Relationships

## server_inventory

Primary table containing server records.

Primary Key:


id


Referenced by:


hardware_components.server_id
operating_systems.server_id
network_interfaces.server_id
maintenance_logs.server_id


---

# hardware_components

Stores hardware information related to a server.

Relationship:


server_inventory.id
|
|
v
hardware_components.server_id


Relationship Type:


One server can have many hardware components


---

# operating_systems

Stores operating system details.

Relationship:


server_inventory.id
|
|
v
operating_systems.server_id


Relationship Type:


One server can have many operating system records


---

# network_interfaces

Stores network interface information.

Relationship:


server_inventory.id
|
|
v
network_interfaces.server_id


Relationship Type:


One server can have multiple network interfaces


---

# maintenance_logs

Stores administrative and maintenance history.

Relationship:


server_inventory.id
|
|
v
maintenance_logs.server_id


Relationship Type:


One server can have many maintenance records


---

# Foreign Key Structure

All supporting tables reference the primary server inventory table.


server_inventory
|
|
+-- hardware_components
|
+-- operating_systems
|
+-- network_interfaces
|
+-- maintenance_logs


Foreign key constraint pattern:

```sql
FOREIGN KEY (server_id)
REFERENCES server_management.server_inventory(id)

## Overview

The `ubuntu_sql_server` database uses a relational database structure centered around the `server_inventory` table.

The `server_inventory` table acts as the primary entity. Supporting information is stored in related tables using foreign key relationships.

Database Schema:


server_management


---

# Entity Relationship Model

                     server_inventory
                           |
                           |
    ------------------------------------------------
    |                  |              |             |
    |                  |              |             |
    v                  v              v             v

hardware_components operating_systems network_interfaces maintenance_logs


---

# Table Relationships

## server_inventory

Primary table containing server records.

Primary Key:


id


Referenced by:


hardware_components.server_id
operating_systems.server_id
network_interfaces.server_id
maintenance_logs.server_id


---

# hardware_components

Stores hardware information related to a server.

Relationship:


server_inventory.id
|
|
v
hardware_components.server_id


Relationship Type:


One server can have many hardware components


---

# operating_systems

Stores operating system details.

Relationship:


server_inventory.id
|
|
v
operating_systems.server_id


Relationship Type:


One server can have many operating system records


---

# network_interfaces

Stores network interface information.

Relationship:


server_inventory.id
|
|
v
network_interfaces.server_id


Relationship Type:


One server can have multiple network interfaces


---

# maintenance_logs

Stores administrative and maintenance history.

Relationship:


server_inventory.id
|
|
v
maintenance_logs.server_id


Relationship Type:


One server can have many maintenance records


---

# Foreign Key Structure

All supporting tables reference the primary server inventory table.


server_inventory
|
|
+-- hardware_components
|
+-- operating_systems
|
+-- network_interfaces
|
+-- maintenance_logs


Foreign key constraint pattern:

```sql
FOREIGN KEY (server_id)
REFERENCES server_management.server_inventory(id)
