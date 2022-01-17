<script context="module">
	import { blurts } from '../stores/blurts';
	import { humanizeDates } from './blurts/utils';

	export const load = async ({ fetch, url }) => {
		const blurtUrl = `${url.origin}/blurts.json`;
		const response = await fetch(blurtUrl);
		const loadedBlurts = await response.json();
		const readyBlurts = humanizeDates(loadedBlurts);
		blurts.set(readyBlurts);
		return {};
	};
</script>

<script>
	import LoginForm from './login/LoginForm.svelte';
	let username = '';

	const handleSubmit = async (username) => {
		if (username === '') {
			username = 'no.';
			return;
		}
		console.log(username);
		localStorage.setItem('username', username.detail.username);
		document.getElementById('login-form').submit();
	};
</script>

<div class="flex flex-col justify-between items-center m-0 h-96 max-w-xs mx-auto">
	<div class="text-center bg-teal-50 rounded-md py-4 w-full max-w-xs mt-24">
		<h1 class="text-7xl text-teal-900" style="font-family: 'Permanent Marker';">BLURT</h1>

		<h2 class="text-xs text-teal-500">An exciting <em>new</em> way to shout at the world!</h2>
	</div>

	<div class="w-full mt-24">
		<LoginForm {username} on:loginSubmit={handleSubmit} />
	</div>
</div>
