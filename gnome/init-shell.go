package gnome

import (
	"os"

	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"gossipauthorxpm.ru/vpn/gnome/buttons"
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

	label := gtk.NewLabel("Нажмите кнопку")

	button := buttons.Init(func() {
		label.SetLabel("Привет, GNOME!")
	})

	content := gtk.NewBox(gtk.OrientationVertical, 12)
	content.SetHAlign(gtk.AlignCenter)
	content.SetVAlign(gtk.AlignCenter)

	content.Append(label)
	content.Append(button)

	headerBar := adw.NewHeaderBar()

	toolbarView := adw.NewToolbarView()
	toolbarView.AddTopBar(headerBar)
	toolbarView.SetContent(content)

	window.SetContent(toolbarView)
	window.Present()
}
