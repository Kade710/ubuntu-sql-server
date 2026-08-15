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
        managed = Falsedb_table = 'server_manager"."network_interfaces'

    def __str__(self):
        return self.interface_name