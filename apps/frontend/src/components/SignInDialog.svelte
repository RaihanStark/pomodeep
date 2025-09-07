<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import { signInSchema, signUpSchema } from '$lib/utils/validation';
	import { signUp } from '$lib/utils/api';

	let email = '';
	let password = '';
	let isSignUp = false;
	let errors: Record<string, string> = {};
	let isLoading = false;
	let successMessage = '';

	async function handleSubmit() {
		// Clear previous errors and success message
		errors = {};
		successMessage = '';

		// Validate form data
		const schema = isSignUp ? signUpSchema : signInSchema;
		const result = schema.safeParse({ email, password });

		if (!result.success) {
			// Set validation errors
			result.error.issues.forEach((issue) => {
				if (issue.path[0]) {
					errors[issue.path[0] as string] = issue.message;
				}
			});
			return;
		}

		// Form is valid, proceed with submission
		if (isSignUp) {
			await handleSignUp(result.data);
		} else {
			// Handle sign in logic
			console.log('Sign in:', result.data);
		}
	}

	async function handleSignUp(data: { email: string; password: string }) {
		isLoading = true;
		try {
			const response = await signUp(data);
			successMessage = `Account created successfully! Welcome, ${response.user.email}`;

			// Reset form after successful signup
			email = '';
			password = '';
			errors = {};

			// Switch to sign in mode after successful signup
			setTimeout(() => {
				isSignUp = false;
				successMessage = '';
			}, 3000);
		} catch (error) {
			errors.general = error instanceof Error ? error.message : 'Sign up failed';
		} finally {
			isLoading = false;
		}
	}

	function toggleMode() {
		isSignUp = !isSignUp;
		email = '';
		password = '';
		errors = {};
		successMessage = '';
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
			{#if successMessage}
				<div class="rounded-md border border-green-200 bg-green-50 p-3">
					<p class="text-sm text-green-600">{successMessage}</p>
				</div>
			{/if}

			{#if errors.general}
				<div class="rounded-md border border-red-200 bg-red-50 p-3">
					<p class="text-sm text-red-600">{errors.general}</p>
				</div>
			{/if}

			<div class="space-y-2">
				<Label for="email">Email</Label>
				<Input
					id="email"
					type="email"
					placeholder="Enter your email"
					bind:value={email}
					class={errors.email ? 'border-red-500' : ''}
					disabled={isLoading}
				/>
				{#if errors.email}
					<p class="text-sm text-red-500">{errors.email}</p>
				{/if}
			</div>

			<div class="space-y-2">
				<Label for="password">Password</Label>
				<Input
					id="password"
					type="password"
					placeholder="Enter your password"
					bind:value={password}
					class={errors.password ? 'border-red-500' : ''}
					disabled={isLoading}
				/>
				{#if errors.password}
					<p class="text-sm text-red-500">{errors.password}</p>
				{/if}
			</div>

			<Button type="submit" variant="tab" class="w-full" disabled={isLoading}>
				{#if isLoading}
					{isSignUp ? 'Creating Account...' : 'Signing In...'}
				{:else}
					{isSignUp ? 'Create Account' : 'Sign In'}
				{/if}
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
