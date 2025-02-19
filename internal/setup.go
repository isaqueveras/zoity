package internal

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/energye/systray"

	"github.com/isaqueveras/zoity/assets"
	"github.com/isaqueveras/zoity/types"
)

func Setup() {
	systray.SetIcon(assets.Icon)
	systray.SetTitle("Zoity")
	systray.SetTooltip("Zoity")

	for idx := range services {
		item := systray.AddMenuItemCheckbox(services[idx].Name, "", false)

		startItem := item.AddSubMenuItem("Start", "start the service")
		stopItem := item.AddSubMenuItem("Stop", "stop the service")
		stopItem.Hide()

		startItem.Click(func() {
			startService(&services[idx])

			item.Check()
			stopItem.Show()
			startItem.Hide()
		})

		stopItem.Click(func() {
			getOpenPortsByPIDAndKillProcess(services[idx].Process)

			item.Uncheck()
			stopItem.Hide()
			startItem.Show()

			totalServiceRunning--
			if totalServiceRunning == 0 {
				systray.SetIcon(assets.Icon)
			}

			log.Println("Stoping service", services[idx].Process.Pid, services[idx].Name)
			services[idx].Process = nil
		})
	}

	systray.AddSeparator()
	systray.AddMenuItem("Exit", "exit").Click(func() {
		stopAllServices()
		systray.Quit()
	})
}

func stopAllServices() {
	for _, service := range services {
		if service.Process == nil {
			continue
		}
		if err := service.Process.Kill(); err != nil {
			log.Println("[ERROR] erro ao enviar sinal para fechar o processo", service.Process.Pid, service.Name)
		}
	}
}

func startService(service *types.Service) {
	command := exec.Command("bash", "-c", fmt.Sprintf(
		"%scd %s && %s", service.GetEnv(), service.Path, service.Command,
	))

	command.Stdout, command.Stderr = os.Stdout, os.Stderr
	if err := command.Start(); err != nil {
		fmt.Printf("[ERROR] Error starting service %s: %v\n", service.Name, err)
		return
	}

	service.Process = command.Process
	totalServiceRunning++

	systray.SetIcon(assets.IconActived)

	log.Println("Starting the service", service.Process.Pid, service.Name)
}

func getOpenPortsByPIDAndKillProcess(process *os.Process) {
	var out bytes.Buffer
	var ports []string

	{
		cmd := exec.Command("netstat", "-tulnp")
		cmd.Stdout = &out
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error running netstat: %v\n", err)
			return
		}
	}

	re := regexp.MustCompile(`:(\d+)\s+.*\s+(\d+)/`)
	lines := strings.Split(out.String(), "\n")

	for _, line := range lines {
		if matches := re.FindStringSubmatch(line); len(matches) > 2 {
			if processID, _ := strconv.Atoi(matches[2]); processID == process.Pid {
				ports = append(ports, matches[1])
			}
		}
	}

	if len(ports) == 0 {
		log.Println("No port found for PID", process.Pid)
		if err := process.Kill(); err != nil {
			log.Println("[ERROR] Error sending signal to stop process", process.Pid)
		}
		return
	}

	fmt.Printf("\n==================\nO processo %d está ouvindo nas portas: %v\n", process.Pid, ports)

	{
		cmd := exec.Command("kill", "-9", strconv.Itoa(process.Pid))
		if err := cmd.Run(); err != nil {
			fmt.Printf("[ERROR] Erro ao finalizar o processo %d: %v\n", process.Pid, err)
			return
		}
	}

	fmt.Printf("[INFO] Processo %d finalizado com sucesso\n", process.Pid)
}
