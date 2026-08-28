# =====================================================
# Ubuntu SQL Server
# Python Administrative Client
# =====================================================

from hardware import get_hardware
from inventory import get_servers
from maintenance import get_maintenance_logs
from network import get_network_interfaces
from reports import generate_reports


def menu():
    """Displays the main menu."""

    print("\n" + "=" * 60)
    print("Ubuntu SQL Server Management Console")
    print("=" * 60)

    print("1. View Server Inventory")
    print("2. View Hardware Components")
    print("3. View Network Interfaces")
    print("4. View Maintenance Logs")
    print("5. Generate Report")
    print("0. Exit")

    return input("\nSelect an option: ").strip()


def display_inventory():
    """Displays all servers."""

    servers = get_servers()

    for server in servers:
        print("-" * 60)
        print(f"Hostname: {server[1]}")
        print(f"IP Address: {server[2]}")
        print(f"OS: {server[3]}")
        print(f"CPU: {server[4]}")
        print(f"RAM: {server[5]}")
        print(f"Storage: {server[6]}")
        print(f"GPU: {server[7]}")
        print(f"Motherboard: {server[8]}")
        print()


def display_hardware():
    """Displays hardware for every server."""

    servers = get_servers()

    for server in servers:
        print("-" * 60)
        print(f"Hardware Components for {server[1]}")
        print("-" * 60)

        hardware = get_hardware(server[0])

        for component in hardware:
            print(f"{component[0]}: {component[1]} {component[2]}")
            print(f"Specs: {component[3]}")
            print()


def display_network():
    """Displays network interfaces."""

    servers = get_servers()

    for server in servers:
        print("-" * 60)
        print(f"Network Interfaces for {server[1]}")
        print("-" * 60)

        interfaces = get_network_interfaces(server[0])

        for interface in interfaces:
            print(f"Interface: {interface[0]}")
            print(f"IP Address: {interface[1]}")
            print(f"MAC Address: {interface[2]}")
            print(f"Type: {interface[3]}")
            print(f"Speed: {interface[4]} Mbps")
            print()


def display_maintenance():
    """Displays maintenance history."""

    servers = get_servers()

    for server in servers:
        print("-" * 60)
        print(f"Maintenance Logs for {server[1]}")
        print("-" * 60)

        logs = get_maintenance_logs(server[0])

        for log in logs:
            print(f"Action: {log[0]}")
            print(f"Performed By: {log[1]}")
            print(f"Date: {log[2]}")
            print()


def main():
    """Runs the administrative client."""

    while True:
        try:
            choice = menu()

            if choice == "1":
                display_inventory()

            elif choice == "2":
                display_hardware()

            elif choice == "3":
                display_network()

            elif choice == "4":
                display_maintenance()

            elif choice == "5":
                generate_reports()

            elif choice == "0":
                print("\nSee ya later!\n")
                break

            else:
                print("\nInvalid option.\n")

        except (EOFError, KeyboardInterrupt):
            print("\n\nExiting.\n")
            break

        except (RuntimeError, OSError, ValueError) as exc:
            print(f"\nError: {exc}\n")


if __name__ == "__main__":
    main()