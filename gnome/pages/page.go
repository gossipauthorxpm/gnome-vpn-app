package pages

import "github.com/diamondburned/gotk4-adwaita/pkg/adw"

type Page interface {
	Get() *adw.PreferencesPage
	GetNavigationPage() *adw.NavigationPage
	InjectInSidebar(stack *adw.ViewStack)
}

type PageViewBuilder interface {
	BuildPageView() *adw.PreferencesPage
}
