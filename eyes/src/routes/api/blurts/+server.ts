import { env } from "$env/dynamic/private"
import type { RequestHandler } from "./$types"

export const GET: RequestHandler = async ({ url }) => {
	const offset = Number(url.searchParams.get("offset")) ?? 0
	const count = Number(url.searchParams.get("count")) ?? 25

	const res = await fetch(`${env.API_ADDRESS}/api/v1/blurts?offset=${offset}&count=${count}`)

	if (!res.ok) return new Response(null, { status: 500 })

	return new Response(JSON.stringify(await res.json()))
}
