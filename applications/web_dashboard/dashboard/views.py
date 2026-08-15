from django.shortcuts import get_object_or_404, render

from .models import HardwareComponent, NetworkInterface, Server

def index(request):
    servers = Server.objects.all().order_by("id")

    context = {
        "servers": servers,
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

    context = {
        "server": server,
        "hardware": hardware,
        "network": network,
    }

    return render(request, "dashboard/server_detail.html", context)