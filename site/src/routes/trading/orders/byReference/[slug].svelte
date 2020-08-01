<script>
    import {onMount, onDestroy} from 'svelte';
    import {stores, goto} from "@sapper/app";
    import {Orders} from '../../../../stores/orders'

    const ordersStore = new Orders();

    const {page} = stores();
    let {slug} = $page.params;
    const title = slug

    onMount(async () => {
        const order = await ordersStore.getByReference(slug)
        goto(`/trading/orders/${order.id}`);
    })


</script>

<svelte:head>
    <title>{title}!</title>
</svelte:head>

<p>Redirecting to order id...</p>