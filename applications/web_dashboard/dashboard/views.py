from datetime import timedelta

from django.shortcuts import get_object_or_404, render
from django.utils import timezone

from .models import HardwareComponent, HealthCheck, MaintenanceLog, NetworkInterface, Server

def index(request):
    servers = Server.objects.all().order_by("id")

    server_data = []

    for server in servers:
        latest_health = HealthCheck.objects.filter(
            server_id=server.id
        ).order_by("-created_at").first()

        agent_status = "UNKNOWN"

        if latest_health and latest_health.created_at:
            last_check = latest_health.created_at

            if timezone.is_naive(last_check):
                last_check = timezone.make_aware(
                    last_check,
                    timezone.get_current_timezone()
                )

            age = timezone.now() - last_check

            if age < timedelta(hours=25):
                agent_status = "ONLINE"
            elif age < timedelta(hours=48):
                agent_status = "STALE"
            else:
                agent_status = "OFFLINE"

        server_data.append(
            {
                "server": server,
                "latest_health": latest_health,
                "agent_status": agent_status,
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

    latest_health = HealthCheck.objects.filter(
        server_id=server_id
    ).order_by("-created_at").first()

    agent_status = "UNKNOWN"

    if latest_health and latest_health.created_at:
        last_check = latest_health.created_at

        if timezone.is_naive(last_check):
            last_check = timezone.make_aware(
                last_check,
                timezone.get_current_timezone()
            )

        age = timezone.now() - last_check

        if age < timedelta(hours=25):
            agent_status = "ONLINE"
        elif age < timedelta(hours=48):
            agent_status = "STALE"
        else:
            agent_status = "OFFLINE"

    health_chart = list(
        reversed(
            [
                {
                    "created_at": check.created_at.strftime("%m/%d %H:%M"),
                    "load_average": float(check.load_average or 0),
                    "memory_percent": float(check.memory_percent or 0),
                    "disk_percent": float(check.disk_percent or 0),
                }
                for check in health_checks
            ]
        )
    )

    maintenance_logs = MaintenanceLog.objects.filter(
        server_id=server_id
    ).order_by("-created_at")[:10]

    context = {
        "server": server,
        "hardware": hardware,
        "network": network,
        "latest_health": latest_health,
        "agent_status": agent_status,
        "health_chart": health_chart,
        "health_checks": health_checks,
        "maintenance_logs": maintenance_logs,
    }

    return render(request, "dashboard/server_detail.html", context)