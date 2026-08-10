package factory

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type SelectContolCreateParams struct {
	Title    string
	Subtitle string
}

type SelectContolButtonParams struct {
	Title string
	Name  string
	Icon  string
}

type SelectControlFactory struct {}

func UpSelectControlFactory() *SelectControlFactory {
	return &SelectControlFactory{}
}

func (self *SelectControlFactory) CreateSelect(params SelectContolCreateParams) (*adw.ComboRow, *adw.ViewStack) {
	stack := adw.NewViewStack()
	selectControl := buildSelect(params, stack)
	return selectControl, stack
}

func (self *SelectControlFactory) InjectButtonInSelectStack(params SelectContolButtonParams, stack *adw.ViewStack) {
	content := gtk.NewBox(gtk.OrientationVertical, 0)
	stack.AddTitledWithIcon(content, params.Name, params.Title, params.Icon)
}

func buildSelect(params SelectContolCreateParams, stack *adw.ViewStack) *adw.ComboRow {
	selectControl := adw.NewComboRow()
	selectControl.SetTitle(params.Title)
	selectControl.SetSubtitle(params.Subtitle)
	
	pages := stack.Pages()
	selectControl.SetModel(pages)
	selectControl.SetExpression(buildExpression(pages))
	return selectControl
}

func buildExpression(model *gtk.SelectionModel) *gtk.PropertyExpression {
	return  gtk.NewPropertyExpression(
		model.ItemType(),
		nil,
		"title",
	)
}