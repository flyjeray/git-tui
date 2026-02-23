package menu

import git "git-tui/git-ops"

// MenuItem is one entry in a navigable menu.
// Exactly one of Submenu, Result, Confirm, or Flow should be set.
type MenuItem struct {
	Label   string
	Submenu func(*git.Repo) []MenuItem // non-nil = push a new menu level
	Result  func(*git.Repo) string     // non-nil = show immediate result string
	Confirm *ConfirmPrompt             // non-nil = show yes/no prompt
	Flow    func(*git.Repo) *InputFlow // non-nil = show multi-step input form
}

// MenuLevel is one frame in the navigation stack.
type MenuLevel struct {
	Items  []MenuItem
	Cursor int
}

func (l *MenuLevel) MoveUp() {
	if l.Cursor > 0 {
		l.Cursor--
	}
}
func (l *MenuLevel) MoveDown() {
	if l.Cursor < len(l.Items)-1 {
		l.Cursor++
	}
}
func (l MenuLevel) Selected() MenuItem { return l.Items[l.Cursor] }

// ConfirmPrompt triggers a yes/no dialog.
type ConfirmPrompt struct {
	Prompt string
	OnYes  func(*git.Repo) string
}

// InputStep is one text field in a multi-step input form.
type InputStep struct {
	Label string
	Hint  string // shown dimmed below the active field
	Value string // accumulated text
}

func (s *InputStep) AppendText(text string) { s.Value += text }
func (s *InputStep) Backspace() {
	if len(s.Value) > 0 {
		s.Value = s.Value[:len(s.Value)-1]
	}
}

// InputFlow is a multi-step text input form.
type InputFlow struct {
	Title    string
	Steps    []InputStep
	Current  int
	OnSubmit func(r *git.Repo, values []string) string
}

func (f *InputFlow) CurrentStep() *InputStep { return &f.Steps[f.Current] }
func (f InputFlow) IsLast() bool             { return f.Current == len(f.Steps)-1 }
func (f *InputFlow) Advance()                { f.Current++ }

func (f InputFlow) CollectValues() []string {
	vals := make([]string, len(f.Steps))
	for i, s := range f.Steps {
		vals[i] = s.Value
	}
	return vals
}
