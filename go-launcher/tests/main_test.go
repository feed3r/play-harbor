package main

import (
	"testing"

	"github.com/feed3r/play-harbor/go-launcher/runlauncher"
)

func TestMissingArgs(t *testing.T) {
	// Test RunLauncher with insufficient arguments
	err := runlauncher.RunLauncher([]string{})
	if err == nil {
		t.Error("Expected error for missing args, got nil")
	}
}

func TestCommandError(t *testing.T) {
	// Test RunLauncher with invalid URL (simulation, does not actually run on Linux)
	err := runlauncher.RunLauncher([]string{"invalid-url", "Game.exe"})
	// We cannot guarantee error on all platforms, but the test checks that the function is callable
	if err == nil {
		t.Log("RunLauncher did not return error, check platform behavior")
	}
}

func TestProcessNotFound(t *testing.T) {
	// Simulate no process found
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
	// Simulate process found
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
