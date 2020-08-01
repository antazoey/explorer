<script>
    import Hash from "./Hash.svelte";

    export let id;
    import {store} from "../stores/bots.js";
    import {onDestroy, onMount} from 'svelte'

    let unsubscribe

    onDestroy(() => {
        if (unsubscribe) {
            unsubscribe();
            unsubscribe = null;
        }
    })

    function updateMarkets(data) {
    }

    onMount(() => {
        if (!unsubscribe) {
            unsubscribe = () => {
            } //store.subscribe();
        }
    })

    let icon = ''
    if (id === 'network') {
        icon = '🔴'
    }

</script>

<div>
<a href='/party/{id}'>
    {#if id.length > 15}
        <Hash text={id} />
    {:else}
        {icon}
        {id}
    {/if}
</a>
</div>