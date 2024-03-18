import { Result } from "@/koksmat/httphelper";
import { AnyARecord } from "dns";
import { connect, Msg, Payload, ServiceError, StringCodec } from "nats.ws";

export interface INatsConfig {
  servers : string[]
}

export const runKoksmatProcess = (command: string, args: string[], timeout: number, channel: string, natsConfig : INatsConfig): Promise<Result<string>> => {
  return new Promise(async (resolve, reject) => {

    let result: Result<any> = {
      hasError: false,
      errorMessage: "",
      data: "",
    };
    try {

      const {servers} = natsConfig

      const nc = await connect({
        servers
      })
      

      const sc = StringCodec();
      const payload: Payload = JSON.stringify({ command, args, timeout, channel:"log."+channel });
      
       console.log("runKoksmatProcess", command, args, timeout, channel, payload)
    
      const msg = await nc.request(command, payload, { timeout: timeout * 1000 });
      result.data = sc.decode(msg.data);
      // if (!ServiceError.isServiceError(msg)) {
      //   result.data = sc.decode(msg.data);
      //   result.hasError = false;
      // } else {
    
      //   result.hasError = true;
      //   console.log("error", ServiceError.toServiceError(msg));
      //   result.errorMessage =ServiceError.toServiceError(msg)?.message
      // }
      
   


    } catch (error ) {
      const e = error as Error;
      
      result.hasError = true;
      console.log("error", e.message ?? error);
      result.errorMessage = "NATS error";

    }


    resolve(result);

  });
};
