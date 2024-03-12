<#---
title: Web deploy to production
tag: webdeployproduction
api: post
---
We start by finding which version tag to use
#>

$appname = "nexi-cava"
$imagename = "nexi-cava"
$dnsname = "cava.home.nexi-intra.com"
$inputFile = join-path  $env:KITCHENROOT $appname ".koksmat","koksmat.json"
$port = "4325"
if (!(Test-Path -Path $inputFile) ) {
   Throw "Cannot find file at expected path: $inputFile"
} 
$json = Get-Content -Path $inputFile | ConvertFrom-Json
$version = "v$($json.version.major).$($json.version.minor).$($json.version.patch).$($json.version.build)"

<#
The we build the deployment file
#>

$image = "ghcr.io/nexi-intra/$($imagename)-web:$($version)"

$config = @"
apiVersion: apps/v1
kind: Deployment
metadata:
  name: $appname
spec:
  selector:
    matchLabels:
      app: $appname
  replicas: 1
  template:
    metadata:
      labels:
        app: $appname
    spec: 
      containers:
      - name: $appname
        image: $image
        ports:
          - containerPort: $port
        env:
        - name: DEBUG
          value: *

        
---
apiVersion: v1
kind: Service
metadata:
  name: $appname
  labels:
    app: $appname
    service: $appname
spec:
  ports:
  - name: http
    port: 5301
    targetPort: $port
  selector:
    app: $appname
---    
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: $appname
spec:
  rules:
  - host: $dnsname
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: $appname
            port:
              number: 5301
    

"@

write-host "Applying config" -ForegroundColor Green

write-host $config -ForegroundColor Gray

$config |  kubectl apply -f -