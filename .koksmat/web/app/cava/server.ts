"use server"

import { runKoksmatProcess } from "@/koksmat/server/runKoksmatProcess";
import {randomBytes} from "crypto"

export async function createRooms(transactionId:string) {
  console.log("createRooms")
  return runKoksmatProcess("meeting-infrastructure.create",[], 600,transactionId)
 
}
export async function updateRooms(transactionId:string) {
  console.log("updateRooms")
  return runKoksmatProcess("meeting-infrastructure.update",[], 600,transactionId)
 
}
export async function deleteRooms(transactionId:string) {
  console.log("deleteRooms")
  return runKoksmatProcess("meeting-infrastructure.delete",[], 600,transactionId)
 
}

export async function getTransactionId() {
  return randomBytes(16).toString("hex")
}


  export async function natsconfig() {
    return {
      servers: ["ws://0.0.0.0:443"]
    }
  }

