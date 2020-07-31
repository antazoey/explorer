<script>
	import Transaction from '../../components/Transaction.svelte'
	import {blockUrl, tendermintBaseUrl} from '../../config/'
	import {onMount, onDestroy} from 'svelte';
	import {stores} from "@sapper/app";
	import {store} from '../../stores/blocks'
	import BlockHeader from "../../components/BlockHeader.svelte";

	const {page} = stores();
	let {slug} = $page.params;
	const title = slug
	let data = [];
	let block = false
	let unsubscribe

	onMount(async () => {
		if (!unsubscribe) {
			unsubscribe = store.subscribe(update);
		}
		let res = await fetch(blockUrl(), {
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

		data = json.map(d => {
			d.Command = JSON.parse(d.Command)
			return d
		})
	})

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

	export let location;

	function navigate(id) {
		console.log(id)
		slug = id
		update()
	}

	function getBlockFromTradeId(id) {
		return Number(id.split('-')[0].replace('V', 0)).toString()
	}

</script>

<svelte:head>
	<title>{title}!</title>
</svelte:head>

<h1>{title}</h1>

<ul class='content'>
	{#if block}
		<BlockHeader block={block.block} />
		<hr>
	{/if}
	{#each data as { Type, Sig, PubKey, Command }, i}
		<Transaction tx={Command} pubKey={PubKey} sig={Sig} type={Type} />
	{/each}
</ul>
