<script>
  import { Markets } from "../../stores/markets.js";
</script>

<div>
  <h1>Markets</h1>
  <ul>
  {#each [...$Markets.entries()] as Market, id}
    <li><a href='markets/{Market[1].id}'>
     {Market[1].name}
    </a></li>
  {/each}
  </ul>
</div>