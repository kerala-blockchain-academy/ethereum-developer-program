import { buildModule } from '@nomicfoundation/hardhat-ignition/modules'

export default buildModule('CertModule', (m) => {
  const cert = m.contract('Cert')

  return { cert }
})
