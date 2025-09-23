<script lang="ts">
	import { API_ADDRESS } from "$lib"
	import mowr from "$lib/assets/mowr.png"
	import { getBlurtState, type Blurt } from "$lib/blurts.svelte"
	import { getUserState } from "$lib/users.svelte"

	const blurts = getBlurtState()

	let loading_blurts = $state(false)
	let more_blurts = $state(true)

	async function get_more_blurts() {
		if (loading_blurts) return

		console.log("blurt count", blurts.list.length)

		loading_blurts = true
		const res = await fetch(`${API_ADDRESS}/blurts?offset=${blurts.list.length}`)

		if (!res.ok) return console.error(res.status, res.statusText)

		const new_blurts = (await res.json()) as Blurt[] | null

		if (new_blurts) blurts.add_bulk(new_blurts)

		if (!new_blurts) more_blurts = false

		console.log("after blurt count", blurts.list.length)
		loading_blurts = false
	}
</script>

{#if more_blurts}
	<button
		class="mx-auto mb-16 flex w-5/6 cursor-pointer rounded bg-white py-2 shadow-lg hover:bg-teal-50 md:mb-0"
		onclick={get_more_blurts}
	>
		<div class="w-full text-right" style="font-family: 'comic sans ms';">clik heer for</div>
		<img src={mowr} alt="mowr blurts" class="ml-auto mr-0 h-16" />
	</button>
{:else}
	<footer class="mb-6 rounded-md bg-white p-4 text-center font-bold md:mb-0">
		<div class="text-teal-500">THAT'S ALL OF THEM!!!</div>
		<div class="text-teal-600">You win in a losing sort of way!</div>
	</footer>
{/if}
