package gotui

import (
	"os"
	"os/exec"
	"strings"
)

func PrintStars(val int) {
	cmd := exec.Command("echo", strings.Repeat("*", val))
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	cmd.Wait()
}
