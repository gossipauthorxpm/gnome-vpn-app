package buttons

import (
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func Init(callback func()) *gtk.Button {
	button := gtk.NewButtonWithLabel("Поздороваться")
	button.ConnectClicked(callback)
	return button
}
