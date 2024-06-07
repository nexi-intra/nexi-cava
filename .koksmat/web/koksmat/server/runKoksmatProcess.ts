import { Result } from "@/koksmat/httphelper";
import { AnyARecord } from "dns";

export interface INatsConfig {
  servers: string[];
}

export const runKoksmatProcess = (
  command: string,
  args: string[],
  timeout: number,
  channel: string,
  natsConfig: INatsConfig
): Promise<Result<string>> => {
  return new Promise(async (resolve, reject) => {
    let result: Result<any> = {
      hasError: false,
      errorMessage: "",
      data: "",
    };
    try {
    } catch (error) {
      const e = error as Error;

      result.hasError = true;
      console.log("error", e.message ?? error);
      result.errorMessage = "NATS error";
    }

    resolve(result);
  });
};
