const net = 'n'
const tendermintBaseUrl = `https://n08.${net}.vega.xyz/tm/`
const apiBaseUrl = `https://n08.${net}.vega.xyz/`
const blockExplorerBaseUrl = 'https://explorer.vega.trading/.netlify/functions/chain-explorer-api'
const topgunBaseUrl = 'https://topgun-service-testnet.ops.vega.xyz/leaderboard'
function apiUrl(path = ''){ return `${apiBaseUrl}${path}` }
function tendermintUrl(path = ''){ return `${tendermintBaseUrl}${path}`}
function blockUrl(path = '') { return `${blockExplorerBaseUrl}${path}` }
function topgunUrl() { return `${topgunBaseUrl}`}
export {
    blockExplorerBaseUrl,
    tendermintBaseUrl,
    apiBaseUrl,
    apiUrl,
    tendermintUrl,
    blockUrl,
    topgunUrl
}