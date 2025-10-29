import { createClient, getContract, http } from 'viem'
import { hardhat } from 'viem/chains'
import details from './deployed_addresses.json'
import Cert from './Cert.json'
import { privateKeyToAccount } from 'viem/accounts'

export const client = createClient({
  chain: hardhat,
  account: privateKeyToAccount(process.env.PRIVATE_KEY as `0x${string}`),
  transport: http(process.env.HTTP_URL),
})

export const contract = getContract({
  abi: Cert.abi,
  address: details.contract as `0x${string}`,
  client,
})
