<script>
    import { enableWYSIWYG, language, writeMode } from "../../editorPref";
    import { languages } from "../../store";
</script>

<div id="settings">
    <input type="checkbox" name="wysiwyg" bind:checked={$enableWYSIWYG} />
    <select name="writeMode" bind:value={$writeMode}>
        <option>default</option>
        <option>vim</option>
        <option>emacs</option>
    </select>

    <select name="language" bind:value={$language}>
        {#each languages as lang}
            <option>{lang}</option>
        {/each}
    </select>
</div>

<style lang="scss">
    #settings {
        position: absolute;
        bottom: 2rem;
        z-index: 100;

        background-color: #000000;
        border-radius: 1000px;

        padding: 1.5rem;
        left: 0;
        right: 0;
        margin-left: auto;
        margin-right: auto;
        width: 25vw;

        select {
            color: white;
            border: none;
        }
    }
</style>
