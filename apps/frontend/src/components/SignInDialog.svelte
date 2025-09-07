<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import { signInSchema, signUpSchema } from '$lib/utils/validation';
	import { signUp, signIn, setAuthToken } from '$lib/utils/api';
	import { authStore } from '$lib/stores/auth';

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
			await handleSignIn(result.data);
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
			}, 2000);
		} catch (error) {
			errors.general = error instanceof Error ? error.message : 'Sign up failed';
		} finally {
			isLoading = false;
		}
	}

	async function handleSignIn(data: { email: string; password: string }) {
		isLoading = true;
		try {
			const response = await signIn(data);

			// Store the token
			setAuthToken(response.token);

			// Update auth store
			authStore.login(response.user, response.token);

			successMessage = `Welcome back, ${response.user.email}!`;

			// Reset form after successful signin
			email = '';
			password = '';
			errors = {};

			// Close dialog after successful signin
			setTimeout(() => {
				successMessage = '';
				// You can emit an event here to notify parent component
				// or redirect to dashboard
			}, 1500);
		} catch (error) {
			errors.general = error instanceof Error ? error.message : 'Sign in failed';
		} finally {
			isLoading = false;
		}
	}

	function toggleMode() {
		isSignUp = !isSignUp;
		errors = {};
		successMessage = '';
		email = '';
		password = '';
	}
</script>

<Dialog.Root>
	<Dialog.Trigger>
		<Button variant="outline">
			{isSignUp ? 'Sign Up' : 'Sign In'}
		</Button>
	</Dialog.Trigger>
	<Dialog.Content class="sm:max-w-[425px]">
		<Dialog.Header>
			<Dialog.Title>{isSignUp ? 'Create Account' : 'Welcome Back'}</Dialog.Title>
			<Dialog.Description>
				{isSignUp
					? 'Enter your details to create a new account'
					: 'Enter your credentials to sign in to your account'}
			</Dialog.Description>
		</Dialog.Header>

		<form on:submit|preventDefault={handleSubmit} class="space-y-4">
			<div class="space-y-2">
				<Label for="email">Email</Label>
				<Input
					id="email"
					type="email"
					bind:value={email}
					placeholder="Enter your email"
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
					bind:value={password}
					placeholder="Enter your password"
					disabled={isLoading}
				/>
				{#if errors.password}
					<p class="text-sm text-red-500">{errors.password}</p>
				{/if}
			</div>

			{#if errors.general}
				<div class="rounded-md border border-red-200 bg-red-50 p-3">
					<p class="text-sm text-red-600">{errors.general}</p>
				</div>
			{/if}

			{#if successMessage}
				<div class="rounded-md border border-green-200 bg-green-50 p-3">
					<p class="text-sm text-green-600">{successMessage}</p>
				</div>
			{/if}

			<Dialog.Footer>
				<Button type="submit" disabled={isLoading} class="w-full">
					{isLoading ? 'Please wait...' : isSignUp ? 'Create Account' : 'Sign In'}
				</Button>
			</Dialog.Footer>
		</form>

		<div class="text-center text-sm text-gray-600">
			{isSignUp ? 'Already have an account?' : "Don't have an account?"}
			<button
				type="button"
				on:click={toggleMode}
				class="ml-1 text-blue-600 underline hover:text-blue-800"
				disabled={isLoading}
			>
				{isSignUp ? 'Sign In' : 'Sign Up'}
			</button>
		</div>
	</Dialog.Content>
</Dialog.Root>
