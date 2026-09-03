import subprocess
from unittest.mock import patch

from django.contrib.auth import get_user_model
from django.contrib.auth.models import Group, Permission
from django.test import TestCase
from django.urls import reverse

from .models import AccessAuditLog, SSHKey


class UserListAPITests(TestCase):
    def setUp(self):
        User = get_user_model()

        self.admin = User.objects.create_superuser(
            username="testadmin",
            email="admin@example.com",
            password="TestPassword123!",
        )

        self.user = User.objects.create_user(
            username="testuser",
            email="user@example.com",
            password="TestPassword123!",
        )

    def test_unauthenticated_request_is_denied(self):
        response = self.client.get(
            reverse("api_user_list")
        )

        self.assertNotEqual(
            response.status_code,
            200,
        )

    def test_admin_can_list_users(self):
        self.client.force_login(self.admin)

        response = self.client.get(
            reverse("api_user_list")
        )

        self.assertEqual(
            response.status_code,
            200,
        )

        data = response.json()

        self.assertEqual(
            len(data["users"]),
            2,
        )

        usernames = [
            user["username"]
            for user in data["users"]
        ]

        self.assertIn(
            "testadmin",
            usernames,
        )

        self.assertIn(
            "testuser",
            usernames,
        )

    def test_user_with_manage_users_permission_can_list_users(self):
        User = get_user_model()

        permitted_user = User.objects.create_user(
            username="permitteduser",
            email="permitted@example.com",
            password="TestPassword123!",
        )

        permission = Permission.objects.get(
            codename="manage_users",
            content_type__app_label="dashboard",
        )

        permitted_user.user_permissions.add(
            permission
        )

        self.client.force_login(
            permitted_user
        )

        response = self.client.get(
            reverse("api_user_list")
        )

        self.assertEqual(
            response.status_code,
            200,
        )

    def test_user_without_manage_users_permission_is_denied(self):
        User = get_user_model()

        staff_user = User.objects.create_user(
            username="staffuser",
            email="staff@example.com",
            password="TestPassword123!",
            is_staff=True,
        )

        self.client.force_login(
            staff_user
        )

        response = self.client.get(
            reverse("api_user_list")
        )

        self.assertEqual(
            response.status_code,
            403,
        )


class UserCreateAPITests(TestCase):
    def setUp(self):
        User = get_user_model()

        self.admin = User.objects.create_superuser(
            username="testadmin",
            email="admin@example.com",
            password="TestPassword123!",
        )

    def test_unauthenticated_request_is_denied(self):
        response = self.client.post(
            reverse("api_user_create"),
            data={
                "username": "newuser",
                "email": "newuser@example.com",
                "password": "SecurePassword123!",
            },
        )

        self.assertNotEqual(
            response.status_code,
            201,
        )

    def test_admin_can_create_user(self):
        self.client.force_login(
            self.admin
        )

        response = self.client.post(
            reverse("api_user_create"),
            data={
                "username": "newuser",
                "email": "newuser@example.com",
                "password": "SecurePassword123!",
            },
        )

        self.assertEqual(
            response.status_code,
            201,
        )

        User = get_user_model()

        user = User.objects.get(
            username="newuser"
        )

        self.assertEqual(
            user.email,
            "newuser@example.com",
        )

        self.assertTrue(
            user.check_password(
                "SecurePassword123!"
            )
        )

        self.assertTrue(
            user.is_active
        )


class UserDisableAPITests(TestCase):
    def setUp(self):
        User = get_user_model()

        self.admin = User.objects.create_superuser(
            username="testadmin",
            email="admin@example.com",
            password="TestPassword123!",
        )

        self.user = User.objects.create_user(
            username="testuser",
            email="user@example.com",
            password="TestPassword123!",
        )

    def test_unauthenticated_request_is_denied(self):
        response = self.client.post(
            reverse(
                "api_user_disable",
                args=[self.user.id],
            )
        )

        self.assertNotEqual(
            response.status_code,
            200,
        )

    def test_admin_can_disable_user(self):
        self.client.force_login(
            self.admin
        )

        response = self.client.post(
            reverse(
                "api_user_disable",
                args=[self.user.id],
            )
        )

        self.assertEqual(
            response.status_code,
            200,
        )

        self.user.refresh_from_db()

        self.assertFalse(
            self.user.is_active
        )


class UserRoleAPITests(TestCase):
    def setUp(self):
        User = get_user_model()

        self.admin = User.objects.create_superuser(
            username="testadmin",
            email="admin@example.com",
            password="TestPassword123!",
        )

        self.user = User.objects.create_user(
            username="testuser",
            email="user@example.com",
            password="TestPassword123!",
        )

    def test_unauthenticated_request_is_denied(self):
        response = self.client.post(
            reverse(
                "api_user_role",
                args=[self.user.id],
            ),
            data={
                "role": "Operator",
            },
        )

        self.assertNotEqual(
            response.status_code,
            200,
        )

    def test_admin_can_assign_role(self):
        self.client.force_login(
            self.admin
        )

        response = self.client.post(
            reverse(
                "api_user_role",
                args=[self.user.id],
            ),
            data={
                "role": "Operator",
            },
        )

        self.assertEqual(
            response.status_code,
            200,
        )

        self.user.refresh_from_db()

        self.assertTrue(
            self.user.groups.filter(
                name="Operator"
            ).exists()
        )

    def test_invalid_role_is_rejected(self):
        self.client.force_login(
            self.admin
        )

        response = self.client.post(
            reverse(
                "api_user_role",
                args=[self.user.id],
            ),
            data={
                "role": "InvalidRole",
            },
        )

        self.assertEqual(
            response.status_code,
            400,
        )

        self.assertFalse(
            self.user.groups.filter(
                name="InvalidRole"
            ).exists()
        )

    def test_assigning_role_replaces_existing_role(self):
        self.client.force_login(
            self.admin
        )

        self.client.post(
            reverse(
                "api_user_role",
                args=[self.user.id],
            ),
            data={
                "role": "Viewer",
            },
        )

        response = self.client.post(
            reverse(
                "api_user_role",
                args=[self.user.id],
            ),
            data={
                "role": "Operator",
            },
        )

        self.assertEqual(
            response.status_code,
            200,
        )

        self.user.refresh_from_db()

        self.assertTrue(
            self.user.groups.filter(
                name="Operator"
            ).exists()
        )

        self.assertFalse(
            self.user.groups.filter(
                name="Viewer"
            ).exists()
        )

class RBACRolePermissionTests(TestCase):
    def test_administrator_has_expected_permissions(self):
        User = get_user_model()

        user = User.objects.create_user(
            username="administrator",
            password="TestPassword123!",
        )

        user.groups.add(
            user.groups.model.objects.get(
                name="Administrator"
            )
        )

        self.assertTrue(
            user.has_perm("dashboard.manage_users")
        )

        self.assertTrue(
            user.has_perm("dashboard.manage_ssh_access")
        )

        self.assertTrue(
            user.has_perm("dashboard.manage_servers")
        )

    def test_operator_has_expected_permissions(self):
        User = get_user_model()

        user = User.objects.create_user(
            username="operator",
            password="TestPassword123!",
        )

        user.groups.add(
            user.groups.model.objects.get(
                name="Operator"
            )
        )

        self.assertFalse(
            user.has_perm("dashboard.manage_users")
        )

        self.assertTrue(
            user.has_perm("dashboard.manage_ssh_access")
        )

        self.assertTrue(
            user.has_perm("dashboard.manage_servers")
        )

    def test_viewer_has_no_management_permissions(self):
        User = get_user_model()

        user = User.objects.create_user(
            username="viewer",
            password="TestPassword123!",
        )

        user.groups.add(
            user.groups.model.objects.get(
                name="Viewer"
            )
        )

        self.assertFalse(
            user.has_perm("dashboard.manage_users")
        )

        self.assertFalse(
            user.has_perm("dashboard.manage_ssh_access")
        )

        self.assertFalse(
            user.has_perm("dashboard.manage_servers")
        )

class UserManagementPageTests(TestCase):
    def setUp(self):
        User = get_user_model()

        self.admin = User.objects.create_superuser(
            username="testadmin",
            email="admin@example.com",
            password="TestPassword123!",
        )

    def test_admin_can_view_user_management_page(self):
        self.client.force_login(self.admin)

        response = self.client.get(
            reverse("user_management")
        )

        self.assertEqual(
            response.status_code,
            200,
        )

        self.assertTemplateUsed(
            response,
            "dashboard/users.html",
        )

    def test_user_without_manage_users_permission_is_denied(self):
        User = get_user_model()

        user = User.objects.create_user(
            username="viewer",
            email="viewer@example.com",
            password="TestPassword123!",
        )

        self.client.force_login(user)

        response = self.client.get(
            reverse("user_management")
        )

        self.assertEqual(
            response.status_code,
            403,
        )

class AccessAuditVerificationTests(TestCase):
    def setUp(self):
        User = get_user_model()

        self.admin = User.objects.create_superuser(
            username="auditadmin",
            email="auditadmin@example.com",
            password="AdminTestPassword123!",
        )

        self.client.force_login(self.admin)

    def test_user_onboarding_and_revocation_are_audited(self):
        User = get_user_model()

        create_response = self.client.post(
            reverse("api_user_create"),
            {
                "username": "lifecycleuser",
                "email": "lifecycle@example.com",
                "password": "LifecyclePassword123!",
            },
        )

        self.assertEqual(
            create_response.status_code,
            201,
        )

        user = User.objects.get(
            username="lifecycleuser"
        )

        role_response = self.client.post(
            reverse(
                "api_user_role",
                args=[user.id],
            ),
            {
                "role": "Operator",
            },
        )

        self.assertEqual(
            role_response.status_code,
            200,
        )

        disable_response = self.client.post(
            reverse(
                "api_user_disable",
                args=[user.id],
            )
        )

        self.assertEqual(
            disable_response.status_code,
            200,
        )

        user.refresh_from_db()

        self.assertFalse(user.is_active)

        self.assertEqual(
            list(
                user.groups.values_list(
                    "name",
                    flat=True,
                )
            ),
            ["Operator"],
        )

        logs = AccessAuditLog.objects.filter(
            target_identifier="lifecycleuser"
        ).order_by("created_at", "id")

        self.assertEqual(
            list(
                logs.values_list(
                    "action",
                    flat=True,
                )
            ),
            [
                "USER_CREATED",
                "USER_ROLE_CHANGED",
                "USER_DISABLED",
            ],
        )

        for log in logs:
            self.assertEqual(
                log.actor,
                self.admin,
            )

            self.assertEqual(
                log.target_type,
                "user",
            )

    def test_password_is_not_stored_in_audit_log(self):
        password = "DoNotAuditThisPassword123!"

        response = self.client.post(
            reverse("api_user_create"),
            {
                "username": "secretcheck",
                "email": "secretcheck@example.com",
                "password": password,
            },
        )

        self.assertEqual(
            response.status_code,
            201,
        )

        log = AccessAuditLog.objects.get(
            action="USER_CREATED",
            target_identifier="secretcheck",
        )

        self.assertNotIn(
            password,
            log.details,
        )

        self.assertNotIn(
            password,
            log.target_identifier,
        )

    def test_browser_user_actions_are_audited(self):
        User = get_user_model()

        response = self.client.post(
            reverse("user_management_create"),
            {
                "username": "browseraudit",
                "email": "browseraudit@example.com",
                "password": "BrowserPassword123!",
            },
        )

        self.assertEqual(
            response.status_code,
            302,
        )

        user = User.objects.get(
            username="browseraudit"
        )

        response = self.client.post(
            reverse(
                "user_management_role",
                args=[user.id],
            ),
            {
                "role": "Viewer",
            },
        )

        self.assertEqual(
            response.status_code,
            302,
        )

        response = self.client.post(
            reverse(
                "user_management_disable",
                args=[user.id],
            )
        )

        self.assertEqual(
            response.status_code,
            302,
        )

        logs = AccessAuditLog.objects.filter(
            target_identifier="browseraudit"
        ).order_by("created_at", "id")

        self.assertEqual(
            list(
                logs.values_list(
                    "action",
                    flat=True,
                )
            ),
            [
                "USER_CREATED",
                "USER_ROLE_CHANGED",
                "USER_DISABLED",
            ],
        )


class SSHAccessAuditTests(TestCase):
    def setUp(self):
        User = get_user_model()

        self.operator = User.objects.create_user(
            username="sshoperator",
            email="sshoperator@example.com",
            password="OperatorPassword123!",
        )

        operator_group = Group.objects.get(
            name="Operator"
        )

        self.operator.groups.add(
            operator_group
        )

        self.client.force_login(
            self.operator
        )

        self.public_key = (
            "ssh-ed25519 "
            "AAAAC3NzaC1lZDI1NTE5AAAAIGZha2V0ZXN0a2V5ZGF0YQ== "
            "test@example.com"
        )

    @patch(
        "dashboard.views._sync_uaccess_ssh_keys"
    )
    def test_ssh_registration_is_audited(
        self,
        mock_sync,
    ):
        response = self.client.post(
            reverse("ssh_key_register"),
            {
                "name": "Audit Test Key",
                "public_key": self.public_key,
            },
        )

        self.assertEqual(
            response.status_code,
            302,
        )

        mock_sync.assert_called_once()

        ssh_key = SSHKey.objects.get(
            name="Audit Test Key"
        )

        self.assertTrue(
            ssh_key.is_active
        )

        log = AccessAuditLog.objects.get(
            action="SSH_KEY_REGISTERED",
            target_identifier=ssh_key.fingerprint,
        )

        self.assertEqual(
            log.actor,
            self.operator,
        )

        self.assertEqual(
            log.target_type,
            "ssh_key",
        )

        self.assertNotIn(
            self.public_key,
            log.details,
        )

    @patch(
        "dashboard.views._sync_uaccess_ssh_keys"
    )
    def test_ssh_revocation_is_audited(
        self,
        mock_sync,
    ):
        from dashboard.views import _ssh_public_key_fingerprint

        fingerprint = _ssh_public_key_fingerprint(
            self.public_key
        )

        ssh_key = SSHKey.objects.create(
            user=self.operator,
            name="Revoke Test Key",
            public_key=self.public_key,
            fingerprint=fingerprint,
        )

        response = self.client.post(
            reverse(
                "ssh_key_revoke",
                args=[ssh_key.id],
            )
        )

        self.assertEqual(
            response.status_code,
            302,
        )

        mock_sync.assert_called_once()

        ssh_key.refresh_from_db()

        self.assertFalse(
            ssh_key.is_active
        )

        self.assertIsNotNone(
            ssh_key.revoked_at
        )

        log = AccessAuditLog.objects.get(
            action="SSH_KEY_REVOKED",
            target_identifier=ssh_key.fingerprint,
        )

        self.assertEqual(
            log.actor,
            self.operator,
        )

        self.assertEqual(
            log.target_type,
            "ssh_key",
        )

    @patch(
        "dashboard.views._sync_uaccess_ssh_keys"
    )
    def test_failed_ssh_registration_rolls_back_without_audit(
        self,
        mock_sync,
    ):
        mock_sync.side_effect = subprocess.SubprocessError(
            "simulated synchronization failure"
        )

        response = self.client.post(
            reverse("ssh_key_register"),
            {
                "name": "Failed Registration",
                "public_key": self.public_key,
            },
        )

        self.assertEqual(
            response.status_code,
            302,
        )

        self.assertFalse(
            SSHKey.objects.filter(
                name="Failed Registration"
            ).exists()
        )

        self.assertFalse(
            AccessAuditLog.objects.filter(
                action="SSH_KEY_REGISTERED"
            ).exists()
        )

    @patch(
        "dashboard.views._sync_uaccess_ssh_keys"
    )
    def test_failed_ssh_revocation_rolls_back_without_audit(
        self,
        mock_sync,
    ):
        from dashboard.views import _ssh_public_key_fingerprint

        fingerprint = _ssh_public_key_fingerprint(
            self.public_key
        )

        ssh_key = SSHKey.objects.create(
            user=self.operator,
            name="Failed Revocation",
            public_key=self.public_key,
            fingerprint=fingerprint,
        )

        mock_sync.side_effect = subprocess.SubprocessError(
            "simulated synchronization failure"
        )

        response = self.client.post(
            reverse(
                "ssh_key_revoke",
                args=[ssh_key.id],
            )
        )

        self.assertEqual(
            response.status_code,
            302,
        )

        ssh_key.refresh_from_db()

        self.assertTrue(
            ssh_key.is_active
        )

        self.assertIsNone(
            ssh_key.revoked_at
        )

        self.assertFalse(
            AccessAuditLog.objects.filter(
                action="SSH_KEY_REVOKED",
                target_identifier=ssh_key.fingerprint,
            ).exists()
        )

    @patch(
        "dashboard.views._sync_uaccess_ssh_keys"
    )
    def test_user_without_permission_cannot_register_ssh_key(
        self,
        mock_sync,
    ):
        User = get_user_model()

        viewer = User.objects.create_user(
            username="sshviewer",
            password="ViewerPassword123!",
        )

        viewer_group = Group.objects.get(
            name="Viewer"
        )

        viewer.groups.add(
            viewer_group
        )

        self.client.force_login(
            viewer
        )

        response = self.client.post(
            reverse("ssh_key_register"),
            {
                "name": "Unauthorized Key",
                "public_key": self.public_key,
            },
        )

        self.assertEqual(
            response.status_code,
            403,
        )

        mock_sync.assert_not_called()

        self.assertFalse(
            SSHKey.objects.filter(
                name="Unauthorized Key"
            ).exists()
        )

        self.assertFalse(
            AccessAuditLog.objects.filter(
                action="SSH_KEY_REGISTERED"
            ).exists()
        )
