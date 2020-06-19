<script>
  import { Markets } from "../../stores/markets.js";
  import { stores } from "@sapper/app";
  const { page } = stores();
  const { slug } = $page.params; 
	
  let market = $Markets.get($page.params.slug)
  
</script>

<div>
  {#if market}
      <h1>{market.name}</h1>
  {:else}
      <h1>Loading...</h1>
  {/if}
</div>
