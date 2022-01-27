<script>
    import TwoColumnData from "./TwoColumnData.svelte";

    export let validator

    console.log(validator)
    let rows = [
        {title: 'Address', value: validator.address, type: 'hash'},
        {title: 'Voting Power', value: validator.voting_power},
        {title: 'Priority', value: validator.proposer_priority},
        {title: 'Tendermint Key', value: validator.pub_key.value}
    ]

</script>

<TwoColumnData rows={rows} />