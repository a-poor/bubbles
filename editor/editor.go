package editor

import (
	_ "github.com/charmbracelet/bubbles/viewport"
	_ "github.com/charmbracelet/bubbletea"
	_ "github.com/charmbracelet/lipgloss"
)

type Editor struct {
	tg *textGrid
}

func NewEditor() *Editor {
	return &Editor{
		tg: newTextGrid(),
	}
}
