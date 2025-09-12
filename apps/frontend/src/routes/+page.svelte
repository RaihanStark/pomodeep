<script lang="ts">
	import Pomodoro from '../components/Pomodoro.svelte';
	import Tasks from '../components/Tasks.svelte';
	import DistractionJot from '../components/DistractionJot.svelte';
</script>

<Pomodoro />

<div class="mt-6">
	<DistractionJot />
</div>

<div class="mt-6">
	<Tasks />
</div>
