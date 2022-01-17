<script>
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { blurts } from '../../stores/blurts';
	import Blurt from './components/Blurt.svelte';
	import { humanizeDates } from './utils';

	let blurt = '';
	let username = '';

	$: displayBlurts = $blurts;

	const updateBlurts = async () => {
		const path = $page.url.origin;
		const url = `${path}/blurts.json`;
		const res = await fetch(url);
		const data = await res.json();
		const datedBlurts = humanizeDates(data);
		blurts.set(datedBlurts);
	};

	setInterval(updateBlurts, 2000);

	onMount(async () => {
		username = localStorage.getItem('username');
		const res = await fetch('/blurts.json');
		const rawBlurts = await (await res).json();
		const dateBlurts = humanizeDates(rawBlurts);
		blurts.set(dateBlurts);
	});

	const typeHandler = () => {
		const remaining = 14 - blurt.length;
		const hue = (360 - 190) / remaining + 190;
		const saturation = (100 - 60) / remaining + 60;
		const countdownDiv = document.getElementById('countdown-box');
		countdownDiv.style.color = `hsl(${hue}, ${saturation}%, 50%)`;
		countdownDiv.innerHTML = remaining;
	};

	const submitHandler = async () => {
		await fetch('/blurts.json', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Accept: 'application/json'
			},
			body: JSON.stringify({ username: username, blurt: blurt })
		});
		blurt = '';
		console.log('what');
		const countdownDiv = document.getElementById('countdown-box');
		countdownDiv.style.color = `hsl(190, 60%, 50%)`;
		countdownDiv.textContent = '14';
	};
</script>

<div class="max-w-md mx-auto">
	<form class="flex justify-between items-baseline" on:submit|preventDefault={submitHandler}>
		<label for="blurt" class="hidden">blurt</label>
		<div class="flex items-baseline gap-2">
			<input
				class="py-1 px-2 rounded-md border-2 border-solid border-teal-900 bg-teal-50 text-teal-800 font-semibold w-40"
				type="text"
				bind:value={blurt}
				id="blurt"
				placeholder="blurt here"
				autocomplete="off"
				maxlength="14"
				on:keyup={typeHandler}
			/>
			<div class="font-bold" style="color: hsl(190, 60%, 55%)" id="countdown-box">14</div>
		</div>
		<button
			type="submit"
			id="blurt-button"
			class="border-2 border-solid border-gray-200 py-1 px-2 rounded-md font-semibold text-teal-50 bg-teal-600"
			>blurt it</button
		>
	</form>

	{#each displayBlurts as blurt (blurt.uid)}
		<Blurt {blurt} />
	{/each}
</div>
