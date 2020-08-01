<script>
    export let id

    function getBlockFromTradeId(id) {
        return Number(id.split('-')[0].replace('V', 0)).toString()
    }

    const block = getBlockFromTradeId(id)
</script>

🟫
<a href='/blocks/{block}'>
     {block}
</a>