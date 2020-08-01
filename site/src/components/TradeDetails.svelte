<script>
    import Hash from "./Hash.svelte";
    import MarketLink from "./MarketLink.svelte";
    import PartyLink from "./PartyLink.svelte";
    import OrderLink from "./OrderLink.svelte";
    import BlockLink from "./BlockLink.svelte";
    import PriceByMarket from "./PriceByMarket.svelte";

    export let trade;
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

{#if trade}
    <dl>
        <dt>ID</dt><dd><Hash text={trade.id} /></dd>
        <dt>Block</dt><dd><BlockLink id={trade.id} /></dd>
        <dt>Buyer</dt><dd><PartyLink id={trade.buyer.id} /></dd>
        <dt>Seller</dt><dd><PartyLink id={trade.seller.id} /></dd>
        <dt>Sell Order</dt><dd><OrderLink id={trade.sellOrder} /></dd>
        <dt>Buy Order</dt><dd><OrderLink id={trade.buyOrder} /></dd>
        <dt>Market</dt><dd><MarketLink id={trade.market.id} /></dd>
        <dt>Size</dt><dd>{trade.size}</dd>
        <dt>Price</dt><dd><PriceByMarket price={trade.price} marketId={trade.market.id} /></dd>
    </dl>
{/if}