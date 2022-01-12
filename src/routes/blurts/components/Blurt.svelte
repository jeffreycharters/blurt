<script>
	import { onMount } from 'svelte';
	export let blurt;
	let users_blurt = false;
	let username = '';
	let userid = -1;
	let likdBlurt = false;

	onMount(() => {
		username = localStorage.getItem('username');
		userid = localStorage.getItem('id');
		users_blurt = username === blurt.user.username;
		likdBlurt = !!checkBlurtLiks(blurt);
	});
	$: own_style = users_blurt ? 'bg-teal-50' : '';

	const likHandler = async () => {
		const res = await fetch(`/blurts/lik/${blurt.uid}`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ uid: blurt.uid, username: username })
		});
		document.getElementById(`lik-box-${blurt.uid}`).textContent = 'Likd!';
		return res;
	};

	const checkBlurtLiks = (blurt) => {
		if (blurt.liks.length > 0) {
			return blurt.liks.find((lik) => lik.user.username === username);
		}
		return false;
	};
</script>

<article class="mt-2 mb-4 py-2 px-4 border-grey-50 border-solid border-2 rounded-md {own_style}">
	<div class="text-teal-700 font-bold border-b-solid border-b-2 border-b-gray-400">
		{blurt.user.username}
	</div>
	<span
		class="font-semibold text-teal-600 text-4xl pl-6 tracking-wide"
		style="font-family: 'Caveat Brush'"
	>
		{blurt.blurt}
	</span>
	<div class="flex justify-between">
		<div class="flex gap-4">
			<div id="lik-box-{blurt.uid}" class="inline">
				{#if likdBlurt}
					Likd!
				{:else}
					<button type="button" on:click={likHandler} class="text-teal-800 cursor-pointer underline"
						>Lik</button
					>
				{/if}
			</div>
			<div class="text-teal-500 font-bold">
				{blurt.liks.length} lik
			</div>
		</div>
		<div>
			{blurt.date}
		</div>
	</div>
</article>
