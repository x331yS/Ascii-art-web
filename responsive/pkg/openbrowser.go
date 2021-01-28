package pkg

import (
	"log"
	"os/exec"
	"runtime"
)

func Openbrowser(zz string) {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", zz).Start()

	case "linux":
		err = exec.Command("xdg-open", zz).Start()
	}
	if err != nil {
		log.Fatal(err)
	}
}
