<script lang="ts">

	import CodeMirror from './CodeMirror.svelte'
	import 'codemirror/mode/gfm/gfm';
	import 'codemirror/keymap/vim';

    import { mdText } from "../../store";
	import {onMount} from 'svelte'
	const code = `function test() {\n  return 42\n}`
	const codee = `# Test\n- test\n- test`
	const options = {
		mode: "gfm",
		lineNumbers: true,
		value: $mdText,
        keyMap: 'vim'
	}
	let editor
	let cursor_activity = false
	onMount(()=>{
		console.log("Editor: ", editor);
	})
	
	function cursorMoved(event) {
		cursor_activity = true;
		console.log('cursor activity');
		
	}
	
	function changed(event) {
		console.log('changed');
        $mdText = editor.doc.getValue();
	}
	
</script>

<CodeMirror on:activity={cursorMoved} on:change={changed} bind:editor {options} />