<script>
  import { Parties } from "../../stores/parties.js";
  import { Leaderboard } from "../../stores/leaderboard.js";
</script>

<div>
  <h1>Trader Leaderboard</h1>
  <ol>
  {#each $Leaderboard as Party, id}
    <li><a href='party/{Party.publicKey}'>
     {Party.publicKey}
    </a></li>
  {/each}
  </ol>
</div>
