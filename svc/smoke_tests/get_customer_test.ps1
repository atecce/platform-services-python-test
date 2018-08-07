param (
    [string]$email = "testing@yahoo.com"
)

$uri = "http://localhost:7050/customers/" + $email

Invoke-WebRequest -Uri $uri