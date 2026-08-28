# Python Client

The Python Client is part of the Ubuntu SQL Server project. It provides a command-line administrative interface for viewing server management data stored in PostgreSQL and generating server reports.

## Features

The client currently supports:

- Viewing server inventory
- Viewing hardware components
- Viewing network interfaces
- Viewing maintenance history
- Generating text-based server reports
- PostgreSQL database connectivity
- Environment-based database configuration

## Technologies

- Python 3
- PostgreSQL
- psycopg
- python-dotenv

## Project Structure

```text
python_client/
├── config.py
├── database.py
├── hardware.py
├── inventory.py
├── main.py
├── maintenance.py
├── models.py
├── network.py
├── reports.py
├── README.md
└── requirements.txt