from django.shortcuts import HttpResponse

# The views

def hello(request):
    return HttpResponse("Hello, world!")