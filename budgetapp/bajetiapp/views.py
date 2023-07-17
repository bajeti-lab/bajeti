from django.shortcuts import render
from django.http import HttpResponse, HttpResponseRedirect
from .models import ExpenseInfo
from django.contrib.auth import logout,login,authenticate
from django.contrib.auth.models import User
from django.contrib.auth.forms import UserCreationForm
from django.db.models import Sum
import matplotlib.pyplot as plt
import numpy as np
from django.db.models import Q

# The views

def index(request):
    expense_items = ExpenseInfo.objects.filter(user_expense=request.user).order_by('-date_added')
    return render(request, 'bajetiapp/index.html',context={'user':request.user,'expense_items':expense_items})

def sign_up(request):
    if request.method == 'POST':
        form = UserCreationForm(request.POST)
        if form.is_valid():
            user=form.save()
            login(request,user)
            return HttpResponseRedirect('app')
        else:
            for msg in form.error_messages:
                print(form.error_messages[msg])
            return render(request, 'bajetiapp/sign_up.html',
                          {'form':form})
    else:
        form = UserCreationForm
        return render(request, 'bajetiapp/sign_up.html',
                      {'form':form})
def add_item(request):
    name = request.POST['expense_name']
    expense_cost = request.POST['cost']
    expense_date = request.POST['expense_date']

    ExpenseInfo.objects.create(expense_name=name,cost=expense_cost,date_added=expense_date,user_expense=request.user)
    return HttpResponseRedirect('app')

def logout_view(request):
    logout(request)
    return HttpResponseRedirect('/')

