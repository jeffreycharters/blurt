<script lang="ts">
	import { getBlurtState } from "$lib/blurts.svelte"
	import { crossfade } from "svelte/transition"
	import { Blurt } from "."
	import { flip } from "svelte/animate"
	import { getUserState } from "$lib/users.svelte"
	import { ContentType, get_websocket, type Message } from "$lib/websocket.svelte"
	import { browser } from "$app/environment"

	const blurts = getBlurtState()
	const users = getUserState()

	const [send, receive] = crossfade({
		duration: (d) => Math.sqrt(d * 200)
	})

	let ws = browser ? get_websocket(blurts) : new WebSocket("ws://localhost:3320/ws")

	async function blurt_likker(id: string, error_func: () => void) {
		const lik_message: Message = {
			content_type: ContentType.Lik,
			payload: { likker: users.active_user, blurt_id: id }
		}

		ws.send(JSON.stringify(lik_message))
	}
</script>

<div class="min-w-xs md:min-w-md flex w-full max-w-screen-xl flex-col gap-4 px-1">
	{#each blurts.list as blurt (blurt.id)}
		<div
			in:receive={{ key: blurt.id }}
			out:send={{ key: blurt.id }}
			animate:flip={{ duration: 250 }}
		>
			<Blurt {blurt} {blurt_likker} />
		</div>
	{/each}
</div>
