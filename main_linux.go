package main

import (
	"os/exec"
)

// for linux
func run(jdk string, burpsuite_pro_jar string) {
	cmd := exec.Command(jdk, "--illegal-access=permit", "-javaagent:BurpLoaderKeygen.jar", "-noverify", "-jar", burpsuite_pro_jar)
	cmd.Start()
}
