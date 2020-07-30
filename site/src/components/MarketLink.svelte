<script>
  export let id;
  let market = {}
</script>

<div>
  <a href='/markets/{id}'>{market ? market.name : id}</a>
</div>