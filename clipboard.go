package main

import (
	"fmt"

	"github.com/atotto/clipboard"
)

func (m *Model) CopyMarksToClipboard() string {
	text := m.ExportMarks()
	if text == "" {
		return "No marks to export"
	}
	if err := clipboard.WriteAll(text); err != nil {
		return fmt.Sprintf("Clipboard write failed: %v", err)
	}
	return fmt.Sprintf("Copied %d marks to clipboard", len(m.marks))
}
