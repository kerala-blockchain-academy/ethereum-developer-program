import assert from 'node:assert/strict'
import { describe, it } from 'node:test'
import { network } from 'hardhat'
import { getCreateAddress, keccak256, toHex } from 'viem'

describe('Cert', async () => {
  const { viem, networkHelpers } = await network.connect()

  async function deployCertFixture() {
    const [admin, other] = await viem.getWalletClients()
    const cert = await viem.deployContract('Cert')
    const client = await viem.getPublicClient()

    return { cert, admin, other, client }
  }

  it('Should set the right admin', async () => {
    const { cert, admin } = await networkHelpers.loadFixture(deployCertFixture)

    const contractAddress = getCreateAddress({
      from: admin.account.address,
      nonce: 0n,
    })

    assert.equal(cert.address, contractAddress.toLowerCase())
  })

  it('Should issue the certificate', async () => {
    const { cert, client } = await networkHelpers.loadFixture(deployCertFixture)

    const hash = await cert.write.issue([
      14n,
      'Deren',
      'MBCC',
      'S',
      '30-05-2025',
    ])
    await client.waitForTransactionReceipt({ hash })

    const event = await cert.getEvents.Issued()

    assert.equal(event.length, 1)
    assert.equal(event[0].args.course, keccak256(toHex('MBCC')))
    assert.equal(event[0].args.id, 14n)
    assert.equal(event[0].args.grade, 'S')
  })

  it('Should read the certificate', async () => {
    const { cert, client } = await networkHelpers.loadFixture(deployCertFixture)

    const hash = await cert.write.issue([
      885n,
      'Shawn',
      'MBCC',
      'A',
      '28-05-2025',
    ])
    await client.waitForTransactionReceipt({ hash })

    const certificate = await cert.read.Certificates([885n])

    assert.equal(certificate[0], 'Shawn')
    assert.equal(certificate[1], 'MBCC')
    assert.equal(certificate[2], 'A')
    assert.equal(certificate[3], '28-05-2025')
  })

  it('Should revert the issuing', async () => {
    const { cert, other } = await networkHelpers.loadFixture(deployCertFixture)

    const certWithOther = await viem.getContractAt('Cert', cert.address, {
      client: { wallet: other },
    })

    await viem.assertions.revertWith(
      certWithOther.write.issue([355n, 'Lisa', 'MBCC', 'B', '31-05-2025']),
      'Access Denied'
    )
  })
})
