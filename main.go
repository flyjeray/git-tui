package main

import (
	"fmt"
	"os"

	git "git-tui/git-ops"
	styles "git-tui/styles"
	ui "git-tui/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	selfInstall()

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(styles.Warn("⚠ could not determine working directory: " + err.Error()))
		os.Exit(1)
	}

	repo, repoErr := git.Find(cwd)
	var repoWarning string = ""
	if repoErr != nil {
		repoWarning = "not a git repository"
	}

	p := tea.NewProgram(ui.InitialModel(repo, repoWarning))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v\n", err)
		os.Exit(1)
	}
}
