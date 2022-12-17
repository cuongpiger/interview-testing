from django.conf.urls import url
from django.contrib import admin
from django.urls import path, include

from rest_framework_simplejwt.views import (
    TokenObtainPairView,
    TokenRefreshView,
)

api = 'api/v1/'

urlpatterns = [
    path('admin/', admin.site.urls),
    path(api, include('customers.urls')),
    path(api + 'api-auth', include('rest_framework.urls')),
    path(api + 'auth/token/', TokenObtainPairView.as_view()),
    path(api + 'auth/token/refresh/', TokenRefreshView.as_view()),
    path(api + 'auth/', include('djoser.urls')),
    path(api + 'auth/', include('djoser.urls.authtoken')),
    path(api + 'auth/', include('djoser.urls.jwt')),
]

