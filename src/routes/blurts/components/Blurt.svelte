<script>
	import { onMount } from 'svelte';
	export let blurt;
	let users_blurt = false;
	onMount(() => {
		let username = localStorage.getItem('username');
		users_blurt = username === blurt.user.username;
		console.log({ user: blurt.user.username, username });
	});
	$: own_style = users_blurt ? 'bg-teal-50 rounded-md border-grey-100 border-solid border-2' : '';
</script>

<article class="mt-2 mb-4 py-2 px-4 {own_style}">
	<h2 class="font-semibold text-teal-600 text-4xl" style="font-family: 'Caveat Brush'">
		{blurt.blurt}
	</h2>
	<div class="flex justify-between">
		<div>
			<a href="/blurt/lik/{blurt.uid}" class="text-teal-800 cursor-pointer underline">lik</a>
			{blurt.user.username}
		</div>
		<div>
			{blurt.date}
		</div>
	</div>
</article>
