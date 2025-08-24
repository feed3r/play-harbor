package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	 if len(os.Args) != 3 {
	 		fmt.Println("ERROR: Needs launch URL and EXE Name")
	 		fmt.Println("Usage: playdock.exe <epicUrl> <exeName>")
	 		return
	 	}

	epicUrl := os.Args[1]
	exeName := os.Args[2]

	fmt.Printf("Starting url: %s\n", epicUrl)

	// Apri l'URL (Windows lo aprirà con il gestore registrato)
	cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", epicUrl)
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Failed to start URL: %v\n", err)
		return
	}

	// Attendi qualche secondo per dare tempo al gioco di partire
	time.Sleep(5 * time.Second)

	// Cerca il processo con quel nome
	procs, err := process.Processes()
	if err != nil {
		fmt.Printf("Error listing processes: %v\n", err)
		return
	}

	var gameProc *process.Process
	for _, p := range procs {
		name, _ := p.Name()
		if strings.EqualFold(name, exeName+".exe") || strings.EqualFold(name, exeName) {
			gameProc = p
			break
		}
	}

	if gameProc == nil {
		fmt.Printf("Could not find a process with name: %s\n", exeName)
		return
	}

	fmt.Println("Game started. Waiting for it to exit...")

	// Attendi che il gioco termini
	for {
		exists, _ := process.PidExists(gameProc.Pid)
		if !exists {
			break
		}
		time.Sleep(2 * time.Second)
	}

	fmt.Println("Game exited.")
}
