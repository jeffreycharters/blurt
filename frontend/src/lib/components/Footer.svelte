<script lang="ts">
	import { browser } from "$app/environment"
	import { API_ADDRESS } from "$lib"
	import { getBlurtState } from "$lib/blurts.svelte"
	import { getUserState } from "$lib/users.svelte"
	import { untrack } from "svelte"

	const blurts = getBlurtState()
	const users = getUserState()

	let footer: HTMLDivElement
	let loading_blurts = false

	$effect(() => {
		if (!browser) return

		const handleIntersect: IntersectionObserverCallback = async (entries, observer) => {
			entries.forEach((entry) => {
				if (!entry.isIntersecting || loading_blurts) return

				loading_blurts = true

				fetch(API_ADDRESS + `/blurts/?offset=${blurts.list.length}&username=${users.active_user}`)
					.then((res) => res.json())
					.then((new_blurts) => {
						if (new_blurts.length == 0) {
							observer.unobserve(footer)
							loading_blurts = false
							return
						}

						blurts.add_bulk(new_blurts)
						loading_blurts = false
					})
			})
		}

		const options = { threshold: 0, rootMargin: "200px 0px 0px" }
		new IntersectionObserver(handleIntersect, options).observe(footer)

		untrack(() => blurts)
	})
</script>

<div class="mb-6 rounded-md bg-white p-4 text-center font-bold md:mb-0">
	<div class="text-teal-500">THAT'S ALL OF THEM!!!</div>
	<div class="text-teal-600">You win in a losing sort of way!</div>
</div>

<div bind:this={footer}></div>
