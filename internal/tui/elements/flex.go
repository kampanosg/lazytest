package elements

import "github.com/rivo/tview"

func (e *Elements) initFlex() {
	sidebar := tview.NewFlex()
	sidebar.SetDirection(tview.FlexRow)
	sidebar.AddItem(e.Tree, 0, 20, true)
	sidebar.AddItem(e.Search, 3, 0, false)

	mainContent := tview.NewFlex()
	mainContent.SetDirection(tview.FlexRow)
	mainContent.AddItem(e.Output, 0, 20, false)
	mainContent.AddItem(e.InfoBox, 3, 0, false)

	app := tview.NewFlex()
	app.AddItem(sidebar, 0, 1, false)
	app.AddItem(mainContent, 0, 2, false)

	footer := tview.NewFlex()
	footer.AddItem(e.Legend, 0, 1, false)

	e.Flex.SetDirection(tview.FlexRow)
	e.Flex.AddItem(app, 0, 30, false)
	e.Flex.AddItem(footer, 0, 1, false)
}
