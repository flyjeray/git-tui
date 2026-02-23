package main

import (
	"fmt"
	"os"

	git "git-tui/git-ops"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	repoWarning string
	repoBranch  string
	repoRemotes []string
}

func initialModel(repo *git.Repo, repoWarning string) model {
	m := model{repoWarning: repoWarning}

	if repo != nil {
		if branch, err := repo.Branch(); err == nil {
			m.repoBranch = branch
		}
		if remotes, err := repo.Remotes(); err == nil {
			m.repoRemotes = remotes
		}
	}

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok {
		if key.String() == "ctrl+c" || key.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.repoWarning != "" {
		return warnStyle.Render("⚠ "+m.repoWarning) + "\n\nPress q to quit.\n"
	}

	s := "Current Branch: " + hintStyle.Render(m.repoBranch) + "\n"
	if len(m.repoRemotes) > 0 {
		s += "Remotes:" + hintStyle.Render("\n")
		for _, remote := range m.repoRemotes {
			s += hintStyle.Render("* "+remote) + "\n"
		}
	}
	s += "\nPress q to quit.\n"
	return s
}

func main() {
	selfInstall()

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(warnStyle.Render("⚠ could not determine working directory: " + err.Error()))
		os.Exit(1)
	}

	repo, repoErr := git.Find(cwd)
	var repoWarning string
	if repoErr != nil {
		repoWarning = "not a git repository — use 'git init' to initialize one"
	}

	p := tea.NewProgram(initialModel(repo, repoWarning))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
