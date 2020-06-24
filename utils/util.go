package utils

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
)

func GetEnv(key, defaultValue string) string {
	if envVal, exists := os.LookupEnv(key); exists {
		return envVal
	}
	return defaultValue
}

func Scanner() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func ClearScreen() {
	osRunning := runtime.GOOS
	if osRunning == "linux" || osRunning == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if osRunning == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
