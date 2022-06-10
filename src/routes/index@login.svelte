<script>
	import LoginForm from '$lib/login/LoginForm.svelte';
	let username = '';

	const handleSubmit = async (username) => {
		if (username.detail === '') username.detail = 'h4x0r';
		localStorage.setItem('username', username.detail);
		document.getElementById('login-form').submit();
	};
</script>

<div
	class="flex flex-col justify-between items-center bg-teal-500 px-2 rounded-lg shadow-md mt-12 mx-2 sm:max-w-4xl sm:mx-auto"
>
	<div class="text-center bg-teal-50 rounded-md py-4 w-full max-w-xs mt-24 max-w-lg">
		<h1
			class="text-7xl text-teal-900 sm:text-8xl md:text-9xl"
			style="font-family: 'Permanent Marker';"
		>
			BLURT
		</h1>

		<h2 class="text-xs text-teal-500 sm:text-base md:text-xl md:my-4">
			An exciting <em>new</em> way to shout at the world!
		</h2>
	</div>

	<div class="w-full my-24">
		<LoginForm {username} on:loginSubmit={handleSubmit} />
	</div>
</div>
