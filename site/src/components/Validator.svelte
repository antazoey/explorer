<script>
    import TwoColumnData from "./TwoColumnData.svelte";

    export let validator

    let rows = [
        {title: 'Address', value: validator.address, type: 'hash'},
        {title: 'Voting Power', value: validator.voting_power},
        {title: 'Priority', value: validator.proposer_priority}
    ]

</script>

<TwoColumnData rows={rows} />