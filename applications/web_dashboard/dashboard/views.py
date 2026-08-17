from django.shortcuts import get_object_or_404, render

from .models import HardwareComponent, HealthCheck, MaintenanceLog, NetworkInterface, Server

def index(request):
    servers = Server.objects.all().order_by("id")

    server_data = []

    for server in servers:
        latest_health = HealthCheck.objects.filter(
            server_id=server.id
        ).order_by("-created_at").first()

        server_data.append(
            {
                "server": server,
                "latest_health": latest_health,
            }
        )

    context = {
        "server_data": server_data,
    }

    return render(request, "dashboard/index.html", context)

def server_detail(request, server_id):
    server = get_object_or_404(Server, id=server_id)

    hardware = HardwareComponent.objects.filter(
        server_id=server_id
    ).order_by("id")

    network = NetworkInterface.objects.filter(
        server_id=server_id
    ).order_by("id")

    health_checks = HealthCheck.objects.filter(
        server_id=server_id
    ).order_by("-created_at")[:10]

    maintenance_logs = MaintenanceLog.objects.filter(
        server_id=server_id
    ).order_by("-created_at")[:10]

    context = {
        "server": server,
        "hardware": hardware,
        "network": network,
        "health_checks": health_checks,
        "maintenance_logs": maintenance_logs,
    }

    return render(request, "dashboard/server_detail.html", context)