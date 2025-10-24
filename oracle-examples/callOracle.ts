import axios from 'axios'
import { createClient, http } from 'viem'
import oracleV1 from './artifacts/contracts/OracleV1.sol/OracleV1.json'
import details from './ignition/deployments/chain-31337/deployed_addresses.json'
import { hardhat } from 'viem/chains'
import { privateKeyToAccount } from 'viem/accounts'
import { waitForTransactionReceipt, writeContract } from 'viem/actions'

const timer = (ms: number) => new Promise((res) => setTimeout(res, ms))

const client = createClient({
  chain: hardhat,
  transport: http(),
  account: privateKeyToAccount(
    '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80'
  ),
})

;(async () => {
  while (true) {
    let data = await axios.get(
      'https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=BTC,USD,EUR'
    )

    let usd = JSON.stringify(data.data.USD)
    console.log(`Read data from truth point. USD rate is $${usd}`)

    let hash = await writeContract(client, {
      address: details['OracleModule#OracleV1'] as `0x${string}`,
      abi: oracleV1.abi,
      functionName: 'setPriceData',
      args: [usd],
      account: privateKeyToAccount(
        '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80'
      ),
    })

    console.log(`Tx: ${hash}`)
    await waitForTransactionReceipt(client, { hash })
    console.log('Updated on-chain Oracle contract!!')

    await timer(30000)
  }
})()
