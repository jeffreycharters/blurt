<script>
	export let user;
	export let editing = false;

	let username = user.username;

	const toggleEdit = () => {
		editing = !editing;
	};

	const nameFocus = (element) => {
		element.focus();
	};

	const updateName = async () => {
		if (user.username === username) {
			toggleEdit();
			return;
		}
		console.log(`updating to ${username}`);
		const res = await fetch('./add-names', {
			method: 'PATCH',
			headers: {
				'Content-Type': 'application/json',
				Accept: 'application/json'
			},
			body: JSON.stringify({ id: user.id, username: username })
		});
		if (!res.ok) {
			console.log('error updating username!');
		}
		user = await res.json();
		toggleEdit();
	};
</script>

<div
	class="bg-slate-100 m-2 py-1 px-2 rounded text-teal-700 font-semibold flex justify-between items-center h-10"
>
	<div class="flex items-center">
		<span class="text-slate-500 mr-2">
			<svg
				xmlns="http://www.w3.org/2000/svg"
				class="icon icon-tabler icon-tabler-user"
				width="18"
				height="18"
				viewBox="0 0 24 24"
				stroke-width="1.75"
				stroke="currentColor"
				fill="none"
				stroke-linecap="round"
				stroke-linejoin="round"
			>
				<path stroke="none" d="M0 0h24v24H0z" fill="none" />

				<circle cx="12" cy="7" r="4" />
				<path d="M6 21v-2a4 4 0 0 1 4 -4h4a4 4 0 0 1 4 4v2" />
			</svg>
		</span>

		{#if editing}
			<input
				type="text"
				bind:value={username}
				use:nameFocus
				on:blur={updateName}
				on:keydown={(event) => {
					if (event.code === 'Enter') updateName();
				}}
				class="w-36 rounded bg-gray-100 font-semibold"
				maxlength="14"
			/>
		{:else}
			<div on:click={toggleEdit} class="whitespace-nowrap">{user.username}</div>
		{/if}
	</div>

	{#if !editing}
		<div class="bg-green-100 py-1 px-2 rounded" on:click={toggleEdit}>
			<svg
				xmlns="http://www.w3.org/2000/svg"
				class="icon icon-tabler icon-tabler-edit"
				width="18"
				height="18"
				viewBox="0 0 24 24"
				stroke-width="1.75"
				stroke="currentColor"
				fill="none"
				stroke-linecap="round"
				stroke-linejoin="round"
			>
				<path stroke="none" d="M0 0h24v24H0z" fill="none" />
				<path d="M7 7h-1a2 2 0 0 0 -2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2 -2v-1" />
				<path d="M20.385 6.585a2.1 2.1 0 0 0 -2.97 -2.97l-8.415 8.385v3h3l8.385 -8.415z" />
				<path d="M16 5l3 3" />
			</svg>
		</div>
	{:else}
		<input
			type="button"
			value="Add it!"
			class="w-3/12 bg-teal-700 text-white p-1 rounded font-semibold text-sm"
			on:click={updateName}
		/>
	{/if}
</div>
