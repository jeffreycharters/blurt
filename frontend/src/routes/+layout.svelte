<script lang="ts">
	import "../app.css"
	import favicon from "$lib/assets/favicon.svg"
	import { setUserState } from "$lib/users.svelte"
	import { fade } from "svelte/transition"
	import { setBlurtState } from "$lib/blurts.svelte"
	import { setChaosState } from "$lib/chaos.svelte"
	import { setPersonaDrawerState } from "$lib/personas.svelte"
	import { onDestroy, onMount } from "svelte"

	let { children, data } = $props()

	const users = setUserState()
	const _blurts = setBlurtState(data.blurts)
	const _chaos = setChaosState()
	const _persona_drawer = setPersonaDrawerState()

	let username = $state("")

	let dialog: HTMLDialogElement
	$effect(() => {
		if (!users.active_user) dialog.showModal()
	})

	function login(event: Event) {
		event.preventDefault()
		if (username.length > 14) username = username.slice(0, 14)
		users.add(username)
		dialog.close()
	}
</script>

<svelte:head>
	<title>Blurtonamo Bay</title>
	<link rel="icon" href={favicon} />
</svelte:head>

<dialog transition:fade bind:this={dialog} class="backdrop:bg-black/50 backdrop:backdrop-blur">
	<div
		class="top-4/12 fixed left-1/2 h-1/2 -translate-x-1/2 -translate-y-1/2 transform rounded-lg bg-white p-4 shadow-lg"
	>
		<div class="flex h-full w-full flex-col justify-between">
			<div class="md:min-w-md w-full rounded-md bg-teal-700 py-4 text-center">
				<h1 class="px-8 text-7xl text-teal-50" style="font-family: 'Permanent Marker';">BLURT</h1>

				<h2 class="text-xs text-teal-100">An exciting <em>new</em> way to shout at the world!</h2>
			</div>

			<div class="flex-grow"></div>

			<div class="mx-auto">
				<form onsubmit={login}>
					<div class="flex flex-col gap-1 text-center text-sm font-bold text-teal-600">
						Enter a name and let's blurt!
						<input
							type="text"
							name="username"
							class="w-full max-w-xs rounded-md px-2 py-1 text-center text-teal-900"
							placeholder="blazer420 is taken, sorry"
							maxlength="14"
							autocomplete="off"
							bind:value={username}
						/>
						<button
							type="submit"
							class="my-1 w-full max-w-xs rounded-md bg-teal-900 px-2 py-1 text-sm uppercase tracking-widest text-teal-100"
							style="font-family: 'Permanent Marker';">Let's Go!!</button
						>
					</div>
				</form>
			</div>

			<div class="flex-grow"></div>
		</div>
	</div>
</dialog>

<div class="min-h-screen min-w-full bg-teal-100 p-2">
	{@render children?.()}
</div>
