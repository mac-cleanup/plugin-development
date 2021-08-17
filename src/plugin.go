package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	color "github.com/fatih/color"
)

type plugin string

var introMsg = "Cleaning up Development Files.."

func (g plugin) Cleanup() {
	fmt.Println(introMsg)
	DeleteFiles("~/.bash_history", "Bash history..")
	DeleteFiles("~/.zhistory", "ZSH history..")
	color.Cyan("Android caches..")
	DeleteFiles("~/.android/cache", "Android caches..")
	DeleteFiles("~/.gradle/caches", "")
	DeleteFiles("~/wget-log", "wget history..")
	DeleteFiles("~/.wget-hsts", "wget history..")
	DeleteFiles("~/Library/Logs/JetBrains/*/", "JetBrains logs..")
	DeleteFiles("~/Library/Developer/Xcode/DerivedData/*", "Xcode data..")
	DeleteFiles("~/Library/Developer/Xcode/Archives/*", "")
	DeleteFiles("~/Library/Developer/Xcode/iOS Device Logs/*", "")
	ShellCommand("gem cleanup", "Gem files..")
	ShellCommand("docker system prune -af", "Docker images..")
	ShellCommand("brew cleanup -s", "Homebrew..")
	ShellCommand("brew tap --repair", "")
	ShellCommand("npm cache clean --force", "NPM cache..")
	ShellCommand("yarn cache clean --force", "Yarn cache..")
	ShellCommand("pod cache clean --all", "Cocoapods cache..")
}

func DeleteFiles(dir string, message string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}

	color.Cyan(message + "...")
	return nil
}

func main() {
	plugin := plugin("")
	plugin.Cleanup()
}

func ShellCommand(command string, message string) {
	_, err := exec.Command("/bin/sh", "-c", command).Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	color.Cyan(message)
}

var Cleanup plugin
