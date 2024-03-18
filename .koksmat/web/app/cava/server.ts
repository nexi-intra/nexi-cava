"use server"

import { runProcess } from "@/koksmat/server/runProcess";
import {randomBytes} from "crypto"


export async function createRooms(transactionId:string) {
  console.log("createRooms")
  
  return runProcess("nats",["request","meeting-infrastructure.create","x"],600,transactionId)
 
}
export async function updateRooms(transactionId:string) {
  console.log("updateRooms")
  return runProcess("nats",["request","meeting-infrastructure.update","x"],600,transactionId)

 
}
export async function deleteRooms(transactionId:string) {
  console.log("deleteRooms")
  return runProcess("nats",["request","meeting-infrastructure.delete","x"],600,transactionId)
 
}

export async function getTransactionId() {
  return randomBytes(16).toString("hex")
}


  