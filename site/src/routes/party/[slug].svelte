<script>
  import MarketLink from '../../components/MarketLink.svelte'
  let accounts = false

</script>

<div>
  {#if accounts}
    <ol>
    {#each accounts as { asset, balance, marketID, type}, i}
      {#if !marketID}
      <li>{asset}: {balance}</li>
      {:else}
      <li><MarketLink id={marketID} />: {asset}: {balance}</li>
      {/if}
    {/each}
    </ol>
  {:else}
      <h1>Soon...</h1>
  {/if}
</div>
