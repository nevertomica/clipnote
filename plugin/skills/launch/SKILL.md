---
name: launch
description: Launch clipnote tmux annotation session. Use when the user
  explicitly mentions clipnote, wants to annotate/mark/highlight AI output,
  wants to review AI responses or take notes, wants a split-pane annotation
  panel, or wants to export/summarize key points from AI output.
---

# clipnote:launch

Launch the clipnote annotation session via tmux.

## Instructions

1. When you detect the user's intent matches this skill, first explain:
   - clipnote will open a **new tmux session** with an AI CLI and an annotation panel side by side
   - This will **leave the current Claude Code environment** — the user will switch to a separate tmux workspace
2. Ask the user to confirm before proceeding (e.g. "Shall I launch clipnote now?")
3. Only after the user confirms, run the following command using the Bash tool:

```bash
CLIPNOTE_CLI=claude "${CLAUDE_PLUGIN_ROOT}/bin/clipnote"
```

This will create a tmux session with Claude CLI on the left and the annotation panel on the right.

Note: `CLIPNOTE_CLI=claude` bypasses the interactive CLI selector, which cannot run inside Claude Code's non-TTY environment.

## Keybindings (tell the user)

| Key | Action |
|-----|--------|
| r | Capture visible area |
| R | Custom range capture |
| Ctrl+r | Clear all content |
| j/k | Move cursor up/down |
| g/G | Jump to top/bottom |
| m | Toggle mark |
| c | Mark + add note |
| v | View note |
| S | Export marks to clipboard |
| P | Paste marks to left pane |
| [/] | Shrink/expand content panel |
| ? | Show help |
| q | Quit |
