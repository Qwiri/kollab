<script>
	import Display from "./components/elements/Display.svelte";
	import WriteField from "./components/elements/WriteField.svelte";
	import { writeMode } from "./store";

	let wysiwyg = true;
</script>

<main>
	<div class="halfScreen">
		<div class="fullscreenY growPls">
			<WriteField />
		</div>
		<div id="settings">
			<label>
				WYSIWYG?
				<input type="checkbox" name="wysiwyg" bind:checked={wysiwyg} />
			</label>
			<label>
				Mode?
				<select bind:value={$writeMode}>
					<option>default</option>
					<option>vim</option>
					<option>emacs</option>
				</select>
			</label>
		</div>
	</div>
	{#if wysiwyg}
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
			height: 100%;
		}
		.halfScreen {
			padding: 0.5rem;
			overflow-wrap: break-word;
			flex-grow: 1;
			width: 40%;
		}

		.growPls {
			display: flex;
			flex-direction: column;
		}

		select {
			color: white;
		}
	}
</style>
