<script>
	import { onDestroy, onMount } from 'svelte';
	import { page } from '$app/stores';
	import Blurt from './components/Blurt.svelte';
	import { blurts } from '$lib/stores/blurts';
	import { humanizeDates } from './utils';
	import { browser } from '$app/env';

	import { crossfade, fade } from 'svelte/transition';
	import { flip } from 'svelte/animate';
	import Loader from './components/Loader.svelte';
	import Processing from './components/Processing.svelte';
	const [send, receive] = crossfade({ duration: 200 });

	let blurt = '';
	let username = '';
	let initiating = true; // loading initial Blurts
	let loading = false; // loading new set of Blurts via infinite scroll

	let processing = false; // processing adding new Blurt

	// need an id for this interval so we can remove it on destroy.
	let fetchInterval = '';
	let likInterval = '';

	$: displayBlurts = $blurts;

	const emptyBlurt = {
		uid: 'howyagoinmayt',
		user: {
			username: 'Blurtbot 🤖',
			id: 1
		},
		blurt: 'No blurts yet!',
		userLikd: true,
		created_at: new Date().toISOString(),
		_count: {
			liks: Math.floor(Math.random() * 1000)
		}
	};

	onMount(async () => {
		username = localStorage.getItem('username');
		if (!username) location.href = '/';
		const res = await fetch(`/blurts.json?username=${encodeURIComponent(username)}`);
		let rawBlurts = await res.json();
		if (rawBlurts.length === 0) rawBlurts = [emptyBlurt];
		const dateBlurts = humanizeDates(rawBlurts);
		blurts.set(dateBlurts);
		setTimeout(() => (initiating = false), 500);
		if (browser) {
			fetchInterval = setInterval(getBlurts, 2000);
			likInterval = setInterval(updateLiks, 10000);
		}
	});

	onDestroy(async () => {
		clearInterval(fetchInterval);
		clearInterval(likInterval);
	});

	const path = $page.url.origin;

	const getBlurts = async () => {
		const mostRecent = displayBlurts[0].created_at;
		const url = `${path}/blurts.json?username=${encodeURIComponent(username)}&after=${mostRecent}`;
		const res = await fetch(url);
		if (res.ok) {
			const rawBlurts = await res.json();
			if (rawBlurts.length > 0) {
				const dateBlurts = humanizeDates(rawBlurts);
				blurts.set([...dateBlurts, ...displayBlurts]);
			}
		}
	};

	const updateLiks = async () => {
		const newestBlurt = encodeURIComponent(displayBlurts[0].created_at);
		const oldestBlurt = encodeURIComponent(displayBlurts[displayBlurts.length - 1].created_at);
		const url = `${path}/blurts/lik.json?from=${oldestBlurt}&to=${newestBlurt}`;
		const res = await fetch(url);
		const blurtLiks = await res.json();
		const updatedBlurts = displayBlurts.map((b) => {
			const blurtLik = blurtLiks.find((l) => l.uid === b.uid);
			if (!blurtLik) return b;
			return {
				...b,
				_count: {
					liks: blurtLik._count.liks
				}
			};
		});
		blurts.set(updatedBlurts);
	};

	const typeHandler = () => {
		const remaining = 14 - blurt.length;
		const hue = (360 - 190) / remaining + 190;
		const saturation = (100 - 60) / remaining + 60;
		const countdownDiv = document.getElementById('countdown-box');
		countdownDiv.style.color = `hsl(${hue}, ${saturation}%, 50%)`;
		countdownDiv.innerHTML = remaining;
	};

	const submitHandler = async () => {
		const processTimer = setTimeout(() => (processing = true), 500);
		const res = await fetch('/blurts.json', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Accept: 'application/json'
			},
			body: JSON.stringify({ username: username, blurt: blurt })
		});
		if (!res.ok) {
			console.log('error saving blurt');
			processing = false;
			return;
		}
		const newBlurt = await res.json();
		newBlurt.liks = [];
		clearTimeout(processTimer);
		processing = false;
		blurts.set([newBlurt, ...displayBlurts]);
		blurt = '';
		const countdownDiv = document.getElementById('countdown-box');
		countdownDiv.style.color = `hsl(190, 60%, 50%)`;
		countdownDiv.textContent = '14';
	};

	const getElementParentWithClassname = (el, className) => {
		while (!el.classList.contains(className)) {
			el = el.parentNode;
		}
		return el;
	};

	const touchHandler = (e) => {
		let blurtDiv = e.target;
		const parentDiv = getElementParentWithClassname(blurtDiv, 'blurt-holder');
		parentDiv.style.transform = 'scale(0.98)';
	};

	const touchEndHandler = (e) => {
		let blurtDiv = e.target;
		const parentDiv = getElementParentWithClassname(blurtDiv, 'blurt-holder');
		parentDiv.style.transform = 'scale(1)';
		parentDiv.style.boxShadow = '';
	};

	// Function is called to load next set of Blurts by intersection observer.
	const loadMoreBlurts = async (entries) => {
		entries.forEach(async (e) => {
			if (e.isIntersecting) {
				const loadTimer = setTimeout(() => (loading = true), 500);
				const path = $page.url.origin;
				const count = 25;
				const url = `${path}/blurts.json?username=${encodeURIComponent(
					username
				)}&take=${count}&cursor=${displayBlurts[displayBlurts.length - 1].uid}`;
				const res = await fetch(url);
				if (res.status === 404) {
					document.getElementById('after-blurt').remove();
					return;
				}
				const rawBlurts = await res.json();
				const dateBlurts = humanizeDates(rawBlurts);
				blurts.set([...displayBlurts, ...dateBlurts]);
				clearTimeout(loadTimer);
				loading = false;
			}
		});
	};

	if (browser) {
		setTimeout(() => {
			const options = {
				root: null,
				rootMargin: '0px 0px 0px 0px',
				threshold: 1.0
			};
			let observer = new IntersectionObserver(loadMoreBlurts, options);
			let target = document.getElementById('after-blurt');
			observer.observe(target);
		}, 1000);
	}
</script>

{#if initiating}
	<div
		out:fade|global={{ duration: 200 }}
		class="bg-white absolute -top-2 bottom-0 left-0 right-0 z-50 text-center pt-8"
	>
		<Processing text="Initiating! Beep Boop..." />
	</div>
{/if}

<div class="max-w-md mx-auto my-2 px-2">
	<form class="flex justify-between items-baseline" on:submit|preventDefault={submitHandler}>
		<label for="blurt" class="hidden">blurt</label>
		<div class="flex items-baseline gap-2">
			<input
				class="py-1 px-2 rounded-md border border-solid border-teal-900 bg-teal-50 text-teal-800 font-semibold w-40"
				type="text"
				bind:value={blurt}
				id="blurt"
				placeholder="blurt here"
				autocomplete="off"
				maxlength="14"
				style="will-change: transform;"
				on:keyup={typeHandler}
			/>
			<div class="font-bold" style="color: hsl(190, 60%, 55%)" id="countdown-box">14</div>
		</div>
		<button
			type="submit"
			id="blurt-button"
			class="border border-solid border-gray-900 py-1 px-4 rounded-md font-semibold text-white bg-teal-600"
			>blurt it</button
		>
	</form>

	{#if processing}
		<Processing text="Blurting!" />
	{/if}

	<main id="blurt-zone">
		{#each displayBlurts as blurt (blurt.uid)}
			<div
				on:touchstart={touchHandler}
				on:touchend={touchEndHandler}
				in:receive|global={{ key: blurt.id }}
				out:send|global={{ key: blurt.uid }}
				animate:flip={{ duration: 400 }}
				class="blurt-holder"
			>
				<Blurt {blurt} />
			</div>
		{/each}
		<Loader showing={loading} />
	</main>
	<div id="after-blurt" />
</div>

<style>
	.blurt-holder {
		transition: all 0.075s cubic-bezier(0.17, 0.67, 0.88, 4.06);
	}
</style>
