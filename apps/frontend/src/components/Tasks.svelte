<script lang="ts">
  type TaskType = 'deep' | 'shallow';
  type Task = { id: number; title: string; type: TaskType };

  let tasks: Task[] = [];
  let newTaskTitle = '';
  let newTaskType: TaskType = 'deep';

  function addTask() {
    const title = newTaskTitle.trim();
    if (!title) return;
    tasks = [...tasks, { id: Date.now(), title, type: newTaskType }];
    newTaskTitle = '';
  }

  function removeTask(id: number) {
    tasks = tasks.filter((t) => t.id !== id);
  }
</script>

<div class="w-full">
  <div class="mx-auto rounded-md" style="background-color:#bf5a5a; padding:18px 18px 22px; max-width:640px;">
    <div class="flex flex-wrap items-center gap-2 sm:gap-3">
      <input class="task-input" placeholder="Brain dump a task..." bind:value={newTaskTitle} on:keydown={(e) => e.key==='Enter' && addTask()} />
      <select class="task-select" bind:value={newTaskType}>
        <option value="deep">Deep</option>
        <option value="shallow">Shallow</option>
      </select>
      <button class="task-add" on:click={addTask}>Add</button>
    </div>

    <div class="grid gap-4 sm:gap-6 mt-4 sm:mt-5 sm:grid-cols-2">
      <div>
        <div class="list-title">Deep</div>
        <ul class="list">
          {#each tasks.filter((t) => t.type === 'deep') as t (t.id)}
            <li class="item">
              <span>{t.title}</span>
              <button class="x" aria-label="remove" on:click={() => removeTask(t.id)}>×</button>
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
          {#each tasks.filter((t) => t.type === 'shallow') as t (t.id)}
            <li class="item">
              <span>{t.title}</span>
              <button class="x" aria-label="remove" on:click={() => removeTask(t.id)}>×</button>
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
  .task-add { background:#fff; color:#a64646; font-weight:700; padding:10px 16px; border-radius:8px; box-shadow:0 2px 0 #c84f4f; }

  .list-title { color:#fff; font-weight:700; margin-bottom:8px; }
  .list { display:flex; flex-direction:column; gap:8px; }
  .item { display:flex; align-items:center; justify-content:space-between; background:rgba(255,255,255,.1); color:#fff; padding:10px 12px; border-radius:8px; }
  .item.empty { opacity:.7; justify-content:center; }
  .x { background:transparent; color:#fff; opacity:.8; font-size:18px; line-height:1; padding:0 4px; }
</style>


