<script>
  import {store} from "../../stores/leaderboard.js";

  import {onDestroy, onMount} from 'svelte'
  import PartyLink from "../../components/PartyLink.svelte";

  let unsubscribe
  let leaderboard

  onDestroy(() => {
    if (unsubscribe) {
      unsubscribe();
      unsubscribe = null;
    }
  });

  function updateLeaderboard(data) {
    leaderboard = data.leaderboard;
  }

  onMount(() => {
    if (!unsubscribe) {
      unsubscribe = store.subscribe(updateLeaderboard);
    }
  })

</script>

<div>
  <h1>Trader Leaderboard</h1>
  {#if leaderboard}
    <table>
      <thead>
      <tr>
        <th>Public Key</th>
        <th>Total Balance</th>
      </tr>
    </thead>
    {#each [...leaderboard] as [ id, { trader } ]}
      <tr>
        <td>
          <PartyLink id={trader.publicKey} />
        </td>
        <td>
          ${trader.totalUsdVal}
        </td>
      </tr>
    {/each}
    </table>
  {:else}
      This will work when testnet ports are open...
  {/if}
</div>
