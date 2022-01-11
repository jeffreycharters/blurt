<script>
	import '../tailwind.css';
</script>

<svelte:head>
	<title>Everybody Blurt Now</title>
	<meta charset="UTF-8" />
	<meta name="viewport" content="width=device-width,initial-scale=1" />
	<meta name="description" content="" />
	<link rel="icon" href="favicon.png" />

	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
	<link
		href="https://fonts.googleapis.com/css2?family=Permanent+Marker&display=swap"
		rel="stylesheet"
	/>
</svelte:head>
<body class="py-2 px-2 min-w-full min-h-full bg-teal-500">
	<slot />
</body>
