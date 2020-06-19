<script>
  import { Markets } from "../stores/markets.js";
  export let id = '';
</script>

<div>
  <a href='/markets/{id}'>{$Markets.get(id) ? $Markets.get(id).name : id}</a>
</div>