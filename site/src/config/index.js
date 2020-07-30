const net = 's'
const tendermintBaseUrl = `https://geo.${net}.vega.xyz:8443/`
const apiBaseUrl = `https://n04.${net}.vega.xyz/`
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