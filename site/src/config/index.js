const net = 's'
const tendermintBaseUrl = `https://geo.${net}.vega.xyz:8443/`
const apiBaseUrl = `https://geo.${net}.vega.xyz/`
const blockExplorerBaseUrl = 'https://explorer.vega.trading/.netlify/functions/chain-explorer-api'
function apiUrl(path = ''){ return `${apiBaseUrl}${path}` }
function tendermintUrl(path = ''){ return `${tendermintBaseUrl}${path}`}
function blockUrl(path = '') { return `${blockExplorerBaseUrl}${path}` }

export {
    blockExplorerBaseUrl,
    tendermintBaseUrl,
    apiBaseUrl,
    apiUrl,
    tendermintUrl,
    blockUrl
}