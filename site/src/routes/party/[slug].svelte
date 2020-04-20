<script>
	import axios from 'axios'
	import { onMount } from 'svelte';
  import { stores } from "@sapper/app";
  import Market from '../../components/Market.svelte'
  const { page } = stores();
  const { slug } = $page.params; 

  let party = {}
  let accounts = []
  onMount(async () => {
    const res = await axios.get(`https://geo.s.vega.xyz/parties/${slug}/accounts`)
    accounts = res.data.accounts
	})	
</script>

<div>
  {#if accounts}
    <ol>
    {#each accounts as { asset, balance, marketID, type}, i}
      {#if !marketID}
      <li>{asset}: {balance}</li>
      {:else}
      <li><Market id={marketID} />: {asset}: {balance}</li>
      {/if}
    {/each}
    </ol>
  {:else}
      <h1>Loading...</h1>
  {/if}
</div>
