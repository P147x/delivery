package main

import (
	"time"

	"delivery/services"

	"github.com/0xAX/notificator"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	var notify *notificator.Notificator
	notify = notificator.New(notificator.Options{
		DefaultIcon: "img/default.png",
		AppName:     "Delivery",
	})
	config := Load("config/services.json")
	systray.SetIcon(icon.Data)
	var menuItems [5]*systray.MenuItem
	var update *systray.MenuItem

	var servs []services.Service

	for i, s := range config.Services {
		servs = append(servs, services.ServiceFactory(s.Delivery, s.TrackingID))
		servs[i].Update()
		menuItems[i] = systray.AddMenuItem(servs[i].GetLastEvent(), s.TrackingID)
		systray.AddSeparator()
	}
	update = systray.AddMenuItem("", "")
	update.Disable()

	for true {
		for i, s := range servs {
			if s.Update() == true {
				menuItems[i].SetTitle(s.GetLastEvent())
				notify.Push("Your package moved !", servs[i].GetLastEvent(), "img/icon.png", notificator.UR_NORMAL)
			}
		}
		update.SetTitle("Last check :" + time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(60 * time.Second)
	}
}

func onExit() {
	// clean up here
}
