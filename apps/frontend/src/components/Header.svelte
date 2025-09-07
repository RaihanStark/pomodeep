<script lang="ts">
	import SignInDialog from './SignInDialog.svelte';
	import { Button } from '$lib/components/ui/button';
	import { authStore } from '$lib/stores/auth';

	function handleLogout() {
		authStore.logout();
	}
</script>

<header class="w-full border-b border-white/20 bg-white/10 backdrop-blur-sm">
	<div class="mx-auto w-full max-w-4xl px-4 py-4 sm:px-6 md:px-8">
		<div class="flex items-center justify-between">
			<!-- Logo/App Name -->
			<div class="flex items-center space-x-2">
				<h1 class="text-2xl font-bold text-white">PomoApp</h1>
			</div>

			<!-- Sign in/Register Button or Logout -->
			<div class="flex items-center space-x-2">
				{#if $authStore.isAuthenticated}
					<Button variant="outline" onclick={handleLogout}>Logout</Button>
				{:else}
					<SignInDialog />
				{/if}
			</div>
		</div>
	</div>
</header>
