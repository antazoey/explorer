<script>
    import Hash from "./Hash.svelte";
    import MarketLink from "./MarketLink.svelte";
    import PartyLink from "./PartyLink.svelte";

    export let order;
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

{#if order}
    <dl>
        <dt>Status</dt><dd>{order.status}</dd>
        <dt>Party</dt><dd><PartyLink id={order.party.id} /></dd>
        <dt>ID</dt><dd><Hash text={order.id} /></dd>
        <dt>Reference</dt><dd><Hash text={order.reference} /></dd>
        <dt>Market</dt><dd><MarketLink id={order.market.id} /></dd>
        <dt>Size (remaining)</dt><dd>{order.size} ({order.remaining})</dd>
        <dt>Price</dt><dd>{order.price}</dd>
        <dt>Side</dt><dd>{order.side}</dd>
        <dt>Time in Force</dt><dd>{order.timeInForce}</dd>
        <dt>Type</dt><dd>{order.type}</dd>
        <dt>Trades</dt><dd>{order.trades.length}</dd>
    </dl>
{/if}