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