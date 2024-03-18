"use server"

import { INatsConfig, runKoksmatProcess } from "@/koksmat/server/runKoksmatProcess";
import {randomBytes} from "crypto"

async function Run(cmd:string,args:string[],transactionId:string) { 
  const natsConfig = await natsconfig()
  return runKoksmatProcess(cmd,args,600,transactionId,natsConfig)
}

export async function createRooms(transactionId:string) {
  console.log("createRooms")
  
  return Run("meeting-infrastructure.create",[], transactionId)
 
}
export async function updateRooms(transactionId:string) {
  console.log("updateRooms")
  return Run("meeting-infrastructure.update",[], transactionId)
 
}
export async function deleteRooms(transactionId:string) {
  console.log("deleteRooms")
  return Run("meeting-infrastructure.delete",[], transactionId)
 
}

export async function getTransactionId() {
  return randomBytes(16).toString("hex")
}


  export async function natsconfig() : Promise<INatsConfig>{
    let natsConnectionString = process.env.NATS;
    if (!natsConnectionString) {
      natsConnectionString = "ws://0.0.0.0:443"
    }
    console.log("natsconfig",natsConnectionString)
    const natsConnections : string[] = natsConnectionString.split(",")
    return {
      servers: natsConnections
    }
  }

