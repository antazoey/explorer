
const networkConfigs = {
    dev: {
        tendermintBaseUrl: `https://n04.d.vega.xyz/tm/`,
        apiBaseUrl: `https://n04.d.vega.xyz/`
    },
    staging: {
        tendermintBaseUrl: `https://n03.s.vega.xyz/tm/`,
        apiBaseUrl: `https://n03.s.vega.xyz/`
    },
    testnet: {
        tendermintBaseUrl: `https://lb.testnet.vega.xyz/tm/`,
        apiBaseUrl: `https://lb.testnet.vega.xyz/`
    }
}

let currentNetwork = 'testnet'
if (typeof localStorage !== 'undefined' && localStorage.getItem('network')) {
 currentNetwork = localStorage.getItem('network') 
}

const blockExplorerBaseUrl = 'https://explorer.vega.trading/.netlify/functions/chain-explorer-api'

function getApiBaseUrl(currentNetwork) {
    return networkConfigs[currentNetwork].apiBaseUrl
}

function getTmUrl(currentNetwork) {
    return networkConfigs[currentNetwork].tendermintBaseUrl
}


function apiUrl(path = ''){ return `${getApiBaseUrl(currentNetwork)}${path}` }
function apiWsUrl(path = ''){ return `${getApiBaseUrl(currentNetwork).replace('https', 'wss')}${path}` }
function tendermintUrl(path = ''){ return `${getTmUrl(currentNetwork)}${path}`}
function blockUrl(path = '') { return `${blockExplorerBaseUrl}${path}` }
export {
    blockExplorerBaseUrl,
    apiUrl,
    apiWsUrl,
    tendermintUrl,
    blockUrl,
}
