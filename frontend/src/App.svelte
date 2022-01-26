<script>
	import Display from "./components/elements/Display.svelte";
	import WriteField from "./components/elements/WriteField.svelte";

</script>
<main>
	<div class="halfScreen growPls">
		<WriteField />
	</div>
	<div class="halfScreen">
		<Display />
	</div>
</main>
<style lang="scss">
	main {
		min-height: 100vh;
		display: flex;

		.halfScreen {
			min-height: 100%;
			width: 50%;
			padding: .5rem;
		}

		.growPls {
			display: flex;
			flex-direction: column;
		}
	}

</style>