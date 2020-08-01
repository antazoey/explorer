<script context="module">
	import {blockUrl, tendermintBaseUrl} from '../../config/'

	export async function preload(page) {
		const { slug } = page.params;
		let res = await this.fetch(blockUrl(), {
			method: 'POST',
			mode: 'cors',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				block_height: parseInt(slug, 10),
				node_url: tendermintBaseUrl
			})
		})

		let json = await res.json()

		let data = json.map(d => {
			d.Command = JSON.parse(d.Command)
			return d
		})

		return { data, slug }
	}
</script>
<script>
	import Transaction from '../../components/Transaction.svelte'
	import {onDestroy} from 'svelte';
	import {store} from '../../stores/blocks'
	import BlockHeader from "../../components/BlockHeader.svelte";

	export let slug
	export let data = [];
	export let block = [];
	let unsubscribe

	if (!unsubscribe) {
		unsubscribe = store.subscribe(update);
	}

	onDestroy(() => {
		if (unsubscribe) {
			unsubscribe();
			unsubscribe = null;
		}
	});

	function update(data) {
		block = data.blocks.get(slug)
        if (!block && data.fetchBlock) {
       		data.fetchBlock(slug)
		}
	}

	function getBlockFromTradeId(id) {
		return Number(id.split('-')[0].replace('V', 0)).toString()
	}
</script>

<svelte:head>
	<title>Block {slug}</title>
</svelte:head>

<h1>Block {slug}</h1>

<ul class='content'>
	{#if block}
		<BlockHeader block={block.block} />
		<hr>
	{/if}
	{#each data as { Type, PubKey, Command }, i}
		<Transaction tx={Command} pubKey={PubKey} type={Type} />
	{/each}
</ul>
