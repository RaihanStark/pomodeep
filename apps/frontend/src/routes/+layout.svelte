<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import Header from '../components/Header.svelte';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<div class="flex min-h-screen w-full flex-col" style="background-color:#bc5151;">
	<div class="flex justify-center">
		<div class="w-full max-w-3xl px-4 sm:px-6 md:px-8">
			<Header />
		</div>
	</div>
	<div class="flex flex-1 items-start justify-center">
		<div class="w-full max-w-3xl px-4 pt-6 sm:px-6 md:px-8">
			{@render children?.()}
		</div>
	</div>
</div>
