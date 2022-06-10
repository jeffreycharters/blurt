<script>
	import { fly } from 'svelte/transition';
	export let header;
	export let text;
	let flexClass = header ? 'flex-col gap-2' : 'flex-row';
</script>

<div
	in:fly={{ y: -20, duration: 200 }}
	class="flex {flexClass} items-center gap-4 mx-auto w-fit my-4 border border-gray-200 px-12 py-2 rounded-md text-teal-500 shadow-sm w-xs"
>
	{#if header}
		<div>
			<strong>{header}</strong>
		</div>
	{/if}

	<svg
		width="40"
		height="10"
		viewBox="0 0 120 30"
		xmlns="http://www.w3.org/2000/svg"
		fill="currentColor"
	>
		<circle cx="15" cy="15" r="15">
			<animate
				attributeName="r"
				from="15"
				to="15"
				begin="0s"
				dur="0.8s"
				values="15;9;15"
				calcMode="linear"
				repeatCount="indefinite"
			/>
			<animate
				attributeName="fill-opacity"
				from="1"
				to="1"
				begin="0s"
				dur="0.8s"
				values="1;.5;1"
				calcMode="linear"
				repeatCount="indefinite"
			/>
		</circle>
		<circle cx="60" cy="15" r="9" fill-opacity="0.3">
			<animate
				attributeName="r"
				from="9"
				to="9"
				begin="0s"
				dur="0.8s"
				values="9;15;9"
				calcMode="linear"
				repeatCount="indefinite"
			/>
			<animate
				attributeName="fill-opacity"
				from="0.5"
				to="0.5"
				begin="0s"
				dur="0.8s"
				values=".5;1;.5"
				calcMode="linear"
				repeatCount="indefinite"
			/>
		</circle>
		<circle cx="105" cy="15" r="15">
			<animate
				attributeName="r"
				from="15"
				to="15"
				begin="0s"
				dur="0.8s"
				values="15;9;15"
				calcMode="linear"
				repeatCount="indefinite"
			/>
			<animate
				attributeName="fill-opacity"
				from="1"
				to="1"
				begin="0s"
				dur="0.8s"
				values="1;.5;1"
				calcMode="linear"
				repeatCount="indefinite"
			/>
		</circle>
	</svg>

	<div class="whitespace-nowrap">{text}</div>
</div>
