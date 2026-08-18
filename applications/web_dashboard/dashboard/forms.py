from django import forms

class MaintenanceLogForm(forms.Form):
    action = forms.CharField(max_length=255)

    description = forms.CharField(
        required=False,
        widget=forms.Textarea
    )

    performed_by = forms.CharField(
        max_length=100,
        required=False
    )