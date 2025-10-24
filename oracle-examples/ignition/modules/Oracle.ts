import { buildModule } from '@nomicfoundation/hardhat-ignition/modules'

export default buildModule('OracleModule', (m) => {
  const oV1 = m.contract('OracleV1')
  const oV2 = m.contract('OracleV2')

  return { oV1, oV2 }
})
