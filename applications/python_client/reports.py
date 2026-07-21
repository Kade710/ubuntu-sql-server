# =====================================================
# Ubuntu SQL Server
# Report Generation
# =====================================================

from datetime import datetime

from inventory import get_servers
from hardware import get_hardware
from network import get_network_interfaces
from maintenance import get_maintenance_logs


def generate_reports():

    servers = get_servers()

    filename + f"server_report_{datetime.now().strftime('%Y%m%d_%H%M%S')}.txt"

    with open(filename, "w") as report:

        reports.write("=" * 60 "\n")
        reports.write("Ubuntu SQL Server Report\n")
        reports.write("=" * 60 "\n\n")

        for server in servers:

            report.write("SERVER INFORMATION\n")
            report.write("-" * 60 + "\n")

            report.write(f"Hostname: {server[1]}\n")
            report.write(f"IP Address: {server[2]}\n")
            report.write(f"OS: {server[3]}\n")
            report.write(f"CPU: {server[4]}\n")
            report.write(f"RAM: {server[5]} GB\n")
            report.write(f"Storage: {server[6]} GB\n")
            report.write(f"GPU: {server[7]}\n")
            report.write(f"Motherboard: {server[8]}\n\n")


            report.write("HARDWARE\n")
            report.write("-" * 60 + "\n")

            for component in get_hardware(server[0]):
                report.write(
                    f"{component[0]}: {component[1]} {component[2]}\n"
                )
                report.write(f"Specs: {component[3]}\n\n")


            report.write("NETWORK\n")
            report.write("-" * 60 + "\n")

            for interface in get_network_interfaces(server[0]):
                report.write(f"Interface: {interface[0]}\n")
                report.write(f"IP Address: {interface[1]}\n")
                report.write(f"MAC Address: {interface[2]}\n")
                report.write(f"Type: {interface[3]}\n")
                report.write(f"Speed: {interface[4]} Mbps\n\n")


            report.write("MAINTENANCE LOGS\n")
            report.write("-" * 60 + "\n")

            for log in get_maintenance_logs(server[0]):
                report.write(f"Action: {log[0]}\n")
                report.write(f"Performed By: {log[1]}\n")
                report.write(f"Date: {log[2]}\n\n")


    print(f"Report created: {filename}")