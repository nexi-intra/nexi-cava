"use client";

import { use, useContext, useEffect, useState } from "react";

import { natsconfig } from "../server";

export const version = 1;

export function useConfig() {
  const [natsservers, setnatsservers] = useState<string[]>([]);

  useEffect(() => {
    const load = async () => {
      const natsConfig = await natsconfig();
      setnatsservers(natsConfig.servers);
    };
    load();
  }, []);

  return {
    natsservers,
  };
}
