package elements

import "github.com/gdamore/tcell/v2"

const helpText = `
	[darkturquoise]1 / 2: [white]Focus on the tree / output 
	[darkturquoise]r: [white]Run the selected test / test suite
	[darkturquoise]a: [white]Run all tests
	[darkturquoise]f: [white]Run all failed tests
	[darkturquoise]p: [white]Run all passed tests
	[darkturquoise]/: [white]Search
	[darkturquoise]Enter: [white](in search mode) Go to the search results
	[darkturquoise]<ESC>: [white]Exit search mode
	[darkturquoise]C: [white](outside search mode) Clear search
	[darkturquoise]q: [white]Quit
	[darkturquoise]?: [white]Show this help message
`

func (e *Elements) initHelp() {
	e.HelpModal.SetText(helpText)
	e.HelpModal.SetBackgroundColor(tcell.ColorBlack)
	e.HelpModal.AddButtons([]string{"Exit <ESC>"})
}
