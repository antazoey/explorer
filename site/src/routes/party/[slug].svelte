<script>
  import { stores } from "@sapper/app";
  import { onMount, tick } from 'svelte'

  import MarketLink from "../../components/MarketLink.svelte";
  import { Parties } from "../../stores/parties.js";

  const { page } = stores();
  let { slug } = $page.params;

  let store = new Parties()
  let party

  let update = async () => {
    party = await store.get(slug)
    console.log(party)
  }

  onMount(update)

  function navigate(id) {
    console.log(id)
    slug = id
    update()
  }

  function getBlockFromTradeId(id) {
    return Number(id.split('-')[0].replace('V', 0)).toString()
  }
</script>

<div>
  {#if party}
    <h1>{party.id}</h1>
    <h2>Accounts</h2>
    <ul>
      {#each party.accounts as acc}
      <li>
        {#if acc.type === 'General'}
        <b>{acc.type} {acc.asset} account: {acc.balance}</b>
        {:else}
        <b>{acc.type} {acc.asset} for <MarketLink id={acc.market.id} />: {acc.balance}</b>
        {/if}
        </li>
      {/each}
    </ul>
    <h2>Recent orders</h2>
    {#each party.orders as or}
        {#if or.remaining === 0}
          &nbsp;
          &nbsp;&nbsp;<s>{or.side} in <MarketLink id={or.market.id} /></s>: {or.size} @ {or.price}
          in <a href="{`/blocks/${getBlockFromTradeId(or.id)}`}">{getBlockFromTradeId(or.id)}</a>
        {:else if (or.remaining === or.size)}
          &nbsp;
          &nbsp;&nbsp;<b>{or.side} in <MarketLink id={or.market.id} /></b>: {or.size} @ {or.price}
          in <a href="{`/blocks/${getBlockFromTradeId(or.id)}`}">{getBlockFromTradeId(or.id)}</a>
        {:else}
          <details>
          <summary>
          {#if or.remaining !== '0'}
            <b>{or.side} in <MarketLink id={or.market.id} /></b>: {or.size} @ {or.price} <i>(remaining: {or.remaining})</i>
            in <a href="{`/blocks/${getBlockFromTradeId(or.id)}`}">{getBlockFromTradeId(or.id)}</a>
          {:else}
            <s>{or.side} in <MarketLink id={or.market.id} /></s>: {or.size} @ {or.price} (filled)
            in <a href="{`/blocks/${getBlockFromTradeId(or.id)}`}">{getBlockFromTradeId(or.id)}</a>
          {/if}
          </summary>
          {#if or.trades && or.trades.length !== 0}
            <div style="margin-left: 60px">
            <h3>Trades</h3>
            {#each or.trades as t}
              <li>
                {#if or.side === 'Buy'}
                <b>{t.size} -> <a on:click="{() => {navigate(t.seller.id)}}" href='/party/{t.seller.id}'>{t.seller.id}</a></b>
                {:else}
                <b>{t.size} -> <a on:click="{() => {navigate(t.buyer.id)}}" href='/party/{t.buyer.id}'>{t.buyer.id}</a></b>
                {/if}
                in <a href="{`/blocks/${getBlockFromTradeId(t.id)}`}">{getBlockFromTradeId(t.id)}</a>
              </li>
            {/each}
            </div>
            {/if}
          </details>
        {/if}
    {/each}
  {:else}
    <h1>Loading...</h1>
  {/if}
</div>
