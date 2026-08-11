from django.shortcuts import render

from .models import Server

def index(request):
    servers = Server.objects.all().order_by("id")

    context = {
        "servers": servers,
    }

    return render(request, "dashbaord/index.html", context)