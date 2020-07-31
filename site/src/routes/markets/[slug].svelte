<script>
  import { stores } from "@sapper/app";
  import { onDestroy, onMount } from 'svelte'

  import { store } from "../../stores/markets.js";
  import Market from "../../components/Market.svelte";

  const { page } = stores();
  const { slug } = $page.params;

  let unsubscribe
  let market

  onDestroy(() => {
      if(unsubscribe) {
          unsubscribe();
          unsubscribe = null;
      }
  });

  function updateMarkets(data) {
      market = data.markets.get(slug);
  }

  onMount (() => {
      if(!unsubscribe) {
          unsubscribe = store.subscribe(updateMarkets);
      }
  })

</script>

<div>
  {#if market}
      <Market id={market.id} />
  {:else}
      <h1>Loading...</h1>
  {/if}
</div>
