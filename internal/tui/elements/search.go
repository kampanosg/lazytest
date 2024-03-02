package elements

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var searchPlaceholderStyle = tcell.StyleDefault.Foreground(tcell.ColorGray)

func (e *Elements) initSearch() {
	e.Search.SetTitle("Search")
	e.Search.SetBorder(true)
	e.Search.SetBackgroundColor(tcell.ColorDefault)
	e.Search.SetTitleAlign(tview.AlignLeft)
	e.Search.SetFieldBackgroundColor(tcell.ColorDefault)
	e.Search.SetPlaceholder("Press / to search")
	e.Search.SetPlaceholderStyle(searchPlaceholderStyle)
	e.Search.SetDoneFunc(e.handlers.handleSearchDone)
	e.Search.SetChangedFunc(e.handlers.handleSearchChanged)
}
