import { mkdirSync, readFileSync, writeFileSync } from 'fs'
import { dirname, join, resolve } from 'path'
import solc from 'solc'

const CONTRACT_PATH = '../../lib/contracts/'
const ARTIFACT_PATH = '../../lib/artifacts/'

const CONTRACT_FILES = ['Cert', 'GovToken', 'MyGovernor', 'TimeLock']

const findImports = (relativePath) => {
  let absolutePath = resolve(CONTRACT_PATH, relativePath)
  if (relativePath.startsWith('@openzeppelin')) {
    absolutePath = resolve(join(CONTRACT_PATH, 'node_modules'), relativePath)
  }

  const source = readFileSync(absolutePath, 'utf8')
  return { contents: source }
}

for (const contract of CONTRACT_FILES) {
  const content = readFileSync(`${CONTRACT_PATH}${contract}.sol`).toString()

  const input = {
    language: 'Solidity',
    sources: {
      [contract]: {
        content: content,
      },
    },
    settings: {
      optimizer: {
        enabled: true,
        runs: 200,
      },
      outputSelection: {
        '*': {
          '*': ['*'],
        },
      },
    },
  }

  const compiled = solc.compile(JSON.stringify(input), { import: findImports })
  const output = JSON.parse(compiled)

  const abi = output.contracts[contract][contract].abi
  const bytecode = output.contracts[contract][contract].evm.bytecode.object

  const abiPath = join(ARTIFACT_PATH, `${contract}.json`)
  mkdirSync(dirname(abiPath), { recursive: true })

  writeFileSync(abiPath, JSON.stringify(abi, null, 2))
  console.log(`Saved ${contract}.json!`)

  const binPath = join(ARTIFACT_PATH, `${contract}.bin`)
  mkdirSync(dirname(binPath), { recursive: true })

  writeFileSync(binPath, bytecode)
  console.log(`Saved ${contract}.bin!`)
}
