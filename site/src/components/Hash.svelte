<script>
    import Clipboard from 'svelte-clipboard';
    export let text;

    let copyElement;

    function copyReaction(){
      copyElement.innerHTML = '✅';
      setTimeout(() => {
          if (copyElement) {
              copyElement.innerHTML = '📋'
          }
      }, 3000)
    }
</script>
<style>
    .firstPart, .lastPart {
        display: inline-block;
        vertical-align: bottom;
        white-space: nowrap;
        overflow: hidden;
    }

    .firstPart {
        max-width: calc(100% - 3.68em);
        min-width: 1.76em;
        text-overflow: ellipsis;
    }

    .lastPart {
        max-width: calc(100% - 2.76em);
        direction: rtl;
    }

    .container {
        display: inline;
        white-space: nowrap;
        overflow: hidden;
    }

    b {
        display: none;
    }

    .hash-wrapper:hover b {
        display: inline-block;
        cursor: pointer;
    }

</style>

{#if text}
<div class="hash-wrapper">
    <Clipboard
        text={text}
        let:copy
        on:copy={copyReaction}>
    {#if text.length > 20}
        <div class="container">
            <span class="firstPart">{text}</span><span class="lastPart">{text.substr(-5)}</span>
        </div>
    {:else}
        {text}
    {/if}
    <b on:click={copy} bind:this={copyElement}>📋</b>
</Clipboard>
</div>
{/if}
