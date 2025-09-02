<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<div class="min-h-screen w-full flex items-start justify-center" style="background-color:#bc5151;">
  <div class="w-full max-w-4xl px-4 sm:px-6 md:px-8 pt-8 md:pt-12">
    {@render children?.()}
  </div>
  
</div>
