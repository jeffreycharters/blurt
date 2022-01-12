<script>
	let username = '';
	const handleSubmit = async (event) => {
		localStorage.setItem('username', username);
		document.getElementById('login-form').submit();
	};
</script>

<div class="flex flex-col justify-evenly items-center m-0 h-96 lg:h-full">
	<div class="text-center bg-teal-50 rounded-md py-4 w-full max-w-xs">
		<h1 class="text-7xl text-teal-900" style="font-family: 'Permanent Marker';">BLURT</h1>

		<h2 class="text-xs text-teal-500">An exciting <em>new</em> way to shout at the world!</h2>
	</div>

	<form action="/login" method="post" on:submit|preventDefault={handleSubmit} id="login-form">
		<div class="text-teal-50 font-bold text-center">
			Enter a name and let's blurt!
			<input
				type="text"
				name="username"
				class="max-w-xs w-full py-1 px-2 rounded-md text-center text-teal-900"
				placeholder="blazer420 is taken, sorry"
				maxlength="14"
				bind:value={username}
			/>
			<button
				type="submit"
				class="bg-teal-900 w-full my-1 py-1 px-2 rounded-md uppercase font-semibold tracking-widest text-sm max-w-xs"
				style="font-family: 'Permanent Marker';">Let's Go!!</button
			>
		</div>
	</form>
</div>
