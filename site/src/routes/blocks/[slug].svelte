<script>
	import Transaction from '../../components/Transaction.svelte'
    import { blockUrl, tendermintBaseUrl } from '../../config/'
	import { onMount } from 'svelte';
	import { stores } from "@sapper/app";

	const { page } = stores();
	const { slug } = $page.params; 
	const title = slug
	let data = [];
  
	onMount(async () => {
		let res = await fetch(blockUrl(), {
			method: 'POST',
			mode: 'cors',
			headers: {
				'Content-Type': 'application/json'
			},
			body: {
	 	       block_height: parseInt(slug, 10),
				node_url: tendermintBaseUrl
			}
		})

		const data = await res.json()

		data = txs.data.map(d => {
			d.Command = JSON.parse(d.Command)
			return d
		})
	})
</script>
<style>
	/*
		By default, CSS is locally scoped to the component,
		and any unused styles are dead-code-eliminated.
		In this page, Svelte can't know which elements are
		going to appear inside the {{{post.html}}} block,
		so we have to use the :global(...) modifier to target
		all elements inside .content
	*/
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
{#each data as { Type, Sig, PubKey, Command }, i}
	<li>
		<b>{Type}</b> - <i>{Command.reference}</i>
		<Transaction tx={Command} pubKey={PubKey}/>
	</li>
{/each}
</ul>
