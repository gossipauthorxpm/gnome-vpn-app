package gnome

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
)

type Sidebar struct {
	view  *adw.NavigationPage
	stack *adw.ViewStack
}

func CreateSidebar(content *Content) *Sidebar {
	stack := buildStackForInnerSidebar(content)
	innerSidebarList := buildInnerSidebarList(content, stack)
	sidebar := adw.NewNavigationPage(innerSidebarList, "Меню")
	return &Sidebar{view: sidebar, stack: stack}
}

func (self *Sidebar) GetView() *adw.NavigationPage {
	return self.view
}

func (self *Sidebar) GetStack() *adw.ViewStack {
	return self.stack
}

func buildStackForInnerSidebar(content *Content) *adw.ViewStack {
	stack := adw.NewViewStack()

	listPages := content.GetPages()

	for _, page := range listPages {
		page.InjectInSidebar(stack)
	}

	return stack
}

func buildInnerSidebarList(content *Content, stack *adw.ViewStack) *adw.NavigationPage {
	stack.SetVisibleChildName("connection")

	sidebar := adw.NewViewSwitcherSidebar()
	sidebar.SetStack(stack)

	page := adw.NewNavigationPage(sidebar, "Меню")
	return page
}
