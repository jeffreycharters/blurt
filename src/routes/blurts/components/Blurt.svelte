<article class="mt-2 mb-4">
	<h2 class="font-semibold text-teal-600 text-4xl" style="font-family: 'Caveat Brush'">
		<slot name="blurt">
			<span>No blurt!</span>
		</slot>
	</h2>
	<div class="flex justify-between">
		<div>
			<slot name="username">
				<span>Anonymous</span>
			</slot>
		</div>
		<div>
			<slot name="date">
				<span>Timeless</span>
			</slot>
		</div>
	</div>
</article>
