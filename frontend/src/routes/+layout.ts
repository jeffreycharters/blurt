import { browser } from '$app/environment';
import type { Blurt } from '$lib/blurts.svelte';

import { API_ADDRESS, BLURT_BATCH_COUNT } from '$lib/index';

export const prerender = true;
export const trailingSlash = 'always'

export async function load({ fetch }) {
	let user = ""
	if (browser) {
		const user_list = JSON.parse(localStorage.getItem('user_list') ?? "[]")
		user = user_list.length > 0 ? user_list[0] : ""
	}

	let blurts: Blurt[] = []
	const res = await fetch(API_ADDRESS + '/blurts?offset=0&count=' + BLURT_BATCH_COUNT)

	if (res.ok) blurts = await res.json()
	if (!res.ok) console.error(res.status, res.statusText)

	return {
		blurts,
		user
	};
}
