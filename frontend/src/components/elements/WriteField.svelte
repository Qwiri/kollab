<script lang="ts">
	import CodeMirrorEditor from "./CodeMirrorEditor.svelte";

	import { mdText } from "../../store";
	import { onMount } from "svelte";

	import CodeMirror from "codemirror";

	// supported languages
	import "codemirror/mode/gfm/gfm";
	import "codemirror/mode/javascript/javascript";

	import "codemirror/keymap/vim";
	import { language, writeMode } from "../../editorPref";

	const options = {
		mode: "gfm",
		lineNumbers: true,
		value: $mdText,
		keyMap: "vim",
		lineWrapping: true,
	};
	let editor;

	$: if (editor) {
		writeMode.subscribe((m) => editor.setOption("keyMap", m));
		language.subscribe(switchLanguage);
	}

	function switchLanguage(lang: string) {
		if (!CodeMirror.modes[lang]) {
			console.error("Language", lang, "not loaded.");
			return;
		}
		console.log("Switching to language: " + lang);
		editor.setOption("mode", lang);
	}

	let cursor_activity = false;
	onMount(() => {
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

<CodeMirrorEditor
	on:activity={cursorMoved}
	on:change={changed}
	bind:editor
	{options}
/>
