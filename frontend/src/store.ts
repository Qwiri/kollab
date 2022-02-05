import { writable } from 'svelte/store';
import type { Writable } from 'svelte/store';

export const mdText: Writable<string> = writable('');
export const writeMode: Writable<string> = writable("vim");
export const language: Writable<string> = writable("gfm");