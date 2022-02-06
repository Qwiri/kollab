import { Writable, writable } from "svelte/store";
import { languages } from "./store";
import type { WriteMode } from "./types";

export const enableWYSIWYG: Writable<boolean> = writable(true);
export const writeMode: Writable<WriteMode> = writable('vim');
export const language: Writable<string> = writable(languages[0]);