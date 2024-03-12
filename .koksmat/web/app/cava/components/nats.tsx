"use client";

import { useEffect, useMemo, useRef, useState } from "react";
import { connect, NatsConnection, StringCodec } from "nats.ws";

import { useConfig } from "./useconfig";

export default function ShowNatsLog(props: { subject: string }) {
  const { subject } = props;
  const {natsservers} = useConfig()
  const [nats, setNats] = useState<NatsConnection>()
  const [version, setversion] = useState(0);
  const [logEntries, setlogEntries] = useState<string[]>([])
  const v2 = useRef(0);
  const logEntries2 = useRef<string[]>([])
  
  useEffect(() => {
    if (natsservers.length===0) return;
    
    (async () => {
      const nc = await connect({
        servers:natsservers,
        name: "cava",
      });
      setNats(nc);
    })();

    return () => {
      nats?.drain();
      console.log("closed NATS connection");
    };
  }, [natsservers]);

  useEffect(() => {
    if (!nats) return;
    const sc = StringCodec();
    logEntries2.current = []
    const sub = nats.subscribe("log.1",{callback: (err, msg) => {
        if (err) {
            console.error(err);
            return;
        }
      
        const item = `[${sub.getProcessed()}]: ${sc.decode(msg.data)}`;
        const message = sc.decode(msg.data)
        console.log(item);
        logEntries2.current.push(message)
      
        v2.current = v2.current + 1;
        setversion(v2.current)
        
    }});


    console.log("connected to NATS");
    return () => {
        sub.unsubscribe();
      };
  }, [nats,subject]);
//return null
  return (
    <div className="">

      {logEntries2.current.map((entry, i) => (
      <div key={i}>
        {entry}</div>
    
      ))}
    </div>
  );
}
