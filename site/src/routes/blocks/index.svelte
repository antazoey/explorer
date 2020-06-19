<script>
  import { onMount } from 'svelte';
  import axios from 'axios'

  let blocks = [];

  onMount(async () => {
	const {data} = await axios.get(`https://geo.s.vega.xyz:8443/blockchain`)
    blocks = data.result.block_metas
  })
</script>

<style>
	ul {
		margin: 0 0 1em 0;
		line-height: 1.5;
	}
</style>

<svelte:head>
	<title>Blocks</title>
</svelte:head>

<h1>Recent blocks</h1>

<ul>
	{#each blocks as block}
		<li>
			<a href='blocks/{block.header.height}'>
				{block.header.height}
			</a>
			({block.num_txs})
		</li>
	{/each}
</ul>