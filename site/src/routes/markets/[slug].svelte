<script>
  import Market from "../../components/Market.svelte";
  import { Markets } from "../../stores/markets.js";
  import { stores } from "@sapper/app";
  const { page } = stores();
  const { slug } = $page.params; 
	
  let market = $Markets.get($page.params.slug)
  
</script>

<div>
  {#if market}
      <Market id={$page.params.slug} /> 
  {:else}
      <h1>Loading...</h1>
  {/if}
</div>
