"use client"
import { Button } from "@/components/ui/button"
import Logo from "@/koksmat/components/logo"
import { useContext, useEffect, useState } from "react"
import { createRooms, deleteRooms, getTransactionId, updateRooms } from "./server"
import ShowNatsLog from "./components/nats"
import { set } from "date-fns"
import { Result } from "@/koksmat/httphelper"
import { MagicboxContext } from "@/koksmat/magicbox-context"



export const dynamic = "force-dynamic"

function sleep(ms: number) {
  return new Promise(resolve => setTimeout(resolve,ms));
}

export default  function CavaHome() {
  const [running, setrunning] = useState(false)
  const [result, setresult] = useState("")
  const [error, seterror] = useState("")
  const [logEntries, setlogEntries] = useState<string[]>([])
  const [transactionId, settransactionid] = useState("")

  const magicbox = useContext(MagicboxContext)

  const run = async (method : (id:string)=>Promise<Result<any>>) => {
    

    setrunning(true)
    setresult("")
    seterror("")

    const transactionId = await getTransactionId()
    settransactionid(transactionId)
    const result = await method(transactionId)
    

    setresult(result.data || "")
    seterror(result.errorMessage || "")
    setrunning(false)
  }

  const doUpdate = async () => run(updateRooms)
  const doCreate = async () => run(createRooms)
  const doDelete = async () => run(deleteRooms)
  return (
    <div >
      <div className="space-y-3 p-4">
        {magicbox?.user?.name && <div>Logged in as {magicbox.user.name}</div>}
      <div>
       <Button disabled={running} onClick={()=>doCreate()}>Create Rooms</Button>
       </div>
       <div>
       <Button disabled={running} onClick={()=>doUpdate()}>Update Rooms</Button>
       </div>
       <div>
       <Button disabled={running} onClick={()=>doDelete()} variant={"destructive"}>Delete Rooms</Button>
       </div>
       {/* <div>
       <Button disabled={running} onClick={()=>run()}>Change Email</Button>
       </div>
       <div>
       <Button disabled={running} onClick={()=>run()}>Run all</Button>
       </div>        */}
       </div>
     {running && <div>Running... </div>}
      {result && <div>Result: {result}</div>}
      {error && <div className="text-red-500">Error: {error}</div>}
      {transactionId &&
      <div>
        transactionId: {transactionId}
     <ShowNatsLog  subject={"log."+transactionId} />
     </div>
     }
    </div>
  )
}
