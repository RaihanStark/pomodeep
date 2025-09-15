<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import * as Dialog from '$lib/components/ui/dialog';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import * as Select from '$lib/components/ui/select/index';
	import * as Tooltip from '$lib/components/ui/tooltip/index';
	type TaskType = 'deep' | 'shallow';
	type Task = { id: number; title: string; type: TaskType; completed: boolean };

	let tasks: Task[] = [];
	let currentTaskId: number | null = null;
	let newTaskTitle = '';
	let newTaskType: TaskType = 'deep';
	let isCreateOpen = false;

	function addTask() {
		const title = newTaskTitle.trim();
		if (!title) return;
		tasks = [...tasks, { id: Date.now(), title, type: newTaskType, completed: false }];
		newTaskTitle = '';
	}

	function handleCreateTask() {
		const title = newTaskTitle.trim();
		if (!title) return;
		addTask();
		isCreateOpen = false;
	}

	function removeTask(id: number) {
		tasks = tasks.filter((t) => t.id !== id);
		if (currentTaskId === id) currentTaskId = null;
	}

	function toggleComplete(id: number) {
		tasks = tasks.map((t) => (t.id === id ? { ...t, completed: !t.completed } : t));
	}

	function orderedByCompletion(items: Task[]): Task[] {
		// Incomplete first, then completed; keep stable order otherwise
		return [...items].sort((a, b) => (a.completed === b.completed ? 0 : a.completed ? 1 : -1));
	}

	function selectTask(id: number) {
		currentTaskId = id;
	}
</script>

<Tooltip.Provider delayDuration={150}>
	<div class="w-full">
		<div class="mx-auto rounded-md" style="background-color:#bf5a5a; padding:18px 18px 22px;">
			<div class="flex flex-wrap items-center gap-2 sm:gap-3">
				<Dialog.Root bind:open={isCreateOpen}>
					<Dialog.Trigger class="w-full">
						<Button variant="primary" class="!h-[44px] !px-10 !w-full">Add task</Button>
					</Dialog.Trigger>
					<Dialog.Content class="sm:max-w-[425px]">
						<Dialog.Header>
							<Dialog.Title>Create task</Dialog.Title>
							<Dialog.Description>Enter details for your new task</Dialog.Description>
						</Dialog.Header>
						<div class="space-y-4">
							<div class="space-y-2">
								<Label for="task-title">Title</Label>
								<Input
									id="task-title"
									placeholder="Brain dump a task..."
									bind:value={newTaskTitle}
								/>
							</div>
							<div class="space-y-2">
								<Label>Type</Label>
								<Select.Root type="single" bind:value={newTaskType}>
									<Select.Trigger class="!h-[44px] border-none bg-[#a64646] capitalize text-white w-full"
										>{newTaskType}</Select.Trigger
									>
									<Select.Content>
										<Select.Item value="deep">Deep</Select.Item>
										<Select.Item value="shallow">Shallow</Select.Item>
									</Select.Content>
								</Select.Root>
							</div>
							<Dialog.Footer>
								<Button class="w-full" variant="tab" onclick={handleCreateTask}>Add task</Button>
							</Dialog.Footer>
						</div>
					</Dialog.Content>
				</Dialog.Root>
			</div>

			<div class="list-title !mt-4">Current task</div>
			<div class="item" style="grid-column: 1 / -1;">
				{#if currentTaskId}
					{#each tasks as t (t.id)}
						{#if t.id === currentTaskId}
							<span class="title">{t.title}</span>
						{/if}
					{/each}
				{:else}
					<span class="title">No task selected</span>
				{/if}
			</div>
			<div class="mt-4 grid gap-4 sm:mt-5 sm:grid-cols-2 sm:gap-6">
				<div>
					<div class="list-title">
						Deep <Tooltip.Root
							><Tooltip.Trigger aria-label="Definitions" style="cursor:help;">ⓘ</Tooltip.Trigger
							><Tooltip.Content side="right" sideOffset={8}>
								<div style="max-width:320px; line-height:1.45;">
									<div style="font-weight:700; margin-bottom:6px;">🔑 Definitions</div>
									<div style="margin-bottom:6px;">Tasks that:</div>
									<ul style="margin-left:16px; margin-bottom:8px; list-style: disc;">
										<li>Require intense focus and concentration</li>
										<li>Push your cognitive abilities to their limits</li>
										<li>Are hard to replicate or automate</li>
										<li>Create high-value output and move your skills or projects forward</li>
									</ul>
									<p style="margin-bottom:10px;">
										Example: designing a new feature, writing a research paper, drafting a legal
										contract, debugging a complex system.
									</p>
								</div>
							</Tooltip.Content></Tooltip.Root
						>
					</div>
					<ul class="list">
						{#each orderedByCompletion(tasks.filter((t) => t.type === 'deep')) as t (t.id)}
							<li
								class="item"
								class:completed={t.completed}
								class:selected={currentTaskId === t.id}
							>
								<Button
									class="!h-[32px] !min-h-[32px] !w-[32px] !min-w-[32px] rounded-lg bg-white text-lg text-[#a64646] hover:bg-white"
									aria-label="toggle complete"
									aria-pressed={t.completed}
									onclick={(e) => {
										e.stopPropagation();
										toggleComplete(t.id);
									}}>{t.completed ? '✓' : ''}</Button
								>
								<button
									class="title-btn"
									type="button"
									aria-label="Select task"
									onclick={() => selectTask(t.id)}
								>
									<span class="title">{t.title}</span>
								</button>
								<div class="actions">
									<Button
										class="x"
										aria-label="remove"
										onclick={(e) => {
											e.stopPropagation();
											removeTask(t.id);
										}}>×</Button
									>
								</div>
							</li>
						{/each}
						{#if tasks.filter((t) => t.type === 'deep').length === 0}
							<li class="item empty">No deep tasks yet</li>
						{/if}
					</ul>
				</div>
				<div>
					<div class="list-title">
						Shallow <Tooltip.Root
							><Tooltip.Trigger aria-label="Definitions" style="cursor:help;">ⓘ</Tooltip.Trigger
							><Tooltip.Content side="right" sideOffset={8}>
								<div style="max-width:320px; line-height:1.45;">
									<div style="font-weight:700; margin-bottom:6px;">🔑 Definitions</div>
									<div style="margin-bottom:6px;">Tasks that:</div>
									<ul style="margin-left:16px; margin-bottom:8px; list-style: disc;">
										<li>Are logistically easy and often done while distracted</li>
										<li>Don’t require deep focus or unique skill</li>
										<li>Are often repetitive, low-value, or easily delegated</li>
										<li>Produce little lasting value</li>
									</ul>
									<p>
										Example: answering routine emails, filling forms, updating spreadsheets,
										checking Slack, attending most status meetings.
									</p>
								</div>
							</Tooltip.Content></Tooltip.Root
						>
					</div>
					<ul class="list">
						{#each orderedByCompletion(tasks.filter((t) => t.type === 'shallow')) as t (t.id)}
							<li
								class="item"
								class:completed={t.completed}
								class:selected={currentTaskId === t.id}
							>
								<Button
									class="!h-[32px] !min-h-[32px] !w-[32px] !min-w-[32px] rounded-lg bg-white text-lg text-[#a64646] hover:bg-white"
									aria-label="toggle complete"
									aria-pressed={t.completed}
									onclick={(e) => {
										e.stopPropagation();
										toggleComplete(t.id);
									}}>{t.completed ? '✓' : ''}</Button
								>
								<button
									class="title-btn"
									type="button"
									aria-label="Select task"
									onclick={() => selectTask(t.id)}
								>
									<span class="title">{t.title}</span>
								</button>
								<div class="actions">
									<Button
										class="x"
										aria-label="remove"
										onclick={(e) => {
											e.stopPropagation();
											removeTask(t.id);
										}}>×</Button
									>
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
</Tooltip.Provider>

<style>

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
		transition:
			outline-color 0.15s ease,
			outline-width 0.15s ease,
			background-color 0.15s ease;
		outline: 2px solid transparent;
	}
	.item:hover {
		background: rgba(255, 255, 255, 0.14);
	}
	.item.selected {
		outline-color: #fff;
		outline-width: 2px;
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
	.title-btn {
		flex: 1 1 auto;
		background: transparent;
		border: 0;
		padding: 0;
		text-align: left;
		color: inherit;
		cursor: pointer;
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
