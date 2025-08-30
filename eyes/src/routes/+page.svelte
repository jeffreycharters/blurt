<script lang="ts">
	import { Toaster, toast } from "svelte-sonner"
	import type { PageData } from "./$types"
	import { slide } from "svelte/transition"

	import { dev } from "$app/environment"
	import SuperDebug, { superForm } from "sveltekit-superforms"

	export let data: PageData
	const { form, enhance, errors, constraints, message } = superForm(data.form)

	message.subscribe(() => {
		if ($message) toast.error($message)
	})
</script>

<body class="min-h-full min-w-full bg-teal-500 px-2 py-2">
	<Toaster expand={true} richColors />

	<div class="m-0 mx-auto flex h-96 max-w-xs flex-col items-center justify-between">
		<div class="mt-24 w-full max-w-xs rounded-md bg-teal-50 py-4 text-center">
			<h1 class="text-7xl text-teal-900" style="font-family: 'Permanent Marker';">BLURT</h1>

			<h2 class="text-xs text-teal-500">An exciting <em>new</em> way to shout at the world!</h2>
		</div>

		<div class="mt-24 w-full">
			<form action="" method="post" id="login-form" use:enhance>
				<div class="flex flex-col gap-1 text-center font-bold text-teal-50">
					<label for="username">Enter a name and let's blurt!</label>
					<input
						type="text"
						name="username"
						class="w-full max-w-xs rounded-md px-2 py-1 text-center text-teal-900"
						placeholder="blazer420 is taken, sorry"
						autocomplete="off"
						aria-invalid={$errors.username ? "true" : undefined}
						bind:value={$form.username}
						{...$constraints.username}
					/>

					{#if $errors.username}
						<div transition:slide={{ duration: 200 }} class="ml-2 mr-auto text-xs text-red-800">
							{$errors.username}
						</div>
					{/if}

					<button
						type="submit"
						class="my-1 w-full max-w-xs rounded-md bg-teal-900 px-2 py-1 text-sm font-semibold uppercase tracking-widest"
						style="font-family: 'Permanent Marker';">Let's Go!!</button
					>
				</div>
				{#if dev}
					<SuperDebug data={$form} />
				{/if}
			</form>
		</div>
	</div>
</body>
