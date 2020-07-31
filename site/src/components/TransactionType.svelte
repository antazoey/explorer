<script>
    export let type;
    let emoji = ''

    if (type === 'SubmitOrder') {
        emoji = '🛒'
        type = 'Order'
    }
</script>

<span>
  {#if emoji}
      <strong>{emoji}</strong>
  {/if}
  <span>{type}</span>
</span>