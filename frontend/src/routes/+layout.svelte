<script lang="ts">
	import "../app.css"
	import favicon from "$lib/assets/favicon.svg"

	let { children } = $props()
</script>

<svelte:head>
	<title>Blurtonamo Bay</title>
	<link rel="icon" href={favicon} />
</svelte:head>

<div class="min-h-screen min-w-full bg-teal-50 px-2 py-2">
	{@render children?.()}
</div>
