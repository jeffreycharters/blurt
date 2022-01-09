<script>
	const trails = [
		{
			uid: 'fgdfhdfg',
			name: 'South Park'
		},
		{
			uid: 'sfgkjhdfg',
			name: 'Homestead'
		}
	];
</script>

Log volunteer time:

<form>
	<label for="" /><input type="text" />
</form>

{#each trails as trail}
	{trail.name}<br />
{/each}
