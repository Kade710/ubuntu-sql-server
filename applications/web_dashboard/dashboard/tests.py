from django.contrib.auth import get_user_model
from django.contrib.auth.models import Permission
from django.test import TestCase
from django.urls import reverse


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