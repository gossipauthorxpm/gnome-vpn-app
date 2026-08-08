package gnome

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
)

type ToolBar struct {
	view *adw.ToolbarView
}

func ToolbarCreate() *ToolBar {
	toolbarView := adw.NewToolbarView()
	header := adw.NewHeaderBar()
	header.SetTitleWidget(adw.NewWindowTitle("GNOME", ""))

	toolbarView.AddTopBar(header)

	return &ToolBar{view: toolbarView}
}

func (self *ToolBar) GetView() *adw.ToolbarView {
	return self.view
}
