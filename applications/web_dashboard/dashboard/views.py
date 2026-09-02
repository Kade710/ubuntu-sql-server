from datetime import timedelta

from django.contrib.auth import get_user_model
from django.contrib.auth.models import Group
from django.http import JsonResponse
from django.shortcuts import get_object_or_404, redirect, render
from django.utils import timezone
from django.views.decorators.http import require_GET, require_POST

from .forms import MaintenanceLogForm
from .models import (
    HardwareComponent,
    HealthCheck,
    MaintenanceLog,
    NetworkInterface,
    Server,
)


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
                    timezone.get_current_timezone(),
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
                timezone.get_current_timezone(),
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

    if request.method == "POST":
        maintenance_form = MaintenanceLogForm(request.POST)

        if maintenance_form.is_valid():
            MaintenanceLog.objects.create(
                server_id=server_id,
                action=maintenance_form.cleaned_data["action"],
                description=maintenance_form.cleaned_data["description"],
                performed_by=maintenance_form.cleaned_data["performed_by"],
                created_at=timezone.now(),
            )

            return redirect(
                "server_detail",
                server_id=server_id,
            )

    else:
        maintenance_form = MaintenanceLogForm()

    context = {
        "server": server,
        "hardware": hardware,
        "network": network,
        "latest_health": latest_health,
        "agent_status": agent_status,
        "health_chart": health_chart,
        "health_checks": health_checks,
        "maintenance_form": maintenance_form,
        "maintenance_logs": maintenance_logs,
    }

    return render(
        request,
        "dashboard/server_detail.html",
        context,
    )


@require_GET
def api_user_list(request):
    if not request.user.has_perm("dashboard.manage_users"):
        return JsonResponse(
            {"detail": "Permission denied."},
            status=403,
        )

    User = get_user_model()

    users = User.objects.all().order_by("username")

    data = [
        {
            "id": user.id,
            "username": user.username,
            "email": user.email,
            "is_active": user.is_active,
            "is_staff": user.is_staff,
            "is_superuser": user.is_superuser,
        }
        for user in users
    ]

    return JsonResponse({"users": data})


@require_POST
def api_user_create(request):
    if not request.user.has_perm("dashboard.manage_users"):
        return JsonResponse(
            {"detail": "Permission denied."},
            status=403,
        )

    username = request.POST.get("username", "").strip()
    email = request.POST.get("email", "").strip()
    password = request.POST.get("password", "")

    if not username or not password:
        return JsonResponse(
            {"detail": "Username and password are required."},
            status=400,
        )

    User = get_user_model()

    if User.objects.filter(username=username).exists():
        return JsonResponse(
            {"detail": "Username already exists."},
            status=400,
        )

    user = User.objects.create_user(
        username=username,
        email=email,
        password=password,
    )

    return JsonResponse(
        {
            "user": {
                "id": user.id,
                "username": user.username,
                "email": user.email,
                "is_active": user.is_active,
                "is_staff": user.is_staff,
                "is_superuser": user.is_superuser,
            }
        },
        status=201,
    )


@require_POST
def api_user_disable(request, user_id):
    if not request.user.has_perm("dashboard.manage_users"):
        return JsonResponse(
            {"detail": "Permission denied."},
            status=403,
        )

    User = get_user_model()

    user = get_object_or_404(
        User,
        id=user_id,
    )

    user.is_active = False
    user.save(update_fields=["is_active"])

    return JsonResponse(
        {
            "user": {
                "id": user.id,
                "username": user.username,
                "is_active": user.is_active,
            }
        },
        status=200,
    )


@require_POST
def api_user_role(request, user_id):
    if not request.user.has_perm("dashboard.manage_users"):
        return JsonResponse(
            {"detail": "Permission denied."},
            status=403,
        )

    role_name = request.POST.get("role", "").strip()

    allowed_roles = {
        "Administrator",
        "Operator",
        "Viewer",
    }

    if role_name not in allowed_roles:
        return JsonResponse(
            {"detail": "Invalid role."},
            status=400,
        )

    User = get_user_model()

    user = get_object_or_404(
        User,
        id=user_id,
    )

    role = get_object_or_404(
        Group,
        name=role_name,
    )

    user.groups.add(role)

    return JsonResponse(
        {
            "user": {
                "id": user.id,
                "username": user.username,
                "roles": list(
                    user.groups.order_by("name").values_list(
                        "name",
                        flat=True,
                    )
                ),
            }
        },
        status=200,
    )

    def test_assigning_role_replaces_existing_role(self):
        self.client.force_login(self.admin)

        self.client.post(
            reverse("api_user_role", args=[self.user.id]),
            data={"role": "Viewer"},
        )

        response = self.client.post(
            reverse("api_user_role", args=[self.user.id]),
            data={"role": "Operator"},
        )

        self.assertEqual(response.status_code, 200)

        self.user.refresh_from_db()

        self.assertTrue(
            self.user.groups.filter(name="Operator").exists()
        )

        self.assertFalse(
            self.user.groups.filter(name="Viewer").exists()
        )
