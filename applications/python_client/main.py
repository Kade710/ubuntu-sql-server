# =====================================================
# Ubuntu SQL Server
# Python Administrative Client 
# =====================================================

from inventory import get_servers
from hardware import get_hardware

def main():

    print("Ubuntu SQL Server Management Console")
    print(" ")

    servers = get_servers()

    for server in servers:
        print(server)
        print(f"Hostname: {server[1]}")
        print(f"IP Adress: {server[2]}")
        print(f"OS: {server[3]}")
        print(f"CPU: {server[4]}")
        print(f"RAM: {server[5]}")
        print(f"Storage: {server[6]}")
        print(f"GPU: {server[7]}")
        print(f"Motherboard: {server[8]}")


        print()
        print("Hardware Components")
        print("  ")

        hardware = get_hardware(server[0])

        for component in hardware:
            print(f"{component[0]}: {component[1]} {component[2]}")
            print(f"Specs: {component[3]}")
            print()

if __name__ == "__main__":
    main()
