"""
URL configuration for budgetapp project.

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/4.2/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.urls import path, include
from bajetiapp import views
from django.contrib.auth import views as auth_views


urlpatterns = [
    path('app', views.index, name='index'),
    path('', auth_views.LoginView.as_view(),name='login'),
    path('sign_up',views.sign_up,name="sign up"),
    path('add_item', views.add_item, name='add item'),
    path('accounts/', include('django.contrib.auth.urls')),
    path('logout', views.logout_view, name='logout')
]
