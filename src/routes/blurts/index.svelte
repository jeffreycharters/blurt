<script>
	import { onDestroy, onMount } from 'svelte';
	import { page } from '$app/stores';
	import Blurt from './components/Blurt.svelte';
	import { blurts } from '$lib/stores/blurts';
	import { humanizeDates } from './utils';
	import { browser } from '$app/env';

	import { crossfade, fly } from 'svelte/transition';
	import { flip } from 'svelte/animate';
	import Loader from './components/Loader.svelte';
	const [send, receive] = crossfade({ duration: 200 });

	let blurt = '';
	let username = '';
	let loading = false;

	// need an id for this interval so we can remove it on destroy.
	let fetchInterval = '';

	$: displayBlurts = $blurts;

	onMount(async () => {
		username = localStorage.getItem('username');
		if (!username) location.href = '/';
		const res = await fetch('/blurts.json');
		const rawBlurts = await res.json();
		const dateBlurts = humanizeDates(rawBlurts);
		blurts.set(dateBlurts);
		if (browser) fetchInterval = setInterval(getBlurts, 2000);
	});

	onDestroy(async () => {
		clearInterval(fetchInterval);
	});

	const getBlurts = async () => {
		const path = $page.url.origin;
		const mostRecent = displayBlurts[0].created_at;
		const url = `${path}/blurts.json?after=${mostRecent}`;
		const res = await fetch(url);
		if (res.ok) {
			const rawBlurts = await res.json();
			if (rawBlurts.length > 0) {
				const dateBlurts = humanizeDates(rawBlurts);
				blurts.set([...dateBlurts, ...displayBlurts]);
			}
		}
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
			return;
		}
		const newBlurt = await res.json();
		newBlurt.liks = [];
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
				console.log('loading');
				loading = true;
				const path = $page.url.origin;
				const count = 25;
				const url = `${path}/blurts.json?take=${count}&cursor=${
					displayBlurts[displayBlurts.length - 1].uid
				}`;
				console.log(url);
				const res = await fetch(url);
				if (res.status === 404) {
					document.getElementById('after-blurt').remove();
					return;
				}
				const rawBlurts = await res.json();
				const dateBlurts = humanizeDates(rawBlurts);
				blurts.set([...displayBlurts, ...dateBlurts]);
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

<div class="max-w-md mx-auto">
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

	<main id="blurt-zone">
		{#each displayBlurts as blurt (blurt.uid)}
			<div
				on:touchstart={touchHandler}
				on:touchend={touchEndHandler}
				in:receive={{ key: blurt.id }}
				out:send={{ key: blurt.uid }}
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
