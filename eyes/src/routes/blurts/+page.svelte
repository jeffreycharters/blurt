<script lang="ts">
	import { onMount } from "svelte"

	import type { Blurt as BlurtType } from "../../app"
	import Blurt from "$lib/c/Blurt.svelte"
	import Loader from "$lib/c/Loader.svelte"
	import Processing from "$lib/c/Processing.svelte"

	import type { PageData } from "./$types"
	import { crossfade, fade } from "svelte/transition"
	import { flip } from "svelte/animate"
	import { toast, Toaster } from "svelte-sonner"
	import { env } from "$env/dynamic/public"
	import { cubicInOut } from "svelte/easing"

	export let data: PageData
	export let userID: string | undefined = data.userID ?? "-1"
	export let blurts = data.blurts ?? []

	let blurt = ""
	let initiating = true // loading initial Blurts
	let loading = false // loading new set of Blurts via "load more" click"

	let countdownBox: HTMLDivElement
	let loadBlurtBox: HTMLDivElement

	let loadableBlurts = true
	const loadCount = 25 // how mant blurts to load at a time
	let loadOffset = loadCount // how many blurts have been loaded

	let userCount = 1

	let conn: WebSocket | null
	let reconnectDelay: number
	let status = false

	function websocketConnect() {
		conn = new WebSocket(`${env.PUBLIC_WS_ADDRESS}/ws`)

		conn.onopen = function () {
			reconnectDelay = 100
			status = true
		}

		conn.onclose = function () {
			status = false
			if (conn) conn = null
			toast.error("Unconnected from BlurtHQ! Reconnecting...")

			setTimeout(() => {
				websocketConnect()
				reconnectDelay = reconnectDelay * 2 ?? 100
			}, reconnectDelay)
		}

		conn.onmessage = function (evt) {
			const data = JSON.parse(evt.data)
			if (!data.success) {
				return toast.error(data.error)
			}

			if (data.payload.blurt) {
				return (blurts = [data.payload.blurt, ...blurts])
			}

			if (data.payload.lik) {
				const likdBlurt = blurts.find((blurt) => blurt.id === data.payload.lik.blurt_id)
				if (!likdBlurt) return toast.error("Likd blurt not found.")

				likdBlurt.edges.liks = [...(likdBlurt.edges.liks ?? []), data.payload.lik]

				const likdBlurtSpot = blurts.findIndex((blurt) => blurt.id === data.payload.lik.blurt_id)
				blurts = blurts.toSpliced(likdBlurtSpot, 1, likdBlurt)
			}

			if (data.payload.userCount) {
				userCount = data.payload.userCount
			}
		}
	}

	const likHandler = (blurt: BlurtType) => {
		if (!conn) return false

		conn.send(
			JSON.stringify({
				type: "lik",
				payload: { blurtID: blurt.id, userID }
			})
		)

		return true
	}

	const [send, receive] = crossfade({ duration: 200 })
	setTimeout(() => (initiating = false), 350)

	let processing = false // processing adding new Blurt

	onMount(async () => {
		const observerOptions = { root: null, rootMargin: "0px" }
		const intersectionObserver = new IntersectionObserver(loadHandler, observerOptions)
		intersectionObserver.observe(loadBlurtBox)

		websocketConnect()
	})

	function typeHandler() {
		const remaining = 14 - blurt.length
		const hue = (360 - 190) / remaining + 190
		const saturation = (100 - 60) / remaining + 60

		countdownBox.style.color = `hsl(${hue}, ${saturation}%, 50%)`
		countdownBox.innerHTML = String(remaining)
	}

	async function submitHandler() {
		if (!conn) return

		conn.send(
			JSON.stringify({
				type: "blurt",
				payload: { content: blurt, userID }
			})
		)

		blurt = ""
		typeHandler()
	}

	async function loadHandler(entries: IntersectionObserverEntry[]) {
		if (entries[0].isIntersecting) {
			loading = true

			let res: Response
			try {
				res = await fetch(`api/blurts?offset=${loadOffset}&count=${loadCount}`)
			} catch (err) {
				console.error(err)
				return toast.error("Failed to get mor blurts!")
			}

			if (!res.ok) return toast.error("Failed to get mor blurts!")

			setTimeout(async () => {
				const newBlurts = (await res.json()) as BlurtType[]
				loadOffset += newBlurts.length
				if (newBlurts.length === 0) {
					loadableBlurts = false
				}
				blurts = [...blurts, ...newBlurts]
				loading = false
			}, 750)
		}
	}
</script>

<Toaster expand={true} richColors />

{#if initiating}
	<div
		out:fade|global={{ duration: 200 }}
		class="absolute -top-2 bottom-0 left-0 right-0 z-50 bg-white pt-8 text-center"
	>
		<Processing text="Initiating! Beep Boop..." />
	</div>
{/if}

<div class="mx-auto my-2 max-w-md px-2">
	<form class="flex items-baseline justify-between" on:submit|preventDefault={submitHandler}>
		<label for="blurt" class="hidden">blurt</label>
		<div class="flex items-baseline gap-2">
			<input
				class="w-40 rounded-md border border-solid border-teal-900 bg-teal-50 px-2 py-1 font-semibold text-teal-800"
				type="text"
				bind:value={blurt}
				id="blurt"
				placeholder="blurt here"
				autocomplete="off"
				maxlength="14"
				style="will-change: transform;"
				on:keyup={typeHandler}
			/>
			<div class="font-bold" style="color: hsl(190, 60%, 55%)" bind:this={countdownBox}>14</div>
		</div>

		<button
			type="submit"
			id="blurt-button"
			class="rounded-md border border-solid border-gray-900 bg-teal-600 px-4 py-1 font-semibold text-white"
			>blurt it</button
		>
	</form>

	<div
		class="mt-2 flex w-full items-center justify-evenly gap-1 rounded bg-slate-100 p-2 text-sm font-bold text-slate-500"
	>
		<div>{userCount} chillin'</div>
		<div class="flex items-center gap-1">
			{#if status}
				<div class="mt-[1px] h-2 w-2 rounded-full border border-slate-400 bg-green-500" />
				<span>cunnekted</span>
			{:else}
				<div class="flex items-center gap-1 rounded bg-slate-200 px-2 py-1">
					<div class="mt-[1px] h-2 w-2 rounded-full border border-slate-400 bg-red-500" />
					<button on:click={websocketConnect}>rekinect!</button>
				</div>
			{/if}
		</div>
	</div>

	{#if processing}
		<Processing text="Blurting!" />
	{/if}

	<main id="blurt-zone">
		{#each blurts as blurt (blurt.id)}
			<div
				in:receive|global={{ key: blurt.id }}
				out:send|global={{ key: blurt.id }}
				animate:flip={{ duration: 400, easing: cubicInOut }}
			>
				<Blurt {blurt} {likHandler} {userID} />
			</div>
		{:else}
			<Blurt
				{userID}
				{likHandler}
				blurt={{
					id: "1",
					author: "BlurtBot",
					create_time: new Date("2022-01-01 01:01:01"),
					content: "Blurtin",
					liks: 420,
					userLikd: false,
					usersBlurt: false,
					edges: {
						author: {
							id: "1",
							username: "BlurtBot",
							create_time: new Date("2022-01-01 01:01:01")
						},
						liks: undefined
					}
				}}
			/>
		{/each}
	</main>

	{#if loadableBlurts}
		<div bind:this={loadBlurtBox} class="h-4" />
	{:else}
		<div class="flex justify-center gap-2 font-bold text-teal-600">
			<img src="cactus.svg" alt="cactus" class="relative inline h-16" />
			<div class="flex flex-col items-start justify-center">
				<div>No more blurts!</div>
				<div>You win!</div>
			</div>
		</div>
	{/if}
	<Loader showing={loading} />
</div>
