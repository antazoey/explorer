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
  <a href='/markets/{id}'>{market ? market.tradableInstrument.instrument.id : id}</a>
</div>