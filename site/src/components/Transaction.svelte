<script>
    import MarketLink from './MarketLink.svelte'
    import TransactionType from "./TransactionType.svelte";
    import PartyLink from "./PartyLink.svelte";
    import OrderReferenceLink from "./OrderReferenceLink.svelte";
    import PriceByMarket from "./PriceByMarket.svelte";
    import TwoColumnData from "./TwoColumnData.svelte";

    export let tx;
    export let pubKey;
    export let type;

    function getPartyFromPubkey(id) {
        return id.substr(2)
    }

    let rows = []

    if (tx.marketId && tx.size && tx.price && pubKey) {
        /* It's an order */
        rows = [
            { title: 'Market', value: tx.marketId, type: 'market' },
            { title: 'Order reference', value: tx.reference, type: 'order-reference' },
            { title: 'Size', value: tx.size },
            { title: 'Party 🎉', value: pubKey, type: "party" },
            { title: 'Price', value: tx.price, type: 'price', marketId: tx.marketId }
        ]
    } else {
        /* Unknown type */
        for (const [key, value] of Object.entries(tx)) {
            rows.push({ title: key, value })
        }
    }

</script>

<details class="{pubKey}">
    <summary>
      &nbsp;&nbsp;<TransactionType type={type} /> {tx.size} @ <PriceByMarket price={tx.price} marketId={tx.marketID} /> in <MarketLink id={tx.marketId} /><br>
    </summary>
    <TwoColumnData rows={rows} />
</details>
