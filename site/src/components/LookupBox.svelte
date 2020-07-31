<script>
    let message = 'Experimental search. Enter an ID above.';
    let link = ''

    function detect(e) {
        const i = e.target.value.toLowerCase()

        if (i.substr(0, 2) === '0x') {
            message = 'This looks like a party id'
            link = `/party/${i}`
        } else if (i.indexOf('-') !== -1) {
            const split = i.split('-')
            const block = Number(split[0].substr(1)).toString()
            if (split.length === 2) {
                // Order
                message = 'This looks like an order'
                link = `/blocks/${block}`
            } else {
                // Trade
                message = 'This looks like a trade'
                link = `/blocks/${block}`
            }
        } else {
            // Could be a market
            message = 'Could be a market?'
        }
    }
</script>

<form>
    <input type="text" on:change={e => detect(e)} />
</form>

<p>
    {message}
    {#if link}
        <a href="{link}">Go!</a>
    {/if}
</p>
