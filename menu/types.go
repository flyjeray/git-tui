package menu

import git "git-tui/git-ops"

func (i MenuItem) DisplayLabel() string {
	if i.Submenu != nil || i.LevelSubmenu != nil {
		return i.Label + " ›"
	}
	return i.Label
}

// MenuItem is one entry in a navigable menu.
// Exactly one of Submenu, LevelSubmenu, Result, Confirm, or Flow should be set.
type MenuItem struct {
	Label        string
	Submenu      func(*git.Repo) []MenuItem // non-nil = push a new menu level
	LevelSubmenu func(*git.Repo) MenuLevel  // non-nil = push a pre-built menu level (supports scroll)
	Result       func(*git.Repo) string     // non-nil = show immediate result string
	Info         func(*git.Repo) string     // non-nil = show immediate info string (leaving will not reset app)
	Confirm      *ConfirmPrompt             // non-nil = show yes/no prompt
	Flow         func(*git.Repo) *InputFlow // non-nil = show multi-step input form
}

// ScrollState enables a MenuLevel to shift its items one-by-one when
// the cursor is at a boundary, giving an infinite-scroll effect.
type ScrollState struct {
	Offset int
	Fetch  func(offset int) []MenuItem
}

// MenuLevel is one frame in the navigation stack.
type MenuLevel struct {
	Items  []MenuItem
	Cursor int
	Scroll *ScrollState // non-nil = level supports scrolling
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

func (l *MenuLevel) ScrollUp() {
	if l.Cursor == 0 && l.Scroll != nil && l.Scroll.Offset > 0 {
		newOffset := l.Scroll.Offset - 1
		if items := l.Scroll.Fetch(newOffset); len(items) > 0 {
			l.Items = items
			l.Scroll.Offset = newOffset
		}
	} else {
		l.MoveUp()
	}
}

func (l *MenuLevel) ScrollDown() {
	if l.Cursor == len(l.Items)-1 && l.Scroll != nil {
		newOffset := l.Scroll.Offset + 1
		if items := l.Scroll.Fetch(newOffset); len(items) == len(l.Items) {
			l.Items = items
			l.Scroll.Offset = newOffset
		}
	} else {
		l.MoveDown()
	}
}

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
