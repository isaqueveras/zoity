package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/energye/systray"

	"github.com/isaqueveras/zoity/assets"
	"github.com/isaqueveras/zoity/types"
)

func main() {
	systray.Run(setup, func() {
		log.Println("Closing Zoity Application")
	})
}

func setup() {
	systray.SetIcon(assets.Icon)
	systray.SetTitle("Zoity")
	systray.SetTooltip("Zoity")

	createItemsInMenu()
}

func stopAllServices() {
	for _, service := range types.Services {
		if service.Process == nil {
			continue
		}
		service.Kill()
	}
}

func startService(service *types.Service) {
	service.Kill()

	command := exec.Command("bash", "-c", fmt.Sprintf(
		"%scd %s && %s", service.GetEnv(), service.Path, service.Command,
	))

	command.Stdout, command.Stderr = os.Stdout, os.Stderr
	if err := command.Start(); err != nil {
		fmt.Printf("[ERROR] Error starting service %s: %v\n", service.Name, err)
		return
	}

	service.Process = command.Process
	types.TotalServiceRunning++

	systray.SetIcon(assets.IconActived)
	log.Println("Starting the service", service.Process.Pid, service.Name, service.Ports)
}

func createItemsInMenu() {
	for idx := range types.Services {
		createItemService(idx, &types.Services)
	}

	systray.AddSeparator()

	settings := systray.AddMenuItem("Settings", "settings")
	settings.AddSubMenuItem("Reload", "reload settings").Click(reloadSettings)
	settings.AddSubMenuItem("Open", "open settings file").Click(openSettingsFile)

	systray.AddMenuItem("Exit", "exit").Click(exit)
}

func reloadSettings() {
	stopAllServices()
	systray.ResetMenu()
	systray.SetIcon(assets.Icon)
	types.InitConfig()
	createItemsInMenu()
}

func exit() {
	stopAllServices()
	systray.Quit()
}

func createItemService(idx int, services *[]types.Service) {
	if services == nil || idx < 0 || idx >= len(*services) {
		fmt.Println("[ERROR] Invalid index or null slice")
		return
	}

	service := (*services)[idx]
	item := systray.AddMenuItemCheckbox(service.Name, service.Name, false)

	startItem := item.AddSubMenuItem("Start", "start the service")
	stopItem := item.AddSubMenuItem("Stop", "stop the service")
	stopItem.Hide()

	startItem.Click(func() {
		startService(&types.Services[idx])
		item.Check()
		stopItem.Show()
		startItem.Hide()
	})

	stopItem.Click(func() {
		service.Kill()
		item.Uncheck()
		stopItem.Hide()
		startItem.Show()

		types.TotalServiceRunning--
		if types.TotalServiceRunning == 0 {
			systray.SetIcon(assets.Icon)
		}

		log.Println("Stoping service", service.Name)
		types.Services[idx].Process = nil
	})
}

func openSettingsFile() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", types.ConfigFile)
	case "darwin":
		cmd = exec.Command("open", types.ConfigFile)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "", types.ConfigFile)
	default:
		log.Printf("unsupported operating system: %s\n", runtime.GOOS)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
