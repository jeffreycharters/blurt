<script>
	import { onMount } from 'svelte';
	import { fade } from 'svelte/transition';
	export let blurt;
	let users_blurt = false;
	let username = '';
	let likdBlurt = blurt.userLikd;
	$: blurtLiks = blurt._count?.liks ?? 0;
	let error = '';

	let displayDate = '';

	onMount(() => {
		username = localStorage.getItem('username');
		users_blurt = username === blurt.user.username;
		const date = new Date();
		const blurtDate = new Date(blurt.created_at);
		displayDate = getDisplayDate(blurtDate, date);
	});
	$: own_style = users_blurt ? 'bg-teal-50' : 'bg-white';

	const getDisplayDate = (blurtDate, date) => {
		const timeGap = date - blurtDate;
		const fewSeconds = 1000 * 60;
		const fewMinutes = 1000 * 60 * 15;
		// const aboutHour = 1000 * 60 * 45;
		const oneHour = 1000 * 60 * 60;
		const hourish = 1000 * 60 * 90;
		const coupleHours = 1000 * 60 * 180;
		const oneDay = 1000 * 60 * 60 * 24;
		if (timeGap < fewSeconds) {
			return 'A few seconds ago.';
		} else if (timeGap < fewMinutes) {
			return 'A few minutes ago.';
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
		blurtLiks += 1;
		likdBlurt = true;
		const res = await fetch(`/blurts/lik/${blurt.uid}`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ uid: blurt.uid, username: username })
		});
		if (!res.ok) {
			likdBlurt = false;
			blurtLiks -= 1;
			error = 'Shit! Trouble likking! Try lik later again.';
			setTimeout(() => (error = ''), 5000);
		}
		return res;
	};
</script>

<article class="mt-2 mb-4 py-2 px-4 border-grey-50 border-solid border rounded-md {own_style}">
	<div
		class="text-teal-500 font-bold border-b-solid border-b border-b-gray-400 pl-2 tracking-wider pb-1"
	>
		{blurt.user.username}
		<img src="verified.svg" alt="verified" class="h-4 inline relative" style="top: -2px;" />
	</div>
	<div
		class="font-semibold text-teal-600 text-4xl pl-6 tracking-wider my-2"
		style="font-family: 'Caveat Brush'"
	>
		{blurt.blurt}
	</div>
	<div class="flex justify-between items-baseline">
		<div class="flex gap-4 items-baseline">
			<div id="lik-box-{blurt.uid}" class="inline">
				{#if likdBlurt}
					<div
						transition:fade|global
						class="border-grey-200 rounded-md border px-4 shadow-sm hover:shadow-none font-bold tracking-wider bg-gray-50 text-slate-400 w-24 grayscale-[50%]"
					>
						<img
							src="heart-icon.svg"
							alt="heart"
							class="h-4 inline relative p-0 -ml-1 mr-1"
							style="top: -2px;"
						/>Likd!
					</div>
				{:else}
					<button
						type="button"
						on:click={likHandler}
						class="text-teal-600 cursor-pointer border-grey-200 rounded-md border px-4 shadow-sm hover:shadow-none font-bold tracking-wider w-24 bg-white"
						><img
							src="heart-icon.svg"
							alt="heart"
							class="h-4 inline relative p-0 -ml-1 mr-1 "
							style="top: -2px;"
						/>Lik</button
					>
				{/if}
			</div>
			<div class="text-teal-500 font-bold">
				{blurtLiks} lik
			</div>
		</div>
		<div class="text-gray-400 text-xs">
			{displayDate}
		</div>
	</div>
	{#if error}
		<div transition:fade|global class="text-sm mt-4 mx-4 text-red-600">{error}</div>
	{/if}
</article>
