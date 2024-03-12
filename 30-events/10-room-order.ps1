<#---
title: Room Order Changes
tag: room-order
api: post
trigger: sharepoint
list: Catering Orders
---

# 

# Handle room order changes

Regardless of this is the first time the record is seen or if it is an update, the same logic is used to handle the change.

On the time of the change the following steps are taken:

The requested delivery time is checked against the current time. If the requested delivery time is in the past, the order is rejected. 

Orders need to be received no later than at noon the **business** day before the delivery.


## Changing an dispatched ordered
An order that has been dispatched can not be changed, but might be cancelled.



#>

param (
    $orderId = "1",
    $outlookdate = "2024-03-24T08:00:00Z"
    )

throw "Not implemented yet"
Connect-PnPOnline -Url $ENV:SITEURL  -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH"

