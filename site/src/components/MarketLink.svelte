<script>
  import { Markets } from "../stores/markets.js";
  export let id;
  let market = $Markets.get(id)
</script>

<div>
  <a href='/markets/{id}'>{market ? market.name : id}</a>
</div>