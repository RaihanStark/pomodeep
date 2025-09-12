<script lang="ts">
	import * as Tooltip from '$lib/components/ui/tooltip/index';

	let distractionText = '';
</script>

<Tooltip.Provider delayDuration={150}>
	<div class="w-full">
		<div class="mx-auto rounded-md" style="background-color:#bf5a5a; padding:18px 18px 22px;">
			<div class="list-title">
				Distraction Jot <Tooltip.Root
					><Tooltip.Trigger aria-label="Definition" style="cursor:help;">ⓘ</Tooltip.Trigger
					><Tooltip.Content side="right" sideOffset={8}>
						<div style="max-width:320px; line-height:1.45;">
							<div style="font-weight:700; margin-bottom:6px;">🔑 Definition</div>
							<p>
								It's a designated capture space (paper notebook, index card, or a single text file)
								that you keep open while doing a deep work block. Whenever a distracting thought,
								urge, or task pops up, you jot it down immediately instead of acting on it.
							</p>
						</div>
					</Tooltip.Content></Tooltip.Root
				>
			</div>
			<p class="mb-4 text-sm text-white opacity-80">
				Brain dump your thoughts here to stay focused on deep work
			</p>
			<textarea
				bind:value={distractionText}
				placeholder="Write down any distracting thoughts, ideas, or tasks that come to mind..."
				class="distraction-textarea"
			></textarea>
		</div>
	</div>
</Tooltip.Provider>

<style>
	.list-title {
		color: #fff;
		font-weight: 700;
		margin-bottom: 8px;
	}

	.distraction-textarea {
		width: 100%;
		height: 128px;
		resize: none;
		background: #fff;
		color: #5b2727;
		padding: 12px;
		border-radius: 8px;
		outline: none;
		border: 0;
		font-size: 14px;
		line-height: 1.4;
	}

	.distraction-textarea::placeholder {
		color: #999;
	}

	.distraction-textarea:focus {
		box-shadow: 0 0 0 2px rgba(255, 255, 255, 0.3);
	}
</style>
