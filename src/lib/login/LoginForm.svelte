<script>
	import { createEventDispatcher } from 'svelte';

	export let username;

	let dispatch = createEventDispatcher();

	const handleSubmit = () => {
		dispatch('loginSubmit', username);
	};
</script>

<form action="/login" method="post" on:submit|preventDefault={handleSubmit} id="login-form">
	<div class="text-teal-50 font-bold text-center flex flex-col gap-1 sm:text-lg md:text-xl">
		<div class="md:my-2">Enter a name and let's blurt!</div>
		<input
			type="text"
			name="username"
			class="max-w-xs w-full py-1 px-2 rounded-md text-center text-teal-900 mx-auto sm:py-2 md:py-3 md:max-w-md"
			placeholder="blazer420 is taken, sorry"
			maxlength="14"
			autocomplete="off"
			bind:value={username}
		/>
		<button
			type="submit"
			class="bg-teal-900 w-full my-1 py-1 px-2 rounded-md uppercase font-semibold tracking-widest text-sm max-w-xs mx-auto sm:py-2 sm:text-lg md:text-lg  md:max-w-md"
			style="font-family: 'Permanent Marker';">Let's Go!!</button
		>
	</div>
</form>
