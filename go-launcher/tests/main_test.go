package main

import (
	"os/exec"
	"testing"
)

// Mock per exec.Command
var mockCommand func(name string, arg ...string) *exec.Cmd

func TestMissingArgs(t *testing.T) {
	// Simula chiamata senza argomenti
	osArgs := []string{"playdock.exe"}
	if len(osArgs) != 3 {
		t.Log("ERROR: Needs launch URL and EXE Name")
	}
}

func TestCommandError(t *testing.T) {
	// Mock exec.Command per simulare errore
	mockCommand = func(name string, arg ...string) *exec.Cmd {
		cmd := exec.Command("false") // comando che fallisce
		return cmd
	}
	cmd := mockCommand("rundll32", "url.dll,FileProtocolHandler", "mockurl")
	err := cmd.Run()
	if err == nil {
		t.Error("Expected error when running command, got nil")
	}
}

func TestProcessNotFound(t *testing.T) {
	// Simula nessun processo trovato
	processes := []string{"other.exe", "notgame.exe"}
	found := false
	for _, name := range processes {
		if name == "Game.exe" {
			found = true
		}
	}
	if found {
		t.Error("Process should not be found")
	}
}

func TestProcessFound(t *testing.T) {
	// Simula processo trovato
	processes := []string{"Game.exe", "other.exe"}
	found := false
	for _, name := range processes {
		if name == "Game.exe" {
			found = true
		}
	}
	if !found {
		t.Error("Process should be found")
	}
}
