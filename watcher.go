package watcher

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func WatchDirectory(dir string, recurse bool) <-chan int {
	notify := make(chan int)
	var prevMax int64
	var currentMax int64
	go func() {
		for {
			currentMax = findMaxTimestamp(dir, currentMax, recurse)

			// If prevMax was zero, directory hasn't really changed yet
			if prevMax < currentMax && prevMax != 0 {
				notify <- 1
			}
			prevMax = currentMax
			time.Sleep(1 * time.Second)
		}
	}()
	return notify
}

func findMaxTimestamp(dir string, startMax int64, recurse bool) (currentMax int64) {
	currentMax = startMax
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		if currentMax < f.ModTime().Unix() {
			currentMax = f.ModTime().Unix()
		}
		if f.IsDir() && recurse {
			currentMax = findMaxTimestamp(dir+"/"+f.Name(), currentMax, recurse)
		}
	}
	return currentMax
}

func ExecuteOnChange(notify <-chan int, f func()) {
	for {
		<-notify
		fmt.Fprintf(os.Stdout, "Detected change at %v\n", time.Now())
		f()
	}
}
