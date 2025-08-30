import { browser } from "$app/environment"
import type { PageLoad } from "./$types"

export const load = (async () => {
	if (browser) localStorage.setItem("username", "test")

	let username = ""
	if (browser) username = localStorage.getItem("username") ?? ""

	console.log(username.length > 0 ? username : "hi")

	return {
		username
	}
}) satisfies PageLoad
