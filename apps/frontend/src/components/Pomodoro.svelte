<script lang="ts">
	import { onDestroy } from 'svelte';
	import { formatSecondsAsMMSS } from '$lib/utils/time';
	import { Button } from '$lib/components/ui/button';

	type Mode = 'pomodoro' | 'short' | 'long';

	const MODE_TO_SECONDS: Record<Mode, number> = {
		pomodoro: 25 * 60,
		short: 5 * 60,
		long: 15 * 60
	};

	let mode: Mode = 'pomodoro';
	let remaining = MODE_TO_SECONDS[mode];
	let running = false;
	let intervalId: number | null = null;

	function setMode(newMode: Mode) {
		mode = newMode;
		reset();
	}

	function start() {
		if (running) return;
		running = true;
		const tick = () => {
			if (remaining <= 0) {
				stop();
				return;
			}
			remaining -= 1;
		};
		// run first tick after 1s so initial display stays intact
		intervalId = window.setInterval(tick, 1000);
	}

	function stop() {
		running = false;
		if (intervalId !== null) {
			clearInterval(intervalId);
			intervalId = null;
			remaining = MODE_TO_SECONDS[mode];
		}
	}

	function reset() {
		stop();
		remaining = MODE_TO_SECONDS[mode];
	}

	onDestroy(() => stop());
</script>

<div class="w-full">
	<div class="mx-auto rounded-md" style="background-color:#bf5a5a; padding:20px 18px 26px">
		<div class="mb-4 flex flex-wrap items-center justify-center gap-3 sm:mb-5 sm:gap-4">
			<Button
				variant="tab"
				class={`${mode === 'pomodoro' ? 'active-tab' : ''}`}
				onclick={() => setMode('pomodoro')}>Pomodoro</Button
			>
			<Button
				variant="tab"
				class={`${mode === 'short' ? 'active-tab' : ''}`}
				onclick={() => setMode('short')}>Short Break</Button
			>
			<Button
				variant="tab"
				class={`${mode === 'long' ? 'active-tab' : ''}`}
				onclick={() => setMode('long')}>Long Break</Button
			>
		</div>

		<div class="timer">{formatSecondsAsMMSS(remaining)}</div>

		<div class="mt-6 flex justify-center sm:mt-7">
			{#if !running}
				<Button onclick={start} variant="primary" class="px-10 py-6">START</Button>
			{:else}
				<Button onclick={stop} variant="primary" class="px-10 py-6">PAUSE</Button>
			{/if}
			<Button onclick={reset} variant="secondary" class="ml-5 px-10 py-6">RESET</Button>
		</div>
	</div>
</div>

<style>
	:global(.active-tab) {
		background-color: #a64646 !important;
	}

	.timer {
		color: #fff;
		text-align: center;
		font-weight: 800;
		font-size: clamp(56px, 20vw, 120px);
		line-height: 1;
		letter-spacing: 2px;
		margin-top: 10px;
		user-select: none;
		transition: transform 0.14s ease;
	}

	@media (min-width: 640px) {
		.timer {
			margin-top: 12px;
		}
	}
</style>
