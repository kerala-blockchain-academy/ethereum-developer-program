import 'dotenv/config'
import { createPublicClient, parseAbiItem, webSocket } from 'viem'
import { hardhat } from 'viem/chains'
import details from './lib/deployed_addresses.json'

const client = createPublicClient({
  chain: hardhat,
  transport: webSocket(process.env.WS_URL),
})

console.log('Listening for Issued Events...')

client.watchEvent({
  address: details.contract as `0x${string}`,
  event: parseAbiItem(
    'event Issued(string indexed course, uint256 id, string grade)'
  ),
  onLogs: (logs) => {
    console.log('================ EVENT OCCURED ================')
    console.log('Course: ', logs[0].args.course)
    console.log('ID: ', logs[0].args.id)
    console.log('Grade: ', logs[0].args.grade)
    console.log('===============================================')
  },
})
