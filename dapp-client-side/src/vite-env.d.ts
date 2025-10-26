/// <reference types="vite/client" />

import { EIP1193Provider } from "viem";

declare global {
  interface Window {
    ethereum: EIP1193Provider
  }
}
