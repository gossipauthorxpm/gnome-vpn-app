package factory

import (
	"fmt"

	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"gossipauthorxpm.ru/vpn/common"
)

type SelectContolCreateParams struct {
	Title    string
	Subtitle string
}

type SelectContolButtonParams struct {
	Title string
	Name  string
}

type SelectControlFactory struct{}

func UpSelectControlFactory() *SelectControlFactory {
	return &SelectControlFactory{}
}

func (self *SelectControlFactory) CreateSelect(params SelectContolCreateParams) (*adw.ComboRow, *adw.ViewStack) {
	stack := adw.NewViewStack()
	selectControl := buildSelect(params, stack)
	return selectControl, stack
}

func (self *SelectControlFactory) InjectButtonInSelectStack(params SelectContolButtonParams, stack *adw.ViewStack) error {
	if common.HasViewStackContainsName(stack, params.Name) {
		return fmt.Errorf("%w: %q", common.DublicateIdError, params.Name)
	}
	content := gtk.NewBox(gtk.OrientationVertical, 0)
	stack.AddTitled(content, params.Name, params.Title)
	return nil
}

func (self *SelectControlFactory) InjectCallbackOnChangeInSelectControl(callback func(*adw.ViewStackPage), selectControl *adw.ComboRow) {
	selectControl.NotifyProperty("selected-item", func() {
		item := selectControl.SelectedItem()
		if item == nil {
			return
		}

		callback(item.Cast().(*adw.ViewStackPage))
	})
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
	return gtk.NewPropertyExpression(
		model.ItemType(),
		nil,
		"title",
	)
}
