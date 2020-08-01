<script>
    import Transaction from '../../../components/Transaction.svelte'
    import {blockUrl, tendermintBaseUrl} from '../../../config/'
    import {onMount, onDestroy} from 'svelte';
    import {stores} from "@sapper/app";
    import {store} from '../../../stores/blocks'
    import BlockHeader from "../../../components/BlockHeader.svelte";
    import {getTx} from '../../../stores/txs'
    import {Orders} from '../../../stores/orders'
    import OrderDetails from "../../../components/OrderDetails.svelte";

    const {page} = stores();
    let {slug} = $page.params;
    const title = slug
    let data = [];
    let block = false
    let order = false
    let unsubscribe
    let ordersStore = new Orders()

    onMount(async () => {
        if (!unsubscribe) {
            unsubscribe = store.subscribe(update);
        }

        order = await ordersStore.get(slug)

//        data = await getTx()
    })

    onDestroy(() => {
        if (unsubscribe) {
            unsubscribe();
            unsubscribe = null;
        }
    });

    function update(data) {
//        block = data.blocks.get(slug)
//        if (!block && data.fetchBlock) {
//           data.fetchBlock(slug)
//        }
    }

    function navigate(id) {
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
    <OrderDetails order={order} />
</ul>
