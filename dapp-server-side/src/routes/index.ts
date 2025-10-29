import { Router } from 'express'
import { waitForTransactionReceipt } from 'viem/actions'
import { client, contract } from '../lib/instance'
const router = Router()

/* GET home page. */
router.get('/', function (req, res) {
  res.send('Hello World!')
})

router.post('/issue', async (req, res) => {
  try {
    const hash = await contract.write.issue([
      req.body.id,
      req.body.name,
      req.body.course,
      req.body.grade,
      req.body.date,
    ])

    const trx = await waitForTransactionReceipt(client, { hash })

    console.log(trx)
    res.json(trx)
  } catch (error) {
    console.log(error)
    res.json(error)
  }
})

router.get('/fetch', async (req, res) => {
  try {
    const result = (await contract.read.Certificates([
      req.query.id,
    ])) as string[]

    console.log(result)
    res.json({
      id: req.query.id,
      name: result[0],
      course: result[1],
      grade: result[2],
      date: result[3],
    })
  } catch (error) {
    console.log(error)
    res.json(error)
  }
})

export default router
