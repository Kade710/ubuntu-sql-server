# =====================================================
# Ubuntu SQL Server
# Python Administrative Client 
# =====================================================

from inventory import get_servers


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

if __name__ == "__main__":
    main()
