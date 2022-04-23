package main

import (
	"os/exec"
	"syscall"
)

// for windows, set HideWindow flag
func run(jdk string, burpsuite_pro_jar string) {
	cmd := exec.Command(jdk, "--illegal-access=permit", "-javaagent:BurpLoaderKeygen.jar", "-noverify", "-jar", burpsuite_pro_jar)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()
}
