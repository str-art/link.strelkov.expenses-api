import { useClientContext } from "@/contexts/ClientProvider";

export default function useClient() {
  const { client } = useClientContext();
  return client;
}
