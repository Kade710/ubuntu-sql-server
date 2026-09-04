import base64
import binascii
import hashlib
import os
import subprocess
import tempfile

from datetime import timedelta

from django.contrib import messages
from django.contrib.auth import get_user_model
from django.contrib.auth.decorators import login_required
from django.contrib.auth.models import Group
from django.http import JsonResponse
from django.shortcuts import get_object_or_404, redirect, render
from django.utils import timezone
from django.views.decorators.http import require_GET, require_POST

from .forms import MaintenanceLogForm, SSHKeyRegistrationForm
from .models import (
    AccessAuditLog,
    HardwareComponent,
    HealthCheck,
    MaintenanceLog,
    NetworkInterface,
    Server,
    SSHKey,
)


@login_required
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


@login_required
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


def _audit_access_event(
    request,
    action,
    target_type,
    target_identifier,
    details="",
):
    AccessAuditLog.objects.create(
        actor=request.user if request.user.is_authenticated else None,
        action=action,
        target_type=target_type,
        target_identifier=str(target_identifier),
        details=details,
    )


@login_required
def user_management(request):
    if not request.user.has_perm("dashboard.manage_users"):
        return JsonResponse(
            {"detail": "Permission denied."},
            status=403,
        )

    User = get_user_model()

    users = User.objects.all().order_by("username")

    context = {
        "users": users,
        "roles": [
            "Administrator",
            "Operator",
            "Viewer",
        ],
    }

    return render(
        request,
        "dashboard/users.html",
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

    _audit_access_event(
        request=request,
        action="USER_CREATED",
        target_type="user",
        target_identifier=user.username,
        details=f"User {user.username} created through the API.",
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

    _audit_access_event(
        request=request,
        action="USER_DISABLED",
        target_type="user",
        target_identifier=user.username,
        details=f"User {user.username} disabled through the API.",
    )

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

    existing_managed_roles = Group.objects.filter(
        name__in=allowed_roles
    )

    user.groups.remove(
        *existing_managed_roles
    )

    user.groups.add(role)

    _audit_access_event(
        request=request,
        action="USER_ROLE_CHANGED",
        target_type="user",
        target_identifier=user.username,
        details=f"User {user.username} assigned to {role_name} through the API.",
    )

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


@require_POST
def user_management_create(request):
    if not request.user.has_perm("dashboard.manage_users"):
        return JsonResponse(
            {"detail": "Permission denied."},
            status=403,
        )

    username = request.POST.get("username", "").strip()
    email = request.POST.get("email", "").strip()
    password = request.POST.get("password", "")

    if not username or not password:
        messages.error(
            request,
            "Username and password are required.",
        )
        return redirect("user_management")

    User = get_user_model()

    if User.objects.filter(username=username).exists():
        messages.error(
            request,
            "Username already exists.",
        )
        return redirect("user_management")

    user = User.objects.create_user(
        username=username,
        email=email,
        password=password,
    )

    _audit_access_event(
        request=request,
        action="USER_CREATED",
        target_type="user",
        target_identifier=user.username,
        details=f"User {user.username} created through the dashboard.",
    )

    messages.success(
        request,
        f"User {user.username} created successfully.",
    )

    return redirect("user_management")


@require_POST
def user_management_disable(request, user_id):
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
    user.save(
        update_fields=["is_active"]
    )

    _audit_access_event(
        request=request,
        action="USER_DISABLED",
        target_type="user",
        target_identifier=user.username,
        details=f"User {user.username} disabled through the dashboard.",
    )

    messages.success(
        request,
        f"User {user.username} disabled.",
    )

    return redirect("user_management")


@require_POST
def user_management_role(request, user_id):
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
        messages.error(
            request,
            "Invalid role.",
        )
        return redirect("user_management")

    User = get_user_model()

    user = get_object_or_404(
        User,
        id=user_id,
    )

    role = get_object_or_404(
        Group,
        name=role_name,
    )

    managed_roles = Group.objects.filter(
        name__in=allowed_roles
    )

    user.groups.remove(
        *managed_roles
    )

    user.groups.add(role)

    _audit_access_event(
        request=request,
        action="USER_ROLE_CHANGED",
        target_type="user",
        target_identifier=user.username,
        details=(
            f"User {user.username} assigned to "
            f"{role_name} through the dashboard."
        ),
    )

    messages.success(
        request,
        f"{user.username} assigned to {role_name}.",
    )

    return redirect("user_management")


def _sync_uaccess_ssh_keys():
    staging_path = "/var/lib/ubuntu-sql-server/ssh/uaccess.keys"
    staging_dir = os.path.dirname(staging_path)

    active_keys = SSHKey.objects.filter(
        is_active=True
    ).order_by("id")

    key_data = "".join(
        f"{ssh_key.public_key.strip()}\n"
        for ssh_key in active_keys
    )

    temp_path = None

    try:
        with tempfile.NamedTemporaryFile(
            mode="w",
            encoding="utf-8",
            dir=staging_dir,
            prefix=".uaccess.",
            delete=False,
        ) as temp_file:
            temp_file.write(key_data)
            temp_file.flush()
            os.fsync(temp_file.fileno())
            temp_path = temp_file.name

        os.chmod(temp_path, 0o600)
        os.replace(temp_path, staging_path)

        subprocess.run(
            [
                "sudo",
                "-n",
                "/usr/local/sbin/u-server-sync-uaccess-keys",
            ],
            check=True,
            capture_output=True,
            text=True,
            timeout=10,
        )

    finally:
        if temp_path and os.path.exists(temp_path):
            os.unlink(temp_path)


def _ssh_public_key_fingerprint(public_key):
    parts = public_key.strip().split()

    if len(parts) < 2:
        raise ValueError("Invalid OpenSSH public key.")

    key_type = parts[0]
    key_data = parts[1]

    allowed_types = {
        "ssh-ed25519",
        "ssh-rsa",
        "ecdsa-sha2-nistp256",
        "ecdsa-sha2-nistp384",
        "ecdsa-sha2-nistp521",
    }

    if key_type not in allowed_types:
        raise ValueError("Unsupported SSH public key type.")

    try:
        decoded_key = base64.b64decode(
            key_data,
            validate=True,
        )
    except (binascii.Error, ValueError):
        raise ValueError("Invalid SSH public key encoding.")

    if not decoded_key:
        raise ValueError("Invalid SSH public key.")

    digest = hashlib.sha256(decoded_key).digest()

    fingerprint = base64.b64encode(
        digest
    ).decode("ascii").rstrip("=")

    return f"SHA256:{fingerprint}"


def ssh_access_management(request):
    if not request.user.has_perm("dashboard.manage_ssh_access"):
        return JsonResponse(
            {"detail": "Permission denied."},
            status=403,
        )

    keys = SSHKey.objects.select_related(
        "user"
    ).order_by(
        "user__username",
        "name",
    )

    form = SSHKeyRegistrationForm()

    context = {
        "ssh_keys": keys,
        "form": form,
    }

    return render(
        request,
        "dashboard/ssh_access.html",
        context,
    )


@require_POST
def ssh_key_register(request):
    if not request.user.has_perm("dashboard.manage_ssh_access"):
        return JsonResponse(
            {"detail": "Permission denied."},
            status=403,
        )

    form = SSHKeyRegistrationForm(request.POST)

    if not form.is_valid():
        messages.error(
            request,
            "A key name and valid public key are required.",
        )
        return redirect("ssh_access_management")

    name = form.cleaned_data["name"].strip()
    public_key = form.cleaned_data["public_key"].strip()

    try:
        fingerprint = _ssh_public_key_fingerprint(public_key)
    except ValueError as exc:
        messages.error(
            request,
            str(exc),
        )
        return redirect("ssh_access_management")

    if SSHKey.objects.filter(
        fingerprint=fingerprint
    ).exists():
        messages.error(
            request,
            "This SSH public key is already registered.",
        )
        return redirect("ssh_access_management")

    ssh_key = SSHKey.objects.create(
        user=request.user,
        name=name,
        public_key=public_key,
        fingerprint=fingerprint,
    )

    try:
        _sync_uaccess_ssh_keys()
    except (OSError, subprocess.SubprocessError):
        ssh_key.delete()

        messages.error(
            request,
            "SSH key synchronization failed. The key was not registered.",
        )
        return redirect("ssh_access_management")

    _audit_access_event(
        request=request,
        action="SSH_KEY_REGISTERED",
        target_type="ssh_key",
        target_identifier=ssh_key.fingerprint,
        details=(
            f'Key "{ssh_key.name}" registered '
            f"for {ssh_key.user.username}."
        ),
    )

    messages.success(
        request,
        f'SSH key "{ssh_key.name}" registered.',
    )

    return redirect("ssh_access_management")


@require_POST
def ssh_key_revoke(request, key_id):
    if not request.user.has_perm("dashboard.manage_ssh_access"):
        return JsonResponse(
            {"detail": "Permission denied."},
            status=403,
        )

    ssh_key = get_object_or_404(
        SSHKey,
        id=key_id,
    )

    if not ssh_key.is_active:
        messages.error(
            request,
            "This SSH key has already been revoked.",
        )
        return redirect("ssh_access_management")

    previous_revoked_at = ssh_key.revoked_at

    ssh_key.is_active = False
    ssh_key.revoked_at = timezone.now()

    ssh_key.save(
        update_fields=[
            "is_active",
            "revoked_at",
        ]
    )

    try:
        _sync_uaccess_ssh_keys()
    except (OSError, subprocess.SubprocessError):
        ssh_key.is_active = True
        ssh_key.revoked_at = previous_revoked_at

        ssh_key.save(
            update_fields=[
                "is_active",
                "revoked_at",
            ]
        )

        messages.error(
            request,
            "SSH key synchronization failed. The key remains active.",
        )
        return redirect("ssh_access_management")

    _audit_access_event(
        request=request,
        action="SSH_KEY_REVOKED",
        target_type="ssh_key",
        target_identifier=ssh_key.fingerprint,
        details=(
            f'Key "{ssh_key.name}" revoked '
            f"for {ssh_key.user.username}."
        ),
    )

    messages.success(
        request,
        f'SSH key "{ssh_key.name}" revoked.',
    )

    return redirect("ssh_access_management")