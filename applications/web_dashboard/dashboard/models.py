from django.db import models

class Server(models.Model):
    id = models.IntegerField(primary_key=True)
    hostname = models.CharField(max_length=100)
    ip_address = models.CharField(max_length=45, null=True, blank=True)
    operating_system = models.CharField(max_length=100, null=True, blank=True)
    cpu = models.CharField(max_length=100, null=True, blank=True)
    ram_gb = models.IntegerField(null=True, blank=True)
    storage_gb = models.IntegerField(null=True, blank=True)
    gpu = models.CharField(max_length=100, null=True, blank=True)
    motherboard = models.CharField(max_length=100, null=True, blank=True)

    class Meta:
        managed = False
        db_table = 'server_management"."server_inventory'

    def __str__(self):
        return self.hostname

class HardwareComponent(models.Model):
    id = models.IntegerField(primary_key=True)
    server_id = models.IntegerField()
    component_type = models.CharField(max_length=50)
    manufacturer = models.CharField(max_length=100, null=True, blank=True)
    model = models.CharField(max_length=100, null=True, blank=True)
    specification = models.TextField(null=True, blank=True)

    class Meta:
        managed = False
        db_table = 'server_management"."hardware_components'

    def __str__(self):
        return f"{self.component_type} - {self.model}"

class NetworkInterface(models.Model):
    id = models.IntegerField(primary_key=True)
    server_id = models.IntegerField()
    interface_name = models.CharField(max_length=100)
    mac_address = models.CharField(max_length=100, null=True, blank=True)
    ip_address = models.CharField(max_length=100, null=True, blank=True)
    network_type = models.CharField(max_length=50, null=True, blank=True)
    speed_mbps = models.IntegerField(null=True, blank=True)

    class Meta:
        managed = False
        db_table = 'server_management"."network_interfaces'

    def __str__(self):
        return self.interface_name

class HealthCheck(models.Model):
    id = models.IntegerField(primary_key=True)
    server_id = models.IntegerField()
    load_average = models.DecimalField(max_digits=6, decimal_places=2, null=True, blank=True)
    memory_percent = models.DecimalField(max_digits=5, decimal_places=2, null=True, blank=True)
    disk_percent = models.DecimalField(max_digits=5, decimal_places=2, null=True, blank=True)
    uptime_hours = models.DecimalField(max_digits=12, decimal_places=2, null=True, blank=True)
    overall_status = models.CharField(max_length=20, null=True, blank=True)
    created_at = models.DateTimeField(null=True, blank=True)

    class Meta:
        managed =False
        db_table = 'server_management"."health_checks'

class MaintenanceLog(models.Model):
    id = models.AutoField(primary_key=True)
    server_id = models.IntegerField()
    action = models.CharField(max_length=100)
    description = models.TextField(null=True, blank=True)
    performed_by = models.CharField(max_length=100, null=True, blank=True)
    created_at = models.DateTimeField(null=True, blank=True)

    class Meta:
        managed = False
        db_table = 'server_management"."maintenance_logs'

class SSHKey(models.Model):
    user = models.ForeignKey(
        "auth.User",
        on_delete=models.CASCADE,
        related_name="ssh_keys",
    )
    name = models.CharField(max_length=100)
    public_key = models.TextField()
    fingerprint = models.CharField(
        max_length=100,
        unique=True
    )
    is_active = models.BooleanField(default=True)
    created_at = models.DateTimeField(auto_now_add=True)
    revoked_at = models.DateTimeField(
        null=True,
        blank=True,
    )

    class Meta:
        ordering = ["user__username", "name"]

    def __str__(self):
        return f"{self.user.username} - {self.name}"
