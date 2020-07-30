<script>
  export let id;
  import { store } from "../stores/markets.js";

  import { onDestroy, onMount } from 'svelte'

  let unsubscribe
  let market

  onDestroy(() => {
    if(unsubscribe) {
      unsubscribe();
      unsubscribe = null;
    }
  });

  function updateMarkets(data) {
    market = data.markets.get(id);
  }

  onMount (() => {
    if(!unsubscribe) {
      unsubscribe = store.subscribe(updateMarkets);
    }
  })
</script>

<div>
  {#if market}
  <h1>{market.id}</h1>
  <pre>{JSON.stringify(market.tradableInstrument, null, 2)}</pre>
    {:else}
      <h1>Loading...</h1>
    {/if}
</div>