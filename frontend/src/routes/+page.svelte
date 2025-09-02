<script lang="ts">
	let blurt = $state("")
	let countdown_box: HTMLDivElement

	function typeHandler() {
		const remaining = 14 - blurt.length
		const hue = (360 - 190) / remaining + 190
		const saturation = (100 - 60) / remaining + 60

		countdown_box.style.color = `hsl(${hue}, ${saturation}%, 50%)`
		countdown_box.innerHTML = String(remaining)
	}
</script>

<div class="mx-auto my-2 max-w-md px-2">
	<form
		class="flex items-baseline justify-between rounded bg-teal-100 p-2 shadow"
		onsubmit={console.log}
	>
		<label for="blurt" class="sr-only">blurt</label>
		<div class="flex items-baseline gap-2">
			<input
				class="w-40 rounded-md border border-solid border-teal-900 bg-teal-50 px-2 py-1 font-semibold text-teal-800"
				type="text"
				id="blurt"
				bind:value={blurt}
				placeholder="blurt here"
				autocomplete="off"
				maxlength="14"
				onkeyup={typeHandler}
			/>
			<div class="font-bold" style="color: hsl(190, 60%, 55%)" bind:this={countdown_box}>14</div>
		</div>

		<button
			type="submit"
			id="blurt-button"
			class="rounded-md border border-solid border-gray-900 bg-teal-600 px-4 py-1 font-semibold text-white"
			>blurt it</button
		>
	</form>
</div>
