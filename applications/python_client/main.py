# =====================================================
# Ubuntu SQL Server
# Python Client Test
# =====================================================

from database import get_connection


def main():

    connection = get_connection()

    print("Connected to PostgreSQL successfully!")

    connection.close()


if __name__ == "__main__":
    main()
