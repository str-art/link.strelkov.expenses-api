"use client";
import { ReactNode, createContext, useContext } from "react";
import axios, { Axios } from "axios";

type ClientContextType = {
  client: Axios;
};

const ClientContext = createContext<ClientContextType>({} as ClientContextType);

export const useClientContext = () => useContext(ClientContext);

type ClientsProviderProps = {
  children: ReactNode;
  apiUrl: string;
};

export default function ClientsProvider({
  children,
  apiUrl,
}: ClientsProviderProps) {
  const client = axios.create({
    baseURL: apiUrl,
    headers: {
      "Content-Type": "application/json",
    },
  });

  return (
    <ClientContext.Provider
      value={{
        client,
      }}
    >
      {children}
    </ClientContext.Provider>
  );
}
