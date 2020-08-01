<script>
    import MarketLink from './MarketLink.svelte'
    import TransactionType from "./TransactionType.svelte";
    import PartyLink from "./PartyLink.svelte";
    import OrderLink from "./OrderLink.svelte";
    import OrderReferenceLink from "./OrderReferenceLink.svelte";

    export let tx;
    export let pubKey;
    export let type;
    let kv = Object.entries(tx)

    function getPartyFromPubkey(id) {
        return id.substr(2)
    }
</script>
<style>

dl {
    margin: 5px 0;
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
flex-basis: 70%;
flex-grow: 1;
margin: 0;
padding: 2px 4px;
border-bottom: 1px solid #333;
}
</style>
<details class="{pubKey}">
    <summary>
      &nbsp;&nbsp;<TransactionType type={type} /> {tx.size} @ {tx.price || 'MARKET'} in <MarketLink id={tx.marketID} /><br>
    </summary>
    <dl>
        <dt>PubKey</dt>
        <dd><PartyLink id={pubKey} /></dd>

        {#each kv as k}
        {#if k[0]!== 'partyID'}
        <dt>{k[0]}</dt>
          {#if k[0]=== 'marketID'}
          <dd><MarketLink id={k[1]} /></dd>
          {:else if k[0]=== 'reference'}
          <dd><OrderReferenceLink reference={k[1]} /></dd>
          {:else if k[0]=== 'terms'}
            <dd>
              <pre>{JSON.stringify(k[1])}</pre>
            </dd>
          {:else}
            <dd>{k[1]}</dd>
          {/if}
        {/if}
        {/each}
    </dl>
</details>
