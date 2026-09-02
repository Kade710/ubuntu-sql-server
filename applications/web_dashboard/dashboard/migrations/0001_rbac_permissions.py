from django.db import migrations

PERMISSIONS = [
    ("manage_users", "Can manage users"),
    ("manage_ssh_access", "Can manage SSH access"),
    ("manage_servers", "Can manage servers"),
]

def create_rbac_permissions(apps, schema_editor):
    ContentType = apps.get_model("contenttypes", "ContentType")
    Permission = apps.get_model("auth", "Permission")

    content_type, _ = ContentType.objects.get_or_create(
        app_label="dashboard",
        model="accesscontrol",
    )

    for codename, name in PERMISSIONS:
        Permission.objects.get_or_create(
            content_type=content_type,
            codename=codename,
            defaults={"name": name},
        )


def remove_rbac_permissions(apps, schema_editor):
    ContentType = apps.get_model("contenttypes", "ContentType")
    Permission = apps.get_model("auth", "Permission")

    content_type = ContentType.objects.filter(
        app_label="dashboard",
        model="accesscontrol",
    ).first()

    if content_type is None:
        return

    Permission.objects.filter(
        content_type=content_type,
        codename__in=[codename for codename, _ in PERMISSIONS],
    ).delete()

    if not Permission.objects.filter(content_type=content_type).exists():
        content_type.delete()


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ("contenttypes", "0002_remove_content_type_name"),
        ("auth", "0012_alter_user_first_name_max_length"),
    ]

    operations = [
        migrations.RunPython(
            create_rbac_permissions,
            remove_rbac_permissions,
        ),
    ]
