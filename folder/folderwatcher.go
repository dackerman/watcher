package main

import (
	"flag"
	"fmt"
	"github.com/dackerman/watcher"
	"os"
	"os/exec"
	"strings"
)

var folder string
var cmd string
var recurse bool

func main() {
	flag.StringVar(&folder, "folder", "~", "The folder to watch")
	flag.StringVar(&cmd, "cmd", "", "The command to run when the folder changes")
	flag.BoolVar(&recurse, "recurse", true, "Controls whether the watcher should recurse into subdirectories")

	flag.Parse()

	splitCmd := strings.Split(cmd, " ")
	if strings.TrimSpace(splitCmd[0]) == "" {
		fmt.Fprintf(os.Stdout, "Command (%v) has too few args\n", cmd)
		return
	}

	notify := watcher.WatchDirectory(folder, recurse)

	fmt.Fprintf(os.Stdout, "Will run `%v` when %v changes\n", cmd, folder)

	watcher.ExecuteOnChange(notify, func() {
		cmdPtr := exec.Command(splitCmd[0], splitCmd[1:]...)
		cmdPtr.Stdout = os.Stdout
		cmdPtr.Stderr = os.Stderr
		err := cmdPtr.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Command failed! %v\n", err)
		}
	})
	fmt.Fprintln(os.Stdout, "Shutting down watcher.")
}
