<script>
    import {store as networkStore} from "../../stores/network.js";
    import {store as blockStore} from "../../stores/blocks.js";

    import {onDestroy, onMount, afterUpdate } from 'svelte'
    import Validator from "../../components/Validator.svelte";
    import BlockSummary from "../../components/BlockSummary.svelte";

    let unsubscribeNetwork
    let unsubscribeBlocks
    let validators
    let blocks

    onDestroy(() => {
        if (unsubscribeNetwork) {
            unsubscribeNetwork();
            unsubscribeNetwork = null;
        }
        if (unsubscribeBlocks) {
            unsubscribeBlocks();
            unsubscribeBlocks = null;
        }
    });

    function updateData(data) {
        if (data.peers) {
            $: validators = [...data.peers.values()];
        }
        if (data.blocks) {
            $: blocks = [...data.blocks.values()].slice(0, 5)
        }
    }

    afterUpdate(() => {
    })

    onMount(() => {
        if (!unsubscribeNetwork) {
            unsubscribeNetwork = networkStore.subscribe(updateData);
        }
        if (!unsubscribeBlocks) {
            unsubscribeBlocks = blockStore.subscribe(updateData);
        }
    })
</script>

<svelte:head>
    <title>Network</title>
</svelte:head>
{#if validators && validators.length > 0}
    <h1>Validators</h1>
    {#each validators as validator}
        <Validator validator={validator} />
    {/each}
{/if}
{#if blocks && blocks.length > 0}
    <h1>Recent blocks</h1>
    {#each blocks as {block}}
        <BlockSummary block={block} />
    {/each}
{/if}
