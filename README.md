# git-tui

A keyboard-driven terminal UI for common git operations.

## Setup

```
go run .
```

On first run the binary installs itself to `~/.local/bin/gt` and adds the directory to your shell `PATH`. Restart the terminal, then use `gt` from any directory.

## Usage

Run `gt` inside a git repository. Navigate with arrow keys or `j`/`k`, confirm with `enter`, go back with `esc`.

## Features

| Item | What it does |
|---|---|
| Pull | Pull from remote(s); shows a submenu when multiple remotes exist |
| Commit | Stage all changes and commit with a message |
| Push | Push to remote(s) |
| Branches | Show current branch; checkout any other branch |
| Remotes | List, add, or delete remotes |
| Log | Scrollable one-line commit log (`↓` past last entry shifts the window) |
| Run custom command | Run any `git <command>` and display the output |

## Project structure

```
main.go            entry point; starts the bubbletea program
install.go         self-install logic (~/.local/bin/gt)

git-ops/           thin wrappers around git CLI commands
  repo.go          Repo type and Find()
  branch.go        GetCurrentBranch, GetInactiveBranchList, Checkout
  remotes.go       GetRemoteNames, GetRemotesWithURLs, AddRemote, RemoveRemote
  stage.go         StageAll
  commit.go        Commit
  pull.go          PullFromRemote
  push.go          Push
  log.go           GetLog (paginated)
  command.go       RunCommand (arbitrary git subcommand)

menu/              menu items and navigation types
  types.go         MenuItem, MenuLevel, ScrollState, InputFlow, ConfirmPrompt
  menu.go          root menu assembly, GetStartMenu
  branch_*.go      branch current / list / root menu items
  remotes_*.go     remotes list / add / delete / root menu items
  pull.go          pull menu item
  push.go          push menu item
  commit.go        commit flow
  log.go           scrollable log menu item
  custom.go        custom git command flow

styles/
  text.go          lipgloss style helpers (Title, Warn, Success, Hint, …)

ui/
  ui.go            Model definition
  update.go        key handling and state transitions
  view.go          rendering
```
