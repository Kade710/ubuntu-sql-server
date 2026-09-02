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

class SSHKeyRegistrationForm(forms.Form):
    name = forms.CharField(
        max_length=100,
        label="Key Name",
    )

    public_key = forms.CharField(
        label="OpenSSH Public Key",
        widget=forms.Textarea(
            attrs={
                "rows": 5,
                "placeholder": "ssh-ed25519 AAAA... user@computer",
            }
        ),
    )