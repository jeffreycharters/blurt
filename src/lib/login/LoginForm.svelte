<script>
	import { createEventDispatcher } from 'svelte';

	export let username;

	let intialPlaceholder = 'blazer420 is taken, sorry';
	let placeholder = intialPlaceholder;
	let dispatch = createEventDispatcher();

	const handleSubmit = () => {
		dispatch('loginSubmit', username);
	};

	const generator = async () => {
		username = '';
		placeholder = 'generating. . .';
		setTimeout(async () => {
			const res = await fetch('./users');
			if (!res.ok) {
				placeholder = 'something broke :(';
				return;
			}
			const body = await res.json();
			username = body.username;
		}, 250);
	};
</script>

<form action="/login" method="post" on:submit|preventDefault={handleSubmit} id="login-form">
	<div class="text-teal-50 font-bold text-center flex flex-col gap-1 sm:text-lg md:text-xl">
		<div class="md:my-2">Enter a name and let's blurt!</div>
		<div class="relative max-w-sm mx-auto md:max-w-md w-full">
			<input
				type="text"
				name="username"
				class="max-w-xs w-full py-1 px-2 rounded-md text-center text-teal-900 mx-auto sm:py-2 md:py-3 md:max-w-md"
				{placeholder}
				maxlength="14"
				autocomplete="off"
				bind:value={username}
			/>
			<div class="absolute top-[0.125rem] right-2 text-teal-800" on:click={generator}>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					class="icon icon-tabler icon-tabler-arrows-shuffle"
					width="28"
					height="28"
					viewBox="0 0 24 24"
					stroke-width="1.75"
					stroke="currentColor"
					fill="none"
					stroke-linecap="round"
					stroke-linejoin="round"
				>
					<path stroke="none" d="M0 0h24v24H0z" fill="none" />
					<path d="M18 4l3 3l-3 3" />
					<path d="M18 20l3 -3l-3 -3" />
					<path d="M3 7h3a5 5 0 0 1 5 5a5 5 0 0 0 5 5h5" />
					<path
						d="M21 7h-5a4.978 4.978 0 0 0 -2.998 .998m-4.002 8.003a4.984 4.984 0 0 1 -3 .999h-3"
					/>
				</svg>
			</div>
		</div>
		<button
			type="submit"
			class="bg-teal-900 w-full my-1 py-1 px-2 rounded-md uppercase font-semibold tracking-widest text-sm max-w-xs mx-auto sm:py-2 sm:text-lg md:text-lg  md:max-w-md"
			style="font-family: 'Permanent Marker';">Let's Go!!</button
		>
	</div>
</form>
