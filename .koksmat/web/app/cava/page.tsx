"use client"
import { Button } from "@/components/ui/button"
import Logo from "@/koksmat/components/logo"
import { useEffect, useState } from "react"
import { updateRooms } from "./server"
import {
  connect,
  NatsConnection,
  StringCodec,
} from "nats.ws";



export const dynamic = "force-dynamic"

function sleep(ms: number) {
  return new Promise(resolve => setTimeout(resolve,ms));
}

export default  function RedirectToLoggedinUse() {
  const [nats, setNats] = useState<NatsConnection>();
  const [running, setrunning] = useState(false)
  const [result, setresult] = useState("")
  const [error, seterror] = useState("")
  const [logEntries, setlogEntries] = useState<string[]>([])
// https://github.com/nats-io/nats.ws

  useEffect(() => {
    (async () => {
      const nc = await connect({
        servers: ["ws://0.0.0.0:443"],
      })
      setNats(nc)
      const sub = nc.subscribe("cava.*");
      const sc = StringCodec();
      (async () => {
        for await (const m of sub) {

          setresult(sc.decode(m.data))
          logEntries.push(`[${sub.getProcessed()}]: ${sc.decode(m.data)}`)
          setlogEntries(logEntries)
        }
        console.log("subscription closed");
      })();

      console.log("connected to NATS")
    })();

    return () => {
      nats?.drain();
      console.log("closed NATS connection")
    }
  }, [])
  const run = async () => {
    setrunning(true)
    const result = await updateRooms()
    
    await sleep(3000)
    setresult(result)
    setrunning(false)
  }
  return (
    <div className="space-y-3">
      <div>
      <div>
       <Button disabled={running} onClick={()=>run()}>Create Rooms</Button>
       </div>
       <div>
       <Button disabled={running} onClick={()=>run()}>Update Rooms</Button>
       </div>
       <div>
       <Button disabled={running} onClick={()=>run()} variant={"destructive"}>Delete Rooms</Button>
       </div>
       <div>
       <Button disabled={running} onClick={()=>run()}>Change Email</Button>
       </div>
       <div>
       <Button disabled={running} onClick={()=>run()}>Run all</Button>
       </div>       
       </div>
     {running && <div>Running...</div>}
      {result && <div>Result: {result}</div>}
      {error && <div className="text-red-500">Error: {error}</div>}
      <pre>
        {JSON.stringify(logEntries, null, 2)}
      </pre>
      {logEntries.map((entry, i) => (
        <div key={i}>{entry}</div>
      ))}
    </div>
  )
}
