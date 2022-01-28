<script lang="ts">
    import SvelteMarkdown from "svelte-markdown";
    import { mdText } from "../../store";
</script>

<SvelteMarkdown source={$mdText} />

<style lang="scss">
    // markdown styling
    :global(h1, h2, h3, h4, h5, h6) {
        color: var(--main-color);
    }
    :global(ul) {
        list-style: none;

        :global(li::before) {
            content: "\2022";
            color: var(--main-color);
            font-weight: bold;
            display: inline-block;
            width: 1em;
            margin-left: -1em;
        }
    }

    :global(ol) {
        list-style: none;
        counter-reset: li;

        :global(li::before) {
            color: var(--main-color);
            content: counter(li)".";
            display: inline-block;
            width: 1em;
            margin-left: -2em;
            counter-increment: li;
        }
    }

    :global(hr) {
        border: 1px solid var(--main-grey);
        background-color: var(--main-grey);
    }
    :global(table, th, td) {
        border: 2px solid var(--main-grey);
        border-collapse: collapse;
    }
    :global(th, td) {
        padding: .5rem;
    }
</style>
