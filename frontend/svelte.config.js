import adapter from "@sveltejs/adapter-static"
import { sveltePreprocess } from "svelte-preprocess"

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://svelte.dev/docs/kit/integrations
	// for more information about preprocessors
	preprocess: sveltePreprocess(),
	kit: {
		adapter: adapter(),
		prerender: {
			handleHttpError: "warn"
		}
	}
}

export default config
