package main

import (
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("java", "-javaagent:BurpLoaderKeygen.jar", "-noverify", "-jar", "burpsuite_pro.jar")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()
}
