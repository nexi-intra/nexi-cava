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
  const {user} = magicbox

  const run = async (method : (id:string)=>Promise<Result<any>>) => {
    

    setrunning(true)
    setresult("")
    seterror("")

    const result = await method(transactionId)
    

    setresult(result.data || "")
    seterror(result.errorMessage || "")
    setrunning(false)
  }

  const doUpdate = async () => run(updateRooms)
  const doCreate = async () => run(createRooms)
  const doDelete = async () => run(deleteRooms)

  useEffect(() => {
   const load = async () => {
    if (!user?.email) return;
    const transactionId = await getTransactionId()
    settransactionid(user.email+"."+transactionId)
   }
  
   load()
  }, [user])
  if (!user?.email) return <div><Button

onClick={async () => {
  const signedIn = await magicbox.signIn(["User.Read"], "");
  
}}
>
Sign In
</Button></div>
  return (
    <div >
      <div className="space-y-3 p-4">
        {user?.name && <div>Logged in as {user.name}</div>}
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
        Copy this command to your terminal to write to "me"
        <div className="mt-4">
        <textarea  value={"nats publish log."+transactionId+" test"} className="w-[100%] font-mono text-sm">
         
         </textarea>
         </div>
     <ShowNatsLog  subject={"log."+transactionId} />
     </div>
     }
    </div>
  )
}
