from django.db import migrations


ROLES = {
    "Administrator": [
        "manage_users",
        "manage_ssh_access",
        "manage_servers",
    ],
    "Operator": [
        "manage_ssh_access",
        "manage_servers",
    ],
    "Viewer": [],
}


def create_rbac_roles(apps, schema_editor):
    Group = apps.get_model("auth", "Group")
    Permission = apps.get_model("auth", "Permission")

    for role_name, permission_codenames in ROLES.items():
        group, _ = Group.objects.get_or_create(name=role_name)

        permissions = Permission.objects.filter(
            content_type__app_label="dashboard",
            codename__in=permission_codenames,
        )

        group.permissions.set(permissions)


def remove_rbac_roles(apps, schema_editor):
    Group = apps.get_model("auth", "Group")

    Group.objects.filter(
        name__in=ROLES.keys(),
    ).delete()


class Migration(migrations.Migration):

    dependencies = [
        ("dashboard", "0001_rbac_permissions"),
    ]

    operations = [
        migrations.RunPython(
            create_rbac_roles,
            remove_rbac_roles,
        ),
    ]
