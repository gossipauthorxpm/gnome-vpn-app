package connection

import (
	"log"

	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"gossipauthorxpm.ru/vpn/adapter"
	"gossipauthorxpm.ru/vpn/adapter/servers"
	"gossipauthorxpm.ru/vpn/gnome/factory"
)

type PageBuilder struct{}

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

	group := adw.NewPreferencesGroup()
	group.SetTitle("Подключение")
	group.SetDescription("Выберите страну и подключитесь к VPN")

	selectControl, stack := factory.UpSelectControlFactory().CreateSelect(
		factory.SelectContolCreateParams{Title: "Выберите сервер подключения", Subtitle: ""},
	)

	buildSelectServerInStack(adapter.ServerController.GetSubscibtionServers(adapter.SubscriptionController), stack)

	factory.UpSelectControlFactory().InjectCallbackOnChangeInSelectControl(func(vsp *adw.ViewStackPage) {
		server := &servers.ServerVisual{Name: vsp.Name(), Title: vsp.Title()}
		adapter.ServerController.ChangeServerCallback(server)
	}, selectControl)

	group.Add(selectControl)
	page.Add(group)
}

func buildSelectServerInStack(servers []servers.ServerVisual, stack *adw.ViewStack) {
	for _, server := range servers {
		error := factory.UpSelectControlFactory().InjectButtonInSelectStack(
			factory.SelectContolButtonParams{Title: server.Title, Name: server.Name}, stack,
		)
		if error != nil {
			log.Fatalf(error.Error())
		}
	}
}

func buildCenterPage(page *adw.PreferencesPage) {

	group := adw.NewPreferencesGroup()
	group.SetTitle("Запуск VPN")

	tunModeSwitcher := adw.NewSwitchRow()
	tunModeSwitcher.SetTitle("TUN mode")
	group.Add(tunModeSwitcher)

	proxyModeSwitcher := adw.NewSwitchRow()
	proxyModeSwitcher.SetTitle("PROXY mode")
	group.Add(proxyModeSwitcher)

	page.Add(group)
}

func buildFooterPage(page *adw.PreferencesPage) {

}
