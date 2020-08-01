<script>
  import { store } from "../../stores/markets.js";

  import { onDestroy, onMount } from 'svelte'

  let unsubscribe
  let markets

  onDestroy(() => {
    if(unsubscribe) {
      unsubscribe();
      unsubscribe = null;
    }
  });

  function updateMarkets(data) {
    markets = data;
  }

  onMount (() => {
    if(!unsubscribe) {
      unsubscribe = store.subscribe(updateMarkets);
    }
  })
</script>

<div>
  <h1>Markets</h1>
  {#if markets}
    <ul>
    {#each [...markets.markets.entries()] as [id, tradableInstrument](id)}
      <li><a href='/trading/markets/{id}'>
       {tradableInstrument.tradableInstrument.instrument.id}
      </a></li>
    {/each}
    </ul>
    {:else}
     Loading...
    {/if}
</div>