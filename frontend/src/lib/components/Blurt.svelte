<script lang="ts">
	import { getDisplayDate } from "$lib"
	import { type Blurt } from "$lib/blurts.svelte"
	import { getUserState } from "$lib/users.svelte"
	import { fade } from "svelte/transition"
	import heart from "$lib/assets/heart.svg"
	import verified from "$lib/assets/verified.svg"

	interface Props {
		blurt: Blurt
		blurt_likker: (id: string, error_func: () => void) => void
	}

	let { blurt, blurt_likker }: Props = $props()

	const users = getUserState()

	let error = $state("")
	let error_func = () => {
		error = "Shit! Trouble likking! Try lik later again."
		setTimeout(() => (error = ""), 5000)
	}

	let this_blurt: HTMLElement | undefined

	function touchHandler() {
		if (!this_blurt) return

		this_blurt.style.transform = "scale(0.98)"
	}

	function touchEndHandler() {
		if (!this_blurt) return

		this_blurt.style.transform = "scale(1)"
		this_blurt.style.boxShadow = ""
	}
</script>

<article
	bind:this={this_blurt}
	class="blurt-holder mx-1 rounded-md px-4 py-2 shadow-md transition-colors duration-500 md:m-0 {users.active_user ===
	blurt.author
		? 'bg-teal-50'
		: 'bg-white'}"
	ontouchstart={touchHandler}
	ontouchend={touchEndHandler}
>
	<div
		class="border-b-solid border-b border-b-gray-400 pb-1 pl-2 font-bold tracking-wider text-teal-500"
	>
		{blurt.author}
		<img src={verified} alt="verified" class="relative inline h-4" style="top: -2px;" />
	</div>
	<div
		class="my-2 pl-6 text-4xl font-semibold tracking-wider text-teal-600"
		style="font-family: 'Caveat Brush'"
	>
		{blurt.content}
	</div>
	<div class="flex items-baseline justify-between">
		<div class="flex items-baseline gap-4">
			<div id="lik-box-{blurt.id}" class="inline">
				{#if blurt.likkers?.includes(users.active_user)}
					<div
						class="border-grey-200 w-24 rounded-md border bg-gray-50 px-4 font-bold tracking-wider text-slate-400 shadow-sm grayscale-[50%] hover:shadow-none"
					>
						<img
							src={heart}
							alt="heart"
							class="relative z-10 -ml-1 mr-1 inline h-4 p-0"
							style="top: -2px;"
						/>Likd!
					</div>
				{:else}
					<button
						type="button"
						onclick={() => blurt_likker(blurt.id, error_func)}
						class="border-grey-200 w-24 cursor-pointer rounded-md border bg-white px-4 font-bold tracking-wider text-teal-600 shadow-sm hover:shadow-none"
						><img
							src={heart}
							alt="heart"
							class="relative z-10 -ml-1 mr-1 inline h-4 p-0"
							style="top: -2px;"
						/>Lik</button
					>
				{/if}
			</div>
			<div class="font-bold text-teal-500">
				{blurt.likkers?.length ?? 0} lik
			</div>
		</div>
		<div class="text-xs text-gray-400">
			{getDisplayDate(new Date(blurt.created), new Date())}
		</div>
	</div>
	{#if error}
		<div transition:fade class="mx-4 mt-4 text-sm text-red-600">{error}</div>
	{/if}
</article>

<style lang="postcss">
	.blurt-holder {
		transition: all 0.075s cubic-bezier(0.17, 0.67, 0.88, 4.06);
	}
</style>
