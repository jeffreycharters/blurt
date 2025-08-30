import type { PageServerLoad } from "./$types"
import { z } from "zod"
import { message, superValidate } from "sveltekit-superforms"
import { zod } from "sveltekit-superforms/adapters"
import { fail, redirect } from "@sveltejs/kit"
import type { User } from "../app"
import { env } from "$env/dynamic/private"

const schema = z.object({
	username: z
		.string()
		.min(1, { message: "Username please!" })
		.max(14, { message: "Username too long! Max 14 characters" })
})

export const load = (async () => {
	const form = await superValidate(zod(schema))
	return { form }
}) satisfies PageServerLoad

export const actions = {
	default: async ({ request, cookies }) => {
		const form = await superValidate(request, zod(schema))
		if (!form.valid) {
			return fail(400, { form })
		}

		let res: Response

		try {
			res = await fetch(`${env.API_ADDRESS}/api/v1/users/${form.data.username}`, {
				method: "POST"
			})
		} catch (err) {
			return message(form, "Login failed! Try later.", { status: 500 })
		}

		if (!res.ok) {
			return message(form, "Unable to login. Try later.", { status: 500 })
		}

		const user: User = await res.json()
		cookies.set("userID", user.id, {
			path: "/",
			httpOnly: true,
			secure: false
		})

		return redirect(302, "/blurts")
	}
}
