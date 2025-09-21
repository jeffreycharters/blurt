<script lang="ts">
	import { getChaosState } from "$lib/chaos.svelte"
	import { getUserState } from "$lib/users.svelte"
	import { ReplaceIcon, UserIcon, XIcon } from "$lib/components"
	import ArrowsIcon from "./ArrowsIcon.svelte"
	import { getPersonaDrawerState } from "$lib/personas.svelte"
	import { crossfade, slide } from "svelte/transition"
	import { flip } from "svelte/animate"

	const [send, receive] = crossfade({
		duration: (d) => Math.sqrt(d * 200)
	})

	let users = getUserState()

	let adding = $state(false)
	let new_username = $state("")

	const chaos = getChaosState()
	let checkbox_chaos = $state(chaos.active)

	const drawer = getPersonaDrawerState()

	function remove_user(username: string) {
		users.remove(username)
	}

	async function add_user(event: Event) {
		event.preventDefault()

		users.add(new_username)

		new_username = ""
		adding = false
	}

	function setChaos() {
		chaos.set_chaos(checkbox_chaos)
	}

	function drawer_key_handler(event: KeyboardEvent) {
		if (!["Enter", " "].includes(event.key)) return

		event.preventDefault()
		drawer.toggle()
	}
</script>

<div
	class="z-100 fixed bottom-2 left-1/2 flex w-full max-w-sm -translate-x-1/2 flex-col-reverse self-start rounded-lg bg-teal-700 p-2 shadow-md md:sticky md:bottom-auto md:left-auto md:top-2 md:mb-0 md:mb-8 md:max-w-sm md:translate-x-0 md:flex-col md:bg-white"
>
	<div
		class="flex w-full items-center justify-between rounded-md bg-slate-50 p-2"
		role="button"
		tabindex="0"
		onclick={() => drawer.toggle()}
		onkeydown={drawer_key_handler}
	>
		<div class="flex items-baseline gap-1">
			<UserIcon />
			<span class="font-bold text-teal-800">{users.active_user}</span>
		</div>
		<div class="flex gap-3">
			<button
				class="rounded border border-slate-400 bg-slate-200 p-[2px] shadow"
				onclick={(event) => (event.stopPropagation(), users.set_random_active())}
			>
				<ReplaceIcon />
			</button>
			<button
				class="rounded border border-slate-400 bg-slate-200 p-[2px] shadow"
				onclick={(event) => (event.stopPropagation(), drawer.toggle())}
			>
				<ArrowsIcon />
			</button>
		</div>
	</div>

	{#if drawer.open}
		<div class="mb-3 md:mb-0" in:slide={{ duration: 200 }} out:slide={{ duration: 125 }}>
			<hr class="mx-auto my-1 hidden w-5/6 rounded border border-slate-200 p-0 md:block" />

			<div class="flex items-baseline justify-between">
				<div class="font-bold text-teal-50 md:text-teal-900">Current Personas</div>
				<div class="text-xs italic text-teal-100 md:text-slate-500">{users.list.length} of 14</div>
			</div>
			<ul class="text-slate-600">
				{#each users.alphabetical_list as user (user)}
					<li
						in:receive={{ key: user }}
						out:send={{ key: user }}
						animate:flip={{ duration: 200 }}
						class:font-bold={user === users.active_user}
						class:text-teal-700={user === users.active_user}
						class:bg-teal-50={user === users.active_user}
						class:text-lg={user === users.active_user}
						class="mb-1 flex cursor-pointer rounded border border-slate-300 bg-slate-50 px-2 py-1 md:bg-slate-100"
					>
						<button
							class="flex-grow cursor-pointer text-left"
							onclick={() => users.set_active_user(user)}
							>{user}
							{#if user === users.active_user}
								<span class="text-sm font-normal italic text-slate-500">(active)</span>
							{/if}
						</button>
						<button onclick={() => remove_user(user)} class="cursor-pointer">
							<XIcon />
						</button>
					</li>
				{/each}
				<li
					class="flex w-full justify-between text-right text-sm font-bold text-slate-200 md:text-slate-400"
				>
					{#if !adding}
						<div
							class="flex items-center gap-2 rounded px-2 text-red-900 text-slate-700 md:bg-slate-100"
						>
							<input
								type="checkbox"
								name="chaos"
								onchange={setChaos}
								bind:checked={checkbox_chaos}
								id="chaos"
								class="h-4 w-4 rounded-sm text-red-700"
							/>
							<label
								for="chaos"
								class="text-teal-50 md:text-slate-900"
								class:font-bold={checkbox_chaos}>Chaos mode</label
							>
						</div>

						<button
							class="h-8 hover:cursor-pointer"
							onclick={() => (adding = true)}
							disabled={users.list.length >= 14}>+ add persona</button
						>
					{:else}
						<form class="flex w-full justify-between" onsubmit={add_user}>
							<div>
								<label for="add-user" class="sr-only"></label>
								<!-- svelte-ignore a11y_autofocus -->
								<input
									type="text"
									bind:value={new_username}
									name="add-user"
									minlength="1"
									maxlength="14"
									class="h-8 w-40 rounded border border-slate-200 text-slate-700 shadow-md"
									autofocus
								/>
							</div>
							<div class="flex gap-1">
								<button
									class="rounded bg-teal-50 px-2 py-1 text-teal-700 md:bg-teal-700 md:text-teal-50"
									>add</button
								>
								<button
									type="button"
									class="px-2 py-1"
									onclick={() => ((adding = false), (new_username = ""))}>cancel</button
								>
							</div>
						</form>
					{/if}
				</li>
			</ul>
			<hr class="mx-auto w-5/6 rounded border border-teal-600 p-0 md:hidden" />
		</div>
	{/if}
</div>
