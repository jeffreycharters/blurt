<script>
	import '../../tailwind.css';
	import favicon from '$lib/assets/favicon.ico';
</script>

<svelte:head>
	<title>Blurffrey</title>
	<meta charset="UTF-8" />
	<meta name="viewport" content="width=device-width,initial-scale=1" />
	<meta name="description" content="Lickum!" />
	<link rel="icon" href={favicon} />
</svelte:head>

<slot />
