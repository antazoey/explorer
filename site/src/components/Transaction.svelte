<script>
    import Market from './Market.svelte'
    export let tx;
    export let pubKey;
    let kv = Object.entries(tx)
</script>

<details>
    <summary>{tx.reference}</summary>
    <dl>
        <dt>PubKey</dt>
        <dd>{pubKey}</dd>
        {#each kv as k}
        <dt>{k[0]}</dt>
        {#if k[0]=== 'marketID'}
          <dd><Market id={k[1]} /></dd>
        {:else}
          {#if k[0]=== 'partyID'}
            <dd><a href='/party/{k[1]}'>{k[1]}</a></dd>
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
