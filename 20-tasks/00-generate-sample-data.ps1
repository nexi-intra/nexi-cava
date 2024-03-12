<#---
title: Generate Sample Data
tag: generate-sample-data
api: post
---#>

# 

# Generate Sample Data

param (
    $upn = "niels.johansen@nexigroup.com",
    $firstdate = "2024-03-24T08:00:00Z",
    $numberoforders = 3
    )


Connect-PnPOnline -Url $ENV:SITEURL  -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH"
$date = [System.DateTime]::Parse($firstdate)
 
$orderData = @"
89;10;35;540;Breakfast 1
90;20;43;540;Breakfast 2
91;30;60.5;540;Breakfast 3
99;120;12;540;Juice
108;120;10;540;Coffee / Tea 
"@

$values = @{
    "Title"          = "Test Title"
    "OrderData"      = $orderData;
    Appointmentstart = "2024-03-24T08:00:00Z"
}


for ($i=1; $i -lt $numberoforders; $i++) {
    $values["Title"] = $upn + ":test" + (New-Guid).Guid
    $orderDate = $date.AddDays(1 * $i)

    $values["Appointmentstart"] = $orderDate.ToString("yyyy-MM-ddTHH:mm:ssZ")
    Add-PnPListItem -List "Catering Orders" -Values $values
}
Add-PnPListItem -List "Catering Orders" -Values $values

