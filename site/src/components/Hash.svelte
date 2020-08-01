<script>
    import Clipboard from 'svelte-clipboard';
    export let text;

    const textEnd = text.substr(-5);
    let copyElement;

    function copyReaction(){
      copyElement.text = '✅';
      setTimeout(() => {
          copyElement.text = '📋'
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
        min-width: 2.76em;
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

    a {
        display: none;
    }

    .hash-wrapper:hover a {
        display: inline-block;
        cursor: pointer;
    }

</style>

<div class="hash-wrapper">
    <Clipboard
        text={text}
        let:copy
        on:copy={copyReaction}>
    {#if text.length > 20}
        <div class="container">
            <span class="firstPart">{text}</span><span class="lastPart">{textEnd}</span>
        </div>
    {:else}
        {text}
    {/if}
    <a on:click={copy} on:copy={console.log} bind:this={copyElement}>📋</a>
</Clipboard>
</div>
