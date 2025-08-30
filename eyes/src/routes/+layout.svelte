<script lang="ts">
	import "../app.css"
</script>

<svelte:head>
	<title>Everybody Blurt Now</title>
	<meta charset="UTF-8" />
	<meta name="viewport" content="width=device-width,initial-scale=1" />
	<meta
		name="description"
		content="It's like Twitter except each Blurt is capped at 14 characters. It's nonsense!"
	/>
	<link rel="icon" href="favicon.png" />

	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="true" />
	<link
		href="https://fonts.googleapis.com/css2?family=Caveat+Brush&family=Permanent+Marker&display=swap"
		rel="stylesheet"
	/>
</svelte:head>

<slot />
