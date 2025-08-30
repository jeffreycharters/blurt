<script lang="ts">
	import { browser } from "$app/environment"
	import { getBlurtState } from "$lib/blurts.svelte"
	import { getChaosState } from "$lib/chaos.svelte"
	import { getUserState } from "$lib/users.svelte"
	import { get_websocket } from "$lib/websocket.svelte"

	let blurt = $state("")
	const users = getUserState()
	const chaos = getChaosState()
	const blurts = getBlurtState()

	enum ContentType {
		Blurt = "blurt",
		Lik = "lik"
	}

	type Message = {
		content_type: ContentType
		payload: LikPayload | BlurtPayload
	}

	type LikPayload = {
		likker: string
		blurt_id: string
	}

	type BlurtPayload = {
		username: string
		content: string
	}

	let ws: WebSocket = browser ? get_websocket(blurts) : new WebSocket("ws://localhost:3320/ws")

	async function sendBlurt(event: Event) {
		event.preventDefault()

		const new_blurt: Message = {
			content_type: ContentType.Blurt,
			payload: { content: blurt, username: users.active_user }
		}

		ws.send(JSON.stringify(new_blurt))

		blurt = ""
		if (chaos.active) users.set_random_active()
	}
</script>

<form
	class="z-100 max-w-screen sticky top-2 mx-auto flex w-full flex-col items-baseline justify-between gap-1 rounded-md bg-teal-700 p-2 shadow-md md:max-w-md"
	onsubmit={sendBlurt}
>
	<div class="flex w-full items-baseline justify-between px-4"></div>
	<div class="flex w-full justify-between gap-2">
		<label for="blurt" class="sr-only">blurt</label>
		<div class="flex flex-grow items-baseline gap-2">
			<input
				class=" w-full rounded-md bg-teal-50 px-2 py-1 font-semibold text-teal-800"
				type="text"
				id="blurt"
				bind:value={blurt}
				placeholder={"blurting as " + users.active_user}
				autocomplete="off"
				maxlength="14"
			/>
		</div>

		<button
			type="submit"
			id="blurt-button"
			class="rounded-md bg-teal-50 px-4 py-1 font-semibold text-teal-800">blurt it</button
		>
	</div>
	<div class="pl-4 text-xs text-teal-50">blurt ammos left: {14 - blurt.length}</div>
</form>
