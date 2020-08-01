<script>
    import Hash from "./Hash.svelte";
    export let block;
</script>
<style>

    dl {
        display: flex;
        flex-flow: row wrap;
        border: solid #333;
        border-width: 1px 1px 0 1px;
    }
    dt {
        flex-basis: 20%;
        padding: 2px 4px;
        text-align: right;
        background: #555555;
        color: #fff;
        border-bottom: 1px solid #333;
        border-right: 1px solid #333;
    }
    dd {
        overflow: hidden;
        flex-basis: 70%;
        flex-grow: 1;
        margin: 0;
        padding: 2px 4px;
        border-bottom: 1px solid #333;
    }
</style>

{#if block}
<dl>
    <dt>Block ID</dt><dd>{block.header.height}</dd>
    <dt>Chain</dt><dd><Hash text={block.header.chain_id} /></dd>
    <dt>App hash</dt><dd><Hash text={block.header.app_hash} /></dd>
    <dt>Consensus hash</dt><dd><Hash text={block.header.consensus_hash} /></dd>
    <dt>Data hash</dt><dd><Hash text={block.header.data_hash} /></dd>
    <dt>Block hash</dt><dd><Hash text={block.block_id.hash} /></dd>
    <dt>Proposer</dt><dd><Hash text={block.header.proposer_address} /></dd>
    <dt>Time</dt><dd>{block.header.time}</dd>
    <dt>Transactions</dt><dd>{block.header.num_txs}</dd>
</dl>
{/if}