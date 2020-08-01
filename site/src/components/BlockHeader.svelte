<script>
    import TwoColumnData from "./TwoColumnData.svelte";
    export let block;

    const rows = [
        { title: 'Block ID', value: block.header.height },
        { title: 'Chain', value: block.header.chain_id, type: 'hash' },

        { title: 'App hash', value: block.header.app_hash, type: 'hash' },
        { title: 'Consensus hash', value: block.header.consensus_hash, type: 'hash' },
        { title: 'Data hash', value: block.header.data_hash, type: 'hash' },
        { title: 'Block hash', value: block.block_id.hash, type: 'hash' },

        { title: 'Proposer', value: block.header.proposer_address, type: 'hash' },
        { title: 'Time', value: block.header.time },

        { title: 'Transactions', value: block.header.num_txs },
   ]
</script>

<TwoColumnData rows={rows} />
