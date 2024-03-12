<#---
title: Tomorrow Orders

---

#>

param (
    $baseDate = "2024-03-11"
)
Connect-PnPOnline -Url $ENV:SITEURL  -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH"

$listItems = Get-PnpListItem -List "Catering Orders"  # | Where-Object { $_.FieldValues.Provisioning_x0020_Status -eq "Provision" }



$today  = [System.DateTime]::Parse($baseDate)

write-host "Base date: $baseDate" $today
write-host "Orders in list: $($listItems.Count)"

foreach ($item in $listItems) {
    #$item.FieldValues.Title

    if ("No title" -eq $item.FieldValues.Title) {
        continue
    }

    
    # write-host $item.FieldValues.Title $item.FieldValues.Appointmentstart $item.FieldValues.Status
   

    $date = [System.DateTime]::Parse($item.FieldValues.Appointmentstart)
    if ($date.Subtract($today).Days -le 0){
        continue
    
    }

   write-host "Order in " $date.Subtract($today).Days "days"
    

    write-host "Day/Month/Year" $date.Day $date.Month $date.Year 



}

