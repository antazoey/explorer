<script>
    import MarketLink from './MarketLink.svelte'
    export let tx;
    export let pubKey;
    let kv = Object.entries(tx)

    function getPartyFromPubkey(id) {
        return id.substr(2)
    }

</script>

<details>
    <summary>{tx.reference}</summary>
    <dl>
        <dt>PubKey</dt>
        <dd><a href='{`/party/${getPartyFromPubkey(pubKey)}`}'>{pubKey}</a></dd>

        {#each kv as k}
        {#if k[0]!== 'partyID'}
        <dt>{k[0]}</dt>
          {#if k[0]=== 'marketID'}
          <dd><MarketLink id={k[1]} /></dd>
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
