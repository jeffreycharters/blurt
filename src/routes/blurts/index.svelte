<script>
	import { onMount } from 'svelte';
	import { blurts } from '../../stores/blurts';
	import Blurt from './components/Blurt.svelte';

	let blurt = '';

	let username = '';

	onMount(async () => {
		username = localStorage.getItem('username');
	});

	const submitHandler = async () => {
		const response = await fetch('/blurts.json', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Accept: 'application/json'
			},
			body: JSON.stringify({ username: username, blurt: blurt })
		});
		blurt = '';
	};
</script>

<div class="max-w-md mx-auto">
	<form class="flex justify-between" on:submit|preventDefault={submitHandler}>
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
			id="blurt-button"
			class="border-2 border-solid border-gray-200 py-1 px-2 rounded-md font-semibold text-teal-50 bg-teal-600"
			>blurt it</button
		>
	</form>

	{#each $blurts as blurt}
		<Blurt>
			<span slot="blurt">{blurt.blurt}</span>
			<span slot="username">{blurt.user.username}</span>
			<span slot="date">{blurt.date}</span>
		</Blurt>
	{/each}
</div>
