<script lang="ts">
  import { onDestroy } from 'svelte';
  import { formatSecondsAsMMSS } from '$lib/utils/time';

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
    }
  }

  function reset() {
    stop();
    remaining = MODE_TO_SECONDS[mode];
  }

  onDestroy(() => stop());
</script>

<div class="w-full">
  <div class="mx-auto rounded-md" style="background-color:#bf5a5a; padding:20px 18px 26px; max-width:640px;">
    <div class="flex flex-wrap justify-center items-center gap-3 sm:gap-4 mb-4 sm:mb-5">
      <button class="tab" class:active-tab={mode==='pomodoro'} onclick={() => setMode('pomodoro')}>Pomodoro</button>
      <button class="tab" class:active-tab={mode==='short'} onclick={() => setMode('short')}>Short Break</button>
      <button class="tab" class:active-tab={mode==='long'} onclick={() => setMode('long')}>Long Break</button>
    </div>

    <div class="timer">{formatSecondsAsMMSS(remaining)}</div>

    <div class="mt-6 sm:mt-7 flex justify-center">
      {#if !running}
        <button onclick={start} class="cta">START</button>
      {:else}
        <button onclick={stop} class="cta">PAUSE</button>
      {/if}
    </div>
  </div>
</div>

<style>
  .tab { background-color:#b65454; color:rgba(255,255,255,.9); padding:6px 10px; border-radius:6px; transition:background-color .2s ease, transform .12s ease; white-space:nowrap; }
  .tab:hover { background-color:#ad4f4f; }
  .tab:active { transform:translateY(1px); }
  .active-tab { background-color:#a64646 !important; }

  .timer { color:#fff; text-align:center; font-weight:800; font-size:clamp(56px, 20vw, 120px); line-height:1; letter-spacing:2px; margin-top:10px; user-select:none; transition:transform .14s ease; }

  .cta { background:#fff; color:#a64646; font-weight:700; letter-spacing:1px; padding:12px 40px; border-radius:8px; box-shadow:0 3px 0 #c84f4f; transition:transform .12s ease, box-shadow .12s ease, opacity .2s ease; }
  .cta:hover { opacity:.95; }
  .cta:active { transform:translateY(2px); box-shadow:0 1px 0 #c84f4f; }

  @media (min-width: 380px) {
    .cta { padding:13px 46px; }
  }

  @media (min-width: 640px) {
    .timer { margin-top:12px; }
    .cta { padding:14px 52px; }
  }

  /* tasks styles moved to Tasks.svelte */
</style>


