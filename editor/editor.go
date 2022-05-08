package editor

import (
	"github.com/charmbracelet/bubbles/viewport"
	_ "github.com/charmbracelet/bubbletea"
	_ "github.com/charmbracelet/lipgloss"
)

// Editor
type Editor struct {
	tg *textGrid
	vp viewport.Model
}

// NewEditor creates and returns a new Editor
// instance.
func NewEditor() *Editor {
	return &Editor{
		tg: newTextGrid(),
		vp: viewport.New(0, 0),
	}
}
