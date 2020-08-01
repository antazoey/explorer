<script context="module">
    import {Orders} from '../../../stores/orders'

    export async function preload(page) {
        let ordersStore = new Orders()
        const { slug } = page.params;

        const order = await ordersStore.get(slug, this.fetch)
        return { order, slug }
    }
</script>

<script>
    import {onMount, onDestroy} from 'svelte';
    import {store} from '../../../stores/blocks'
    import BlockHeader from "../../../components/BlockHeader.svelte";
    import OrderDetails from "../../../components/OrderDetails.svelte";
    import TradeDetails from "../../../components/TradeDetails.svelte";

    export let slug

    const title = slug
    let data = [];
    let block = false
    export let order = false
    let unsubscribe

    onMount(async () => {
        if (!unsubscribe) {
            unsubscribe = store.subscribe(update);
        }
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

    function navigate(id) {
        slug = id
        update()
    }

</script>
<style>
    .trades-container {
        margin: 0 30px;
    }
</style>

<svelte:head>
    <title>{title}!</title>
</svelte:head>

<h1>{title}</h1>

<ul class='content'>
    {#if block}
        <BlockHeader block={block.block} />
        <hr>
    {/if}
    <OrderDetails order={order} />
    {#if order.trades}
        <h2>Trades</h2>
        <div class="trades-container">
            {#each order.trades as trade}
                 <TradeDetails trade={trade} /><br>
            {/each}
        </div>
    {/if}
</ul>
