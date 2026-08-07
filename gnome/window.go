package gnome

import "github.com/diamondburned/gotk4-adwaita/pkg/adw"

func InitWindow(app *adw.Application) *adw.ApplicationWindow {
	window := adw.NewApplicationWindow(&app.Application)
	window.SetTitle("Basic App")
	window.SetDefaultSize(420, 240)
	return window
}
