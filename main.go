package main

import (
	"os"
	"os/exec"
	"path"
	"syscall"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		os.WriteFile("error.log", []byte(err.Error()), 0666)
		return
	}

	burpsuite_pro_jar := path.Join(pwd, "burpsuite_pro.jar")
	if !PathExists(burpsuite_pro_jar) {
		os.WriteFile("error.log", []byte("burpsuite_pro.jar not found \n"), 0666)
		return
	}

	default_jdk := path.Join(pwd, "jre", "bin", "java.exe")
	jdk := "java"
	if PathExists(default_jdk) {
		jdk = default_jdk
	}

	cmd := exec.Command(jdk, "--illegal-access=permit", "-javaagent:BurpLoaderKeygen.jar", "-noverify", "-jar", burpsuite_pro_jar)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()
}
