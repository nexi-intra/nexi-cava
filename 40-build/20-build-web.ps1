<#---
title: Build Web
tag: build-web
---#>

#

# Build Web


if ((Split-Path -Leaf (Split-Path  -Parent -Path $PSScriptRoot)) -eq "sessions"){
    $path = join-path $PSScriptRoot ".." ".."
}
else{
  $path = join-path $PSScriptRoot ".." ".koksmat/"

}

$koksmatDir = Resolve-Path $path
$env:NODE_ENV="development"
Set-Location (join-path $koksmatDir "web")

pnpm build