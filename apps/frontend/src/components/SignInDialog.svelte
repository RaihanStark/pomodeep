<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';

	let email = '';
	let password = '';
	let isSignUp = false;

	function handleSubmit() {
		if (isSignUp) {
			// Handle sign up logic
			console.log('Sign up:', { email, password });
		} else {
			// Handle sign in logic
			console.log('Sign in:', { email, password });
		}
		// Reset form
		email = '';
		password = '';
	}

	function toggleMode() {
		isSignUp = !isSignUp;
		email = '';
		password = '';
	}
</script>

<Dialog.Root>
	<Dialog.Trigger>
		<Button variant="primary" size="sm">Sign in</Button>
	</Dialog.Trigger>
	<Dialog.Content class="sm:max-w-md">
		<Dialog.Header>
			<Dialog.Title class="text-xl font-semibold">
				{isSignUp ? 'Create Account' : 'Sign In'}
			</Dialog.Title>
			<Dialog.Description>
				{isSignUp
					? 'Enter your details to create a new account'
					: 'Enter your credentials to sign in to your account'}
			</Dialog.Description>
		</Dialog.Header>

		<form on:submit|preventDefault={handleSubmit} class="space-y-4">
			<div class="space-y-2">
				<Label for="email">Email</Label>
				<Input id="email" type="email" placeholder="Enter your email" bind:value={email} required />
			</div>

			<div class="space-y-2">
				<Label for="password">Password</Label>
				<Input
					id="password"
					type="password"
					placeholder="Enter your password"
					bind:value={password}
					required
				/>
			</div>

			<Button type="submit" variant="tab" class="w-full">
				{isSignUp ? 'Create Account' : 'Sign In'}
			</Button>

			<div class="flex flex-col items-center">
				<span>
					{isSignUp ? 'Already have an account?' : "Don't have an account?"}
					<Button
						type="button"
						variant="link"
						class="inline p-0 align-baseline"
						onclick={toggleMode}
					>
						{isSignUp ? 'Sign in' : 'Sign up'}
					</Button>
				</span>
			</div>
		</form>
	</Dialog.Content>
</Dialog.Root>
