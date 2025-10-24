import axios from 'axios'
import oracleV2 from './artifacts/contracts/OracleV2.sol/OracleV2.json'
import details from './ignition/deployments/chain-31337/deployed_addresses.json'
import { hardhat } from 'viem/chains'
import { createPublicClient, parseAbiItem, webSocket } from 'viem'
import { waitForTransactionReceipt, writeContract } from 'viem/actions'
import { privateKeyToAccount } from 'viem/accounts'

const contractAddress = details['OracleModule#OracleV2'] as `0x${string}`

const client = createPublicClient({
  chain: hardhat,
  transport: webSocket('ws://127.0.0.1:8545'),
})

;() => {
  console.log('Listening for events...')
  client.watchEvent({
    address: contractAddress,
    event: parseAbiItem('event PriceDataRequest()'),
    onLogs: async (logs) => {
      console.log(`${logs[0].eventName} occured!!`)

      let data = await axios.get(
        'https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=BTC,USD,EUR'
      )

      let usd = JSON.stringify(data.data.USD)
      console.log(`Read data from truth point. USD rate is $${usd}`)

      let hash = await writeContract(client, {
        address: contractAddress,
        abi: oracleV2.abi,
        functionName: 'updatePriceData',
        args: [usd],
        account: privateKeyToAccount(
          '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80'
        ),
      })

      console.log(`Tx: ${hash}`)
      await waitForTransactionReceipt(client, { hash })
      console.log('Updated on-chain Oracle contract!!')
    },
  })
}
