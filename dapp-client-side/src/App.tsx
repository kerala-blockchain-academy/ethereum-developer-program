import { ChangeEvent, FormEvent, useState } from 'react'
import './App.css'
import { createWalletClient, custom } from 'viem'
import { contractAddress } from './deployed_addresses.json'
import { abi } from './Cert.json'
import {
  readContract,
  waitForTransactionReceipt,
  writeContract,
} from 'viem/actions'
import { hardhat } from 'viem/chains'

function App() {
  const [output, setOutput] = useState('')
  const [queryID, setQueryID] = useState(0)

  const [formData, setFormData] = useState({
    id: 0,
    name: '',
    course: '',
    grade: '',
    date: '',
  })

  const client = createWalletClient({
    chain: hardhat,
    transport: custom(window.ethereum),
  })

  const connectMetaMask = async () => {
    const accounts = await client.requestAddresses()
    console.log(accounts)

    alert(`Successfully Connected ${accounts[0]}`)
  }

  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target
    setFormData((prevState) => ({ ...prevState, [name]: value }))
  }

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault()

    console.log(formData)

    const accounts = await client.getAddresses()
    const hash = await writeContract(client, {
      address: contractAddress as `0x${string}`,
      abi: abi,
      functionName: 'issue',
      args: [
        formData.id,
        formData.name,
        formData.course,
        formData.grade,
        formData.date,
      ],
      account: accounts[0],
    })
    console.log('Transaction Hash:', hash)

    await waitForTransactionReceipt(client, { hash })

    resetForm()
  }

  const resetForm = () => {
    setFormData({ id: 0, name: '', course: '', grade: '', date: '' })
  }

  const getCertificate = async () => {
    let result = (await readContract(client, {
      address: contractAddress as `0x${string}`,
      abi: abi,
      functionName: 'Certificates',
      args: [queryID],
    })) as string[]

    if (result) {
      console.log(result)
      setOutput(
        `Name: ${result[0]}, Course: ${result[1]}, Grade: ${result[2]}, Date: ${result[3]}`
      )
    }
  }

  return (
    <div>
      <h1>Certificate DApp</h1>
      <button onClick={connectMetaMask}>Connect MetaMask</button>
      <br />
      <br />
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor='id'>ID: </label>
          <input
            type='number'
            id='id'
            name='id'
            value={formData.id}
            onChange={handleChange}
          />
        </div>
        <div>
          <label htmlFor='name'>Name: </label>
          <input
            type='text'
            id='name'
            name='name'
            value={formData.name}
            onChange={handleChange}
          />
        </div>
        <div>
          <label htmlFor='course'>Course: </label>
          <input
            type='text'
            id='course'
            name='course'
            value={formData.course}
            onChange={handleChange}
          />
        </div>
        <div>
          <label htmlFor='grade'>Grade: </label>
          <input
            type='text'
            id='grade'
            name='grade'
            value={formData.grade}
            onChange={handleChange}
          />
        </div>
        <div>
          <label htmlFor='date'>Date: </label>
          <input
            type='date'
            id='date'
            name='date'
            value={formData.date}
            onChange={handleChange}
          />
        </div>
        <div>
          <button type='submit'>Submit</button>
          <button type='button' onClick={resetForm}>
            Reset
          </button>
        </div>
      </form>
      <br />
      <br />
      <div>
        <label htmlFor='queryID'>Query ID: </label>
        <input
          type='number'
          id='queryID'
          name='queryID'
          value={queryID}
          onChange={(e) => setQueryID(Number(e.target.value))}
        />
      </div>
      <button onClick={getCertificate}>Get</button>
      <p>{output}</p>
    </div>
  )
}

export default App
