<script>
    import TwoColumnData from "../components/TwoColumnData.svelte";
    export let account;
    export let marketId = false;

    if (!marketId && account.market) {
        marketId = account.market.id;
    }

    const rows = [
        { title: 'Account Type', value: account.type },
        { title: 'Balance', value: account.balance, type: 'price', marketId: marketId },
        { title: 'Asset', value: account.asset },
    ];

    if (account.market) {
        rows.push({ title: 'Market', value: marketId, type: 'market' })
    }

</script>

<TwoColumnData rows={rows} />
