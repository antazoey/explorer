<script>
	import { store } from "../../stores/blocks.js";

	import { onDestroy, onMount } from 'svelte'

	let unsubscribe
	let blocks

	onDestroy(() => {
		if(unsubscribe) {
			unsubscribe();
			unsubscribe = null;
		}
	});

	function updateData(data) {
		blocks = data.blocks;
	}

	onMount (() => {
		if(!unsubscribe) {
			unsubscribe = store.subscribe(updateData);
		}

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
{#if blocks}
<ul>
	{#each [...blocks].sort().reverse() as [id, {block}](id)}
		<li>
			<a href='blocks/{block.header.height}'>
				{block.header.height}
			</a>
			(txs: {block.num_txs}, size: {block.block_size}, hash: {block.block_id.hash})
		</li>
	{/each}
</ul>
{/if}