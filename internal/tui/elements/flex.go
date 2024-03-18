package elements

import "github.com/rivo/tview"

func (e *Elements) initFlex(sidebarSize, mainContentSize int) {
	e.Flex.Clear()

	sidebar := e.buildSidebar()
	mainContent := e.buildMainContent()
	app := e.buildApp(sidebar, mainContent, sidebarSize, mainContentSize)
	footer := e.buildFooter()

	e.Flex.SetDirection(tview.FlexRow)
	e.Flex.AddItem(app, 0, 30, false)
	e.Flex.AddItem(footer, 0, 1, false)
}

func (e *Elements) ResizeFlex(sidebarSize, mainContentSize int) {
	e.initFlex(sidebarSize, mainContentSize)
}

func (e *Elements) buildSidebar() *tview.Flex {
	sidebar := tview.NewFlex()
	sidebar.SetDirection(tview.FlexRow)
	sidebar.AddItem(e.Tree, 0, 20, true)
	sidebar.AddItem(e.Search, 3, 0, false)
	return sidebar
}

func (e *Elements) buildMainContent() *tview.Flex {
	ht := tview.NewFlex()
	ht.AddItem(e.History, 0, 6, false)
	ht.AddItem(e.Timings, 0, 6, false)

	mainContent := tview.NewFlex()
	mainContent.SetDirection(tview.FlexRow)
	mainContent.AddItem(e.Output, 0, 15, false)
	mainContent.AddItem(ht, 0, 5, false)
	mainContent.AddItem(e.InfoBox, 3, 0, false)
	return mainContent
}

func (e *Elements) buildApp(sidebar, mainContent *tview.Flex, sidebarSize, mainContentSize int) *tview.Flex {
	app := tview.NewFlex()
	app.AddItem(sidebar, 0, sidebarSize, false)
	app.AddItem(mainContent, 0, mainContentSize, false)
	return app
}

func (e *Elements) buildFooter() *tview.Flex {
	footer := tview.NewFlex()
	footer.AddItem(e.Legend, 0, 1, false)
	return footer
}
