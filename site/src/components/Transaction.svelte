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

    if (tx.marketID && tx.partyID && tx.size && tx.price) {
        /* It's an order */
        rows = [
            { title: 'Market', value: tx.marketID, type: 'market' },
            { title: 'Party', value: tx.partyID, type: 'party' },
            { title: 'Order reference', value: tx.reference, type: 'order-reference' },
            { title: 'Size', value: tx.size },
            { title: 'Price', value: tx.price, type: 'price', marketId: tx.marketID }
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
      &nbsp;&nbsp;<TransactionType type={type} /> {tx.size} @ <PriceByMarket price={tx.price} marketId={tx.marketID} /> in <MarketLink id={tx.marketID} /><br>
    </summary>
    <TwoColumnData rows={rows} />
</details>
