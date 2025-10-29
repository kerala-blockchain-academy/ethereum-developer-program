import { Router } from 'express'
import { createPublicClient, http, parseAbiItem } from 'viem'
import { hardhat } from 'viem/chains'
import details from '../lib/deployed_addresses.json'
const router = Router()

router.get('/', async (req, res) => {
  try {
    let events
    if (req.query.course) {
      const client = createPublicClient({
        chain: hardhat,
        transport: http(process.env.HTTP_URL),
      })

      const filter = await client.createEventFilter({
        address: details.contract as `0x${string}`,
        event: parseAbiItem(
          'event Issued(string indexed course, uint256 id, string grade)'
        ),
        fromBlock: 'earliest',
        args: { course: req.query.course.toString() },
      })

      events = await client.getFilterLogs({ filter })
    }

    res.json(events)
  } catch (error) {
    console.log(error)
    res.json(error)
  }
})

export default router
