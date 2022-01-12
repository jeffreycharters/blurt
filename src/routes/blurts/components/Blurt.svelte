<script>
	import { onMount } from 'svelte';
	export let blurt;
	let users_blurt = false;
	let username = '';
	let userid = -1;
	let likdBlurt = false;
	let displayDate = '';

	onMount(() => {
		username = localStorage.getItem('username');
		userid = localStorage.getItem('id');
		users_blurt = username === blurt.user.username;
		likdBlurt = !!checkBlurtLiks(blurt);
		const date = new Date();
		const blurtDate = new Date(blurt.created_at);
		displayDate = getDisplayDate(blurtDate, date);
	});
	$: own_style = users_blurt ? 'bg-teal-50' : '';

	const getDisplayDate = (blurtDate, date) => {
		const timeGap = date - blurtDate;
		const fewSeconds = 1000 * 15;
		const fewMinutes = 1000 * 60 * 45;
		const oneHour = 1000 * 60 * 60;
		const hourish = 1000 * 60 * 90;
		const coupleHours = 1000 * 60 * 180;
		const oneDay = 1000 * 60 * 60 * 24;
		if (timeGap < fewSeconds) {
			return 'A few seconds ago.';
		} else if (timeGap < fewMinutes) {
			return 'Less than an hour ago.';
		} else if (timeGap < hourish) {
			return 'About an hour ago.';
		} else if (timeGap < coupleHours) {
			return 'A couple of hours ago.';
		} else if (timeGap < oneDay) {
			return `About ${Math.ceil(timeGap / oneHour)} hours ago.`;
		} else {
			return blurt.date;
		}
	};

	const likHandler = async () => {
		const res = await fetch(`/blurts/lik/${blurt.uid}`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ uid: blurt.uid, username: username })
		});
		document.getElementById(`lik-box-${blurt.uid}`).innerHTML =
			'<div class="text-gray-300 font-bold px-2">Likd!</div>';
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
	<div class="text-teal-700 font-bold border-b-solid border-b-2 border-b-gray-400 pl-2">
		{blurt.user.username}
	</div>
	<div
		class="font-semibold text-teal-600 text-4xl pl-6 tracking-wider my-2"
		style="font-family: 'Caveat Brush'"
	>
		{blurt.blurt}
	</div>
	<div class="flex justify-between">
		<div class="flex gap-4">
			<div id="lik-box-{blurt.uid}" class="inline">
				{#if likdBlurt}
					<div class="text-gray-300 font-bold px-2">u Lik!</div>
				{:else}
					<button
						type="button"
						on:click={likHandler}
						class="text-teal-600 cursor-pointer border-grey-200 rounded-md border-2 px-4 shadow-md hover:shadow-sm font-bold tracking-wider"
						>Lik</button
					>
				{/if}
			</div>
			<div class="text-teal-500 font-bold">
				{blurt.liks.length} lik
			</div>
		</div>
		<div class="text-gray-400">
			{displayDate}
		</div>
	</div>
</article>
