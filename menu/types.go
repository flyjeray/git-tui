package menu

import git "git-tui/git-ops"

// Operation is returned by Activate to tell the model what to show next.
type Operation struct {
	Title         string
	ConfirmPrompt *ConfirmPrompt
	Flow          *InputFlow
}

// MenuItem is one entry in a navigable menu.
// Set exactly one of: Activate (leaf), Children (static submenu), Dynamic (runtime submenu).
type MenuItem struct {
	Label     string
	Operation func(*git.Repo) Operation  // leaf: returns what to show next
	Children  []MenuItem                 // static submenu
	Submenu   func(*git.Repo) []MenuItem // runtime-built submenu (e.g. list of remotes)
}

// MenuLevel is one frame in the navigation stack.
type MenuLevel struct {
	Items  []MenuItem
	Cursor int
}

// non-nil ConfirmPrompt triggers a yes/no prompt
type ConfirmPrompt struct {
	Prompt string
	OnYes  func(*git.Repo) string
}

// SelectOption is one choice in a select-style input step.
type selectOption struct {
	Value string // submitted to onSubmit
	Label string // human-readable description
}

// InputStep is one field in a multi-step input form.
// If options is non-nil the step renders as a navigable select; otherwise as a text field.
type inputStep struct {
	Label   string
	Hint    string         // shown dimmed below an active text field
	Value   string         // accumulated text, or selected option value
	Options []selectOption // non-nil = select mode
	Cursor  int            // active cursor within options
}

// non-nil InputFlow triggers a multi-step text input form.
type InputFlow struct {
	Title    string
	Steps    []inputStep
	Current  int
	OnSubmit func(r *git.Repo, values []string) string
}
