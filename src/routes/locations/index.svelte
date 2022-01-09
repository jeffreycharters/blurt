<script context="module">
	export const load = async ({ fetch }) => {
		const locations = await fetch('/locations.json');

		if (locations.ok) {
			return {
				props: {
					locations: await locations.json()
				}
			};
		}
		return {
			status: locations.status,
			error: new Error(`Could not load locations!`)
		};
	};
</script>

<script>
	export let locations;
	$: locations;
	let name = '';

	const submitHandler = async (event) => {
		const response = await fetch('/locations.json', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Accept: 'application/json'
			},
			body: JSON.stringify({ name: name })
		});
		const newLocation = await response.json();
		if (response.status === 201) {
			locations = [...locations, newLocation];
			name = '';
		}
	};

	const deleteHandler = async (uid) => {
		const response = await fetch(`/locations/${uid}.json`, {
			method: 'DELETE'
		});
		if (response.status === 204) {
			const filteredLocations = locations.filter((loc) => loc.uid !== uid);
			locations = filteredLocations;
		}
	};
</script>

<h1>Current locations list:</h1>

<h2>Add new</h2>

<form>
	<label for="name">Name</label>
	<input type="text" bind:value={name} />
	<button type="submit" on:click|preventDefault={submitHandler}>Add</button>
</form>

<ul>
	{#each locations as location}
		<li>{location.name} - <span on:click={() => deleteHandler(location.uid)}>delete</span></li>
	{/each}
</ul>
