<script>
    export let text;
    const textEnd = text.substr(-5)
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
        white-space: nowrap;
        overflow: hidden;
    }
</style>

<div class="container">
    <span class="firstPart">{text}</span><span class="lastPart">{textEnd}</span>
</div>