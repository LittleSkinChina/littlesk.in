Set-Location ./api
Invoke-WebRequest https://download.ip2location.com/lite/IP2LOCATION-LITE-DB1.IPV6.BIN.ZIP -OutFile ./db.zip
unzip db.zip
