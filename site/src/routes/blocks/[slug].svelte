<script>
	import Transaction from '../../components/Transaction.svelte'
    import { blockUrl, tendermintBaseUrl } from '../../config/'
	import { onMount, onDestroy } from 'svelte';
	import { stores } from "@sapper/app";
	import { store } from '../../stores/blocks'
	const { page } = stores();
	let { slug } = $page.params;
	const title = slug
	let data = [];
	let block = false
	let unsubscribe

	onMount(async () => {
		if(!unsubscribe) {
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
		if(unsubscribe) {
			unsubscribe();
			unsubscribe = null;
		}
	});

	function update(data) {
		block = data.blocks.get(slug)
        console.log(block)
	}

	function navigate(id) {
		console.log(id)
		slug = id
		update()
	}

	function getBlockFromTradeId(id) {
		return Number(id.split('-')[0].replace('V', 0)).toString()
	}

</script>
<style>
	.content :global(h2) {
		font-size: 1.4em;
		font-weight: 500;
	}

	.content :global(pre) {
		background-color: #f9f9f9;
		box-shadow: inset 1px 1px 5px rgba(0,0,0,0.05);
		padding: 0.5em;
		border-radius: 2px;
		overflow-x: auto;
	}

	.content :global(pre) :global(code) {
		background-color: transparent;
		padding: 0;
	}

	.content :global(ul) {
		line-height: 1.5;
	}

	.content :global(li) {
		margin: 0 0 0.5em 0;
	}
</style>

<svelte:head>
	<title>{title}!</title>
</svelte:head>

<h1>{title}</h1>

<ul class='content'>
	{JSON.stringify(block)}
{#each data as { Type, Sig, PubKey, Command }, i}
	<li>
		<b>{Type}</b> - <i>{Command.reference}</i>
		<Transaction tx={Command} pubKey={PubKey}/>
	</li>
{/each}
</ul>
