# Web Dashboard

The Web Dashboard provides a browser-based view of U-Server information stored in the PostgreSQL database.

It is built with Django and provides an easier way to view server information without working directly with the database.

## Features

- View server inventory
- View hardware information
- View operating system information
- View network interfaces
- View system health
- View health history
- View maintenance logs
- View Go Agent status

## Database

The dashboard reads server management data from the `ubuntu_sql_server` PostgreSQL database.

Most project data is stored in the `server_management` schema.

## Technologies

- Python
- Django
- PostgreSQL
- HTML
- CSS

## Status

The Web Dashboard is working and actively being developed. More monitoring and management features will be added as the project grows.
