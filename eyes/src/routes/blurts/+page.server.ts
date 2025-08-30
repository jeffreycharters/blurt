import { fail, redirect } from "@sveltejs/kit"
import type { Blurt } from "../../app"
import type { PageServerLoad } from "./$types"
import { env } from "$env/dynamic/private"

export const load = (async ({ cookies }) => {
	const userID = cookies.get("userID")

	if (!userID) return redirect(301, "/")

	let res: Response

	try {
		res = await fetch(`${env.API_ADDRESS}/api/v1/blurts`)
	} catch (err) {
		return fail(500, { error: "Could not fetch blurts" })
	}

	const blurts = (await res.json()) as Blurt[]
	// console.log(blurts)

	return {
		blurts,
		userID
	}
}) satisfies PageServerLoad
