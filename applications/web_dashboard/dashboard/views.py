from django.shortcuts import get_object_or_404, render

from .models import Server

def index(request):
    servers = Server.objects.all().order_by("id")

    context = {
        "servers": servers,
    }

    return render(request, "dashboard/index.html", context)

def server_detail(request, server_id):
    server = get_object_or_404(Server, id=server_id)

    context = {
        "server": server,
    }

    return render(request, "dashboard/server_detail.html", context)