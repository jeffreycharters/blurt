<script lang="ts">
	import { browser } from "$app/environment"
	import { API_ADDRESS } from "$lib"
	import { getBlurtState } from "$lib/blurts.svelte"
	import { BlurtList, NewBlurtForm, PersonaManager } from "$lib/components"
	import { getUserState } from "$lib/users.svelte"
	import { onMount } from "svelte"

	const blurts = getBlurtState()
	const users = getUserState()

	let footer: HTMLDivElement
	let loading_blurts = false

	onMount(() => {
		if (browser) {
			const handleIntersect: IntersectionObserverCallback = async (entries, observer) => {
				try {
					if (!entries[0].isIntersecting || loading_blurts) return console.log("false alarm")

					console.log("loading more blurts")
					loading_blurts = true

					const res = await fetch(
						API_ADDRESS + `/blurts/?offset=${blurts.list.length}&username=${users.active_user}`
					)

					if (!res.ok) return console.error("ERROR!", res.status, res.statusText)

					const new_blurts = await res.json()
					console.log(new_blurts.length)
					if (new_blurts.length == 0)
						return (console.log("unobserving"), observer.unobserve(footer))

					blurts.add_bulk(new_blurts)

					loading_blurts = false
					console.log("done loading more blurts")
				} catch (e) {
					console.log("yup")
				}
			}
			const options = { threshold: 0, rootMargin: "200px 0px 0px 0px" }
			const footer_observer = new IntersectionObserver(handleIntersect, options)
			footer_observer.observe(footer)
		}
	})
</script>

<div
	class="mx-auto my-2 flex w-full flex-col items-center px-2 md:max-w-screen-md md:flex-row md:gap-4"
>
	<PersonaManager />
	<div class="flex w-full flex-col gap-4">
		<NewBlurtForm />
		<BlurtList />

		<div class="mb-6 rounded-md bg-white p-4 text-center font-bold md:mb-0">
			<div class="text-teal-500">THAT'S ALL OF THEM!!!</div>
			<div class="text-teal-600">You win in a losing sort of way!</div>
		</div>

		<div bind:this={footer}></div>
	</div>
</div>
