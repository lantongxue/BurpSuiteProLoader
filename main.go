package main

import (
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("java", "-javaagent:BurpLoaderKeygen.jar", "-noverify", "-jar", "burpsuite_pro_v2021.5.3.jar")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()
}
