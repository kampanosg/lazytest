package clipboard

import "github.com/atotto/clipboard"

// ClipboardManager is a wrappen of atotto/clipboard so it can easily be used as paramater and mocked
type ClipboardManager struct {}

func NewClipboardManager() *ClipboardManager {
	return &ClipboardManager{}
}

func (c *ClipboardManager) WriteAll(text string) error {
	return clipboard.WriteAll(text)
}