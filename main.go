package main

import (
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {

	var burpsuite_pro_jar string = ""
	filepath.Walk("./", func(path string, info fs.FileInfo, err error) error {
		name := strings.ToLower(info.Name())
		if strings.HasPrefix(name, "burpsuite_pro") {
			burpsuite_pro_jar = info.Name()
			return nil
		}
		return nil
	})
	if burpsuite_pro_jar == "" {
		os.WriteFile("error.log", []byte("burpsuite_pro.jar not found \n"), 0666)
		return
	}
	cmd := exec.Command("java", "-javaagent:BurpLoaderKeygen.jar", "-noverify", "-jar", burpsuite_pro_jar)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()
}
