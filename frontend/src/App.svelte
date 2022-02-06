<script>
	import EditorSettings from "./components/controls/EditorSettings.svelte";
	import Display from "./components/elements/Display.svelte";
	import WriteField from "./components/elements/WriteField.svelte";
	import { enableWYSIWYG } from "./editorPref";
</script>

<main>
	<div class="halfScreen">
		<div class="fullscreenY growPls">
			<EditorSettings />
			<WriteField />
		</div>
	</div>
	{#if $enableWYSIWYG}
		<div class="halfScreen fullscreenY displayWrapper">
			<Display />
		</div>
	{/if}
</main>

<style lang="scss">
	main {
		min-height: 100vh;
		width: 100%;
		display: flex;
		flex-wrap: wrap;
		justify-content: center;

		.fullscreenY {
			min-height: 100vh;
		}
		.halfScreen {
			padding: 0.5rem;
			overflow-wrap: break-word;
			flex-grow: 1;
			width: 40%;
			min-width: 420px;
		}

		.growPls {
			display: flex;
			flex-direction: column;
			position: relative;
		}
	}
</style>
