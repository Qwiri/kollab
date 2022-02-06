import { writable } from 'svelte/store';
import type { Writable } from 'svelte/store';

export const mdText: Writable<string> = writable('');

export const languages = [
    "gfm",
    "qml",
    "javascript",
];