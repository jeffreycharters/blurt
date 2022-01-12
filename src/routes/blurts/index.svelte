<script context="module">
	export const load = async ({ fetch }) => {
		const blurts = await fetch('/blurts.json');

		if (blurts.ok) {
			let dateBlurts = await blurts.json();
			for (let blurt of dateBlurts) {
				blurt.date = new Date(blurt.created_at);
			}
			return {
				props: {
					blurts: await dateBlurts
				}
			};
		}
		return {
			status: blurts.status,
			error: new Error(`Could not load blurts!`)
		};
	};
</script>

<script>
	import { onMount } from 'svelte';
	import Blurt from './components/Blurt.svelte';

	export let blurts;
	$: blurts;
	let blurt = 'testing';

	let username = '';

	onMount(async () => {
		username = localStorage.getItem('username');
	});

	const submitHandler = async (event) => {
		const response = await fetch('/blurts.json', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Accept: 'application/json'
			},
			body: JSON.stringify({ username: username, blurt: blurt })
		});
		const newBlurt = await response.json();
		if (response.status === 201) {
			blurts = [newBlurt, ...blurts];
			blurt = '';
		}
	};

	const deleteHandler = async (uid) => {
		const response = await fetch(`/locations/${uid}.json`, {
			method: 'DELETE'
		});
		if (response.status === 204) {
			const filteredLocations = locations.filter((loc) => loc.uid !== uid);
			locations = filteredLocations;
		}
	};
	const month = [
		'January',
		'February',
		'March',
		'April',
		'May',
		'June',
		'July',
		'August',
		'September',
		'October',
		'November',
		'December'
	];
</script>

<div class="max-w-md mx-auto">
	<h1
		class="text-3xl tracking-widest text-teal-600 text-center mb-4 border-b-2 border-b-solid border-b-gray-200 pb-2"
		style="font-family: 'Permanent Marker';"
	>
		Megablurts
	</h1>

	<form class="flex justify-between">
		<label for="blurt" class="hidden">blurt</label>
		<input
			class="py-1 px-2 rounded-md border-2 border-solid border-teal-900 bg-teal-50 text-teal-800 font-semibold"
			type="text"
			bind:value={blurt}
			id="blurt"
			placeholder="blurt here"
			maxlength="14"
		/>
		<button
			type="submit"
			on:click|preventDefault={submitHandler}
			class="border-2 border-solid border-gray-200 py-1 px-2 rounded-md font-semibold text-teal-50 bg-teal-600"
			>Add</button
		>
	</form>

	{#each blurts as blurt}
		<Blurt>
			<span slot="blurt">{blurt.blurt}</span>
			<span slot="username">{blurt.user.username}</span>
			<span slot="date"
				>{month[blurt.date.getMonth()]} {blurt.date.getDate()}, {blurt.date.getFullYear()}</span
			>
		</Blurt>
	{/each}
</div>
