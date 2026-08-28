# =====================================================
# Ubuntu SQL Server
# Report Generation
# =====================================================

from datetime import datetime
from pathlib import Path

from hardware import get_hardware
from inventory import get_servers
from maintenance import get_maintenance_logs
from network import get_network_interfaces


def generate_reports(output_dir="."):
    """
    Generates a server report and returns the created report path.
    """

    output_path = Path(output_dir).expanduser()

    if output_path.exists() and not output_path.is_dir():
        raise NotADirectoryError(
            f"Report output path is not a directory: {output_path}"
        )

    output_path.mkdir(parents=True, exist_ok=True)

    timestamp = datetime.now().strftime("%Y%m%d_%H%M%S_%f")

    filename = output_path / f"server_report_{timestamp}.txt"
    temp_filename = filename.with_suffix(".tmp")

    servers = get_servers()

    try:
        with temp_filename.open("w", encoding="utf-8") as report:
            report.write("=" * 60 + "\n")
            report.write("Ubuntu SQL Server Report\n")
            report.write("=" * 60 + "\n\n")

            for server in servers:
                server_id = server[0]

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

                for component in get_hardware(server_id):
                    report.write(
                        f"{component[0]}: {component[1]} {component[2]}\n"
                    )
                    report.write(f"Specs: {component[3]}\n\n")

                report.write("NETWORK\n")
                report.write("-" * 60 + "\n")

                for interface in get_network_interfaces(server_id):
                    report.write(f"Interface: {interface[0]}\n")
                    report.write(f"IP Address: {interface[1]}\n")
                    report.write(f"MAC Address: {interface[2]}\n")
                    report.write(f"Type: {interface[3]}\n")
                    report.write(f"Speed: {interface[4]} Mbps\n\n")

                report.write("MAINTENANCE LOGS\n")
                report.write("-" * 60 + "\n")

                for log in get_maintenance_logs(server_id):
                    report.write(f"Action: {log[0]}\n")
                    report.write(f"Performed By: {log[1]}\n")
                    report.write(f"Date: {log[2]}\n\n")

        temp_filename.replace(filename)

    except Exception:
        if temp_filename.exists():
            temp_filename.unlink()

        raise

    print(f"Report created: {filename}")

    return filename