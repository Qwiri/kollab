<script>
	import Display from "./components/elements/Display.svelte";
	import WriteField from "./components/elements/WriteField.svelte";
	import { language, writeMode } from "./store";

	let wysiwyg = true;
</script>

<main>
	<div class="halfScreen">
		<div class="fullscreenY growPls">
			<WriteField />
		</div>
		<div id="settings">
			<h1>Settings</h1>
			<table>
				<tr>
					<td>
						<label for="wysiwyg">WYSIWYG?</label>
					</td>
					<td>
						<input type="checkbox" name="wysiwyg" bind:checked={wysiwyg} />
					</td>
				</tr>
				<tr>
					<td>
						<label for="writeMode">
							Mode?
						</label>
					</td>
					<td>
						<select name="writeMode" bind:value={$writeMode}>
							<option>default</option>
							<option>vim</option>
							<option>emacs</option>
						</select>
					</td>

				</tr>
				<tr>
					<td>
						<label for="lang">Lang</label>
					</td>
					<td>
						<input type="text" name="lang" bind:value={$language} />
					</td>
				</tr>
			</table>
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
		}

		select {
			color: white;
		}
		#settings {

			h1 {
				margin-top: 0;
			}
			border-radius: .2rem;
			border: 2px solid var(--main-grey);
			padding: .5rem;
			padding-left: 2rem;

			table {
				text-align: center;
			}
		}
	}
</style>
