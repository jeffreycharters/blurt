<script lang="ts">
	import { fade } from "svelte/transition"
	import type { Blurt } from "../../app"
	import { displayDate } from "$lib"

	export let userID: string = ""
	export let blurt: Blurt
	export let likHandler: (blurt: Blurt) => boolean

	$: likdBlurt = blurt.edges.liks?.find((lik) => {
		return lik.user_id === userID
	})
	$: blurtLiks = blurt.edges.liks?.length ?? 0
	let error = ""

	const touchHandler = (e: TouchEvent) => {
		if (e.target) (e.currentTarget as HTMLDivElement).style.transform = "scale(0.98)"
	}

	const touchEndHandler = (e: TouchEvent) => {
		if (e.target) {
			const currDiv = e.currentTarget as HTMLDivElement

			currDiv.style.transform = "scale(1)"
			currDiv.style.boxShadow = ""
		}
	}

	const own_style = blurt.edges.author.id === userID ? "bg-teal-50" : "bg-white"
</script>

<div on:touchstart={touchHandler} on:touchend={touchEndHandler} class="blurt-holder">
	<article class="border-grey-50 mb-4 mt-2 rounded-md border border-solid px-4 py-2 {own_style}">
		<div
			class="border-b-solid border-b border-b-gray-400 pb-1 pl-2 font-bold tracking-wider text-teal-500"
		>
			{decodeURI(blurt.edges.author.username)}
			<img src="verified.svg" alt="verified" class="relative inline h-4" style="top: -2px;" />
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
					{#if likdBlurt}
						<div
							transition:fade|global
							class="border-grey-200 w-24 rounded-md border bg-gray-50 px-4 font-bold tracking-wider text-slate-400 shadow-sm grayscale-[50%] hover:shadow-none"
						>
							<img
								src="heart-icon.svg"
								alt="heart"
								class="relative -ml-1 mr-1 inline h-4 p-0"
								style="top: -2px;"
							/>Likd!
						</div>
					{:else}
						<button
							type="button"
							on:click={() => likHandler(blurt)}
							class="border-grey-200 w-24 cursor-pointer rounded-md border bg-white px-4 font-bold tracking-wider text-teal-600 shadow-sm hover:shadow-none"
							><img
								src="heart-icon.svg"
								alt="heart"
								class="relative -ml-1 mr-1 inline h-4 p-0"
								style="top: -2px;"
							/>Lik</button
						>
					{/if}
				</div>
				<div class="font-bold text-teal-500">
					{blurtLiks} lik{blurtLiks === 1 ? "" : "s"}
				</div>
			</div>
			<div class="text-xs text-gray-400">
				{displayDate(new Date(blurt.create_time), new Date())}
			</div>
		</div>
		{#if error}
			<div transition:fade|global class="mx-4 mt-4 text-sm text-red-600">{error}</div>
		{/if}
	</article>
</div>

<style>
	.blurt-holder {
		transition: all 0.075s cubic-bezier(0.17, 0.67, 0.88, 4.06);
	}
</style>
