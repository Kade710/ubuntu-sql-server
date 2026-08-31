from django.contrib.auth import get_user_model
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
        response = self.client.get(reverse("api_user_list"))

        self.assertNotEqual(response.status_code, 200)

    def test_admin_can_list_users(self):
        self.client.force_login(self.admin)

        response = self.client.get(reverse("api_user_list"))

        self.assertEqual(response.status_code, 200)

        data = response.json()

        self.assertEqual(len(data["users"]), 2)

        usernames = [user["username"] for user in data["users"]]

        self.assertIn("testadmin", usernames)
        self.assertIn("testuser", usernames)