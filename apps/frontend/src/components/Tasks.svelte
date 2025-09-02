<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import * as Select from '$lib/components/ui/select/index';
	type TaskType = 'deep' | 'shallow';
	type Task = { id: number; title: string; type: TaskType; completed: boolean };

	let tasks: Task[] = [];
	let newTaskTitle = '';
	let newTaskType: TaskType = 'deep';

	function addTask() {
		const title = newTaskTitle.trim();
		if (!title) return;
		tasks = [...tasks, { id: Date.now(), title, type: newTaskType, completed: false }];
		newTaskTitle = '';
	}

	function removeTask(id: number) {
		tasks = tasks.filter((t) => t.id !== id);
	}

	function toggleComplete(id: number) {
		tasks = tasks.map((t) => (t.id === id ? { ...t, completed: !t.completed } : t));
	}

	function orderedByCompletion(items: Task[]): Task[] {
		// Incomplete first, then completed; keep stable order otherwise
		return [...items].sort((a, b) => (a.completed === b.completed ? 0 : a.completed ? 1 : -1));
	}
</script>

<div class="w-full">
	<div
		class="mx-auto rounded-md"
		style="background-color:#bf5a5a; padding:18px 18px 22px; max-width:640px;"
	>
		<div class="flex flex-wrap items-center gap-2 sm:gap-3">
			<input
				class="task-input"
				placeholder="Brain dump a task..."
				bind:value={newTaskTitle}
				onkeydown={(e) => e.key === 'Enter' && addTask()}
			/>

			<Select.Root type="single" bind:value={newTaskType}>
				<Select.Trigger class="bg-[#a64646] text-white border-none !h-[44px] capitalize">{newTaskType}</Select.Trigger>
				<Select.Content>
					<Select.Item value="deep">Deep</Select.Item>
					<Select.Item value="shallow">Shallow</Select.Item>
				</Select.Content>
			</Select.Root>
			<Button variant="primary" class="!h-[44px] !px-10" onclick={addTask}>Add</Button>
		</div>

		<div class="mt-4 grid gap-4 sm:mt-5 sm:grid-cols-2 sm:gap-6">
			<div>
				<div class="list-title">Deep</div>
				<ul class="list">
					{#each orderedByCompletion(tasks.filter((t) => t.type === 'deep')) as t (t.id)}
						<li class="item" class:completed={t.completed}>
							<Button
								class="!h-[32px] !w-[32px] !min-w-[32px] !min-h-[32px] bg-white text-[#a64646] hover:bg-white rounded-lg text-lg"
								aria-label="toggle complete"
								aria-pressed={t.completed}
								onclick={() => toggleComplete(t.id)}>{t.completed ? '✓' : ''}</Button
							>
							<span class="title">{t.title}</span>
							<div class="actions">
								<Button class="x" aria-label="remove" onclick={() => removeTask(t.id)}>×</Button>
							</div>
						</li>
					{/each}
					{#if tasks.filter((t) => t.type === 'deep').length === 0}
						<li class="item empty">No deep tasks yet</li>
					{/if}
				</ul>
			</div>
			<div>
				<div class="list-title">Shallow</div>
				<ul class="list">
					{#each orderedByCompletion(tasks.filter((t) => t.type === 'shallow')) as t (t.id)}
						<li class="item" class:completed={t.completed}>
							<Button
								class="!h-[32px] !w-[32px] !min-w-[32px] !min-h-[32px] bg-white text-[#a64646] hover:bg-white rounded-lg text-lg"
								aria-label="toggle complete"
								aria-pressed={t.completed}
								onclick={() => toggleComplete(t.id)}>{t.completed ? '✓' : ''}</Button
							>
							<span class="title">{t.title}</span>
							<div class="actions">
								<Button class="x" aria-label="remove" onclick={() => removeTask(t.id)}>×</Button>
							</div>
						</li>
					{/each}
					{#if tasks.filter((t) => t.type === 'shallow').length === 0}
						<li class="item empty">No shallow tasks yet</li>
					{/if}
				</ul>
			</div>
		</div>
	</div>
</div>

<style>
	.task-input {
		flex: 1 1 160px;
		min-width: 0;
		background: #fff;
		color: #5b2727;
		padding: 10px 12px;
		border-radius: 8px;
		outline: none;
		border: 0;
	}
	:global(.task-add) {
		background: #fff;
		color: #a64646;
		font-weight: 700;
		padding: 10px 16px;
		border-radius: 8px;
		box-shadow: 0 2px 0 #c84f4f;
	}

	.list-title {
		color: #fff;
		font-weight: 700;
		margin-bottom: 8px;
	}
	.list {
		display: flex;
		flex-direction: column;
		gap: 8px;
	}
	.item {
		display: flex;
		align-items: center;
		gap: 10px;
		background: rgba(255, 255, 255, 0.1);
		color: #fff;
		padding: 10px 12px;
		border-radius: 8px;
	}
	.item.completed {
		opacity: 0.5;
	}
	.item.completed .title {
		text-decoration: line-through;
	}
	.item.empty {
		opacity: 0.7;
		justify-content: center;
	}
	.title {
		flex: 1 1 auto;
		min-width: 0;
		word-break: break-word;
		overflow-wrap: anywhere;
	}
	.actions {
		margin-left: auto;
		display: flex;
		align-items: center;
		gap: 8px;
	}
	:global(.x) {
		background: transparent;
		color: #fff;
		opacity: 0.8;
		font-size: 18px;
		line-height: 1;
		padding: 0 4px;
	}
</style>
