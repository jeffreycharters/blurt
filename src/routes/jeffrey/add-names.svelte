<script>
	import User from '$lib/jeffrey/User.svelte';

	export let users;
	let newName = '';

	const addHandler = async (event) => {
		if (event.keyCode === 13) event.preventDefault();
		if (event.keyCode && event.keyCode != 13) return;
		// console.log(newName);
		const res = await fetch('./add-names', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Accept: 'application/json'
			},
			body: JSON.stringify({ username: newName })
		});
		if (!res.ok) {
			console.log('error saving user');
			return;
		}
		const user = await res.json();
		users = [user, ...users];
		newName = '';
	};
</script>

<div class="mx-auto min-w-sm w-full max-w-2xl xl:flex xl:gap-4 xl:justify-evenly">
	<div class="max-w-sm xl:mt-12 md:max-w-md">
		<div class="font-mono text-xl mb-6">/jeffrey/add-names</div>

		<div class="text-teal-800 text-center text-lg mb-4 font-semibold">Add Generator Names</div>

		<div class="flex justify-evenly">
			<input
				type="text"
				placeholder="some fake name"
				class="border rounded w-7/12 px-2 py-1 text-teal-800 text-semibold border-gray-100 shadow-inner"
				maxlength="14"
				bind:value={newName}
				on:keydown={addHandler}
			/>
			<input
				type="button"
				value="Add it!"
				class="w-3/12 bg-teal-700 text-white p-1 rounded font-semibold"
				on:click={addHandler}
			/>
		</div>

		<div class="text-sm my-6 p-2 rounded bg-slate-100 text-slate-700">
			Generator names are the ones that come up as options on the main login page when you click the
			refresh button.
		</div>
	</div>

	<div class="max-w-sm w-full md:max-w-md xl:mt-12">
		<div class="text-lg text-teal-800 font-semibold border-b">Recent</div>
		{#if users}
			{#each users as user (user.id)}
				<User {user} />
			{/each}
		{:else}
			<div class="text-sm m-2">Where my users at dog?</div>
		{/if}
	</div>
</div>
