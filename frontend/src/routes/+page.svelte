<script lang="ts">
	import { BlurtList, Footer, NewBlurtForm, PersonaManager } from "$lib/components"
</script>

<div
	class="mx-auto my-2 flex w-full flex-col items-center px-2 md:max-w-screen-md md:flex-row md:gap-4"
>
	<PersonaManager />
	<div class="flex w-full flex-col gap-4">
		<NewBlurtForm />
		<BlurtList />
		<Footer />
	</div>
</div>
