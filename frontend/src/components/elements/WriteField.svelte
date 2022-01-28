<script lang="ts">
	import CodeMirror from "./CodeMirror.svelte";
	import "codemirror/mode/gfm/gfm";
	import "codemirror/keymap/vim";

    import { mdText, writeMode } from "../../store";
	import { onMount } from 'svelte';
	

	const options = {
		mode: "gfm",
		lineNumbers: true,
		value: $mdText,
        keyMap: 'vim',
		lineWrapping: true
	}
	let editor;
	$: if (editor) {
		writeMode.subscribe(m => editor.setOption("keyMap", m));
	}
	let cursor_activity = false
	onMount(()=>{
		console.log("Editor: ", editor);
	});

	function cursorMoved(event) {
		cursor_activity = true;
		console.log("cursor activity");
	}

	function changed(event) {
		console.log("changed");
		$mdText = editor.doc.getValue();
	}
</script>

<CodeMirror
	on:activity={cursorMoved}
	on:change={changed}
	bind:editor
	{options}
/>
