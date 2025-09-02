<script lang="ts">
  import { Button } from '$lib/components/ui/button';
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
  <div class="mx-auto rounded-md" style="background-color:#bf5a5a; padding:18px 18px 22px; max-width:640px;">
    <div class="flex flex-wrap items-center gap-2 sm:gap-3">
      <input class="task-input" placeholder="Brain dump a task..." bind:value={newTaskTitle} onkeydown={(e) => e.key==='Enter' && addTask()} />
      <select class="task-select" bind:value={newTaskType}>
        <option value="deep">Deep</option>
        <option value="shallow">Shallow</option>
      </select>
      <Button class="task-add" onclick={addTask}>Add</Button>
    </div>

    <div class="grid gap-4 sm:gap-6 mt-4 sm:mt-5 sm:grid-cols-2">
      <div>
        <div class="list-title">Deep</div>
        <ul class="list">
          {#each orderedByCompletion(tasks.filter((t) => t.type === 'deep')) as t (t.id)}
            <li class="item" class:completed={t.completed}>
              <Button class="check" aria-label="toggle complete" aria-pressed={t.completed} onclick={() => toggleComplete(t.id)}>{t.completed ? '✓' : ''}</Button>
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
              <Button class="check" aria-label="toggle complete" aria-pressed={t.completed} onclick={() => toggleComplete(t.id)}>{t.completed ? '✓' : ''}</Button>
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
  .task-input { flex:1 1 160px; min-width:0; background:#fff; color:#5b2727; padding:10px 12px; border-radius:8px; outline:none; border:0; }
  .task-select { background:#a64646; color:#fff; padding:10px 10px; border-radius:8px; border:0; }
  :global(.task-add) { background:#fff; color:#a64646; font-weight:700; padding:10px 16px; border-radius:8px; box-shadow:0 2px 0 #c84f4f; }

  .list-title { color:#fff; font-weight:700; margin-bottom:8px; }
  .list { display:flex; flex-direction:column; gap:8px; }
  .item { display:flex; align-items:center; gap:10px; background:rgba(255,255,255,.1); color:#fff; padding:10px 12px; border-radius:8px; }
  .item.completed { opacity:.5; }
  .item.completed .title { text-decoration:line-through; }
  .item.empty { opacity:.7; justify-content:center; }
  .title { flex:1 1 auto; min-width:0; word-break:break-word; overflow-wrap:anywhere; }
  .actions { margin-left:auto; display:flex; align-items:center; gap:8px; }
  :global(.check) { width:26px; height:26px; min-width:26px; min-height:26px; display:inline-flex; align-items:center; justify-content:center; background:#fff; color:#a64646; border-radius:6px; font-weight:800; box-shadow:0 2px 0 #c84f4f; }
  :global(.x) { background:transparent; color:#fff; opacity:.8; font-size:18px; line-height:1; padding:0 4px; }
</style>


