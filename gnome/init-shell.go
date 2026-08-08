package gnome

import (
	"os"

	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
)

func InitShell(applicationSid string, osArgs []string) {
	initGnome(applicationSid, osArgs)
}

func initGnome(applicationSid string, osArgs []string) {
	app := adw.NewApplication(applicationSid, 0)
	app.ConnectActivate(func() {
		interfaceUp(app)
	})
	os.Exit(app.Run(osArgs))
}

func interfaceUp(app *adw.Application) {
	window := InitWindow(app)

	content := CreateContent()

	toolbar := ToolbarCreate()
	sidebar := CreateSideBar(content)

	splitView := adw.NewNavigationSplitView()
	splitView.SetSidebar(sidebar.GetView())
	splitView.SetContent(adw.NewNavigationPage(sidebar.GetStack(), "Контент"))

	toastOverlay := adw.NewToastOverlay()
	toastOverlay.SetChild(splitView)

	toolbar.GetView().SetContent(toastOverlay)

	window.SetContent(toolbar.GetView())
	window.Present()
}
