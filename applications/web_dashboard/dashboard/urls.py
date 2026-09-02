from django.urls import path

from . import views


urlpatterns = [
    path(
        "",
        views.index,
        name="index",
    ),
    path(
        "server/<int:server_id>/",
        views.server_detail,
        name="server_detail",
    ),

    # User management dashboard
    path(
        "users/",
        views.user_management,
        name="user_management",
    ),
    path(
        "users/create/",
        views.user_management_create,
        name="user_management_create",
    ),
    path(
        "users/<int:user_id>/disable/",
        views.user_management_disable,
        name="user_management_disable",
    ),
    path(
        "users/<int:user_id>/role/",
        views.user_management_role,
        name="user_management_role",
    ),

    # SSH access management
    path(
        "ssh/",
        views.ssh_access_management,
        name="ssh_access_management",
    ),
    path(
        "ssh/register/",
        views.ssh_key_register,
        name="ssh_key_register",
    ),
    path(
        "ssh/<int:key_id>/revoke/",
        views.ssh_key_revoke,
        name="ssh_key_revoke",
    ),


    # User management API
    path(
        "api/users/",
        views.api_user_list,
        name="api_user_list",
    ),
    path(
        "api/users/create/",
        views.api_user_create,
        name="api_user_create",
    ),
    path(
        "api/users/<int:user_id>/disable/",
        views.api_user_disable,
        name="api_user_disable",
    ),
    path(
        "api/users/<int:user_id>/role/",
        views.api_user_role,
        name="api_user_role",
    ),
]
