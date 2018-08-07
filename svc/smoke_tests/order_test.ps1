param (
    [string]$email = "testing@yahoo.com",
    [double]$total = 20
)

$body = '{"email":"' + $email + '","total":' + $total + '}'

Invoke-WebRequest -Uri http://localhost:7050/order -Method Patch -Body $body -ContentType "application/json"