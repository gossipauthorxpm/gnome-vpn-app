package routing

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
)

type PageBuilder struct {}

func (self *PageBuilder) BuildPageView() *adw.PreferencesPage {
	return buildPage()
}

func buildPage() *adw.PreferencesPage {
	page := adw.NewPreferencesPage()

	buildHeaderPage(page)
	buildCenterPage(page)
	buildFooterPage(page)

	return page
}

func buildHeaderPage(page *adw.PreferencesPage) {

}

func buildCenterPage(page *adw.PreferencesPage) {

}

func buildFooterPage(page *adw.PreferencesPage) {

}
