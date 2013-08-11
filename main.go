package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var SleepTimer time.Duration
	SleepTimer = 2000000000 // 2 seconds
	fs := NewGenericFileSystem()
	fs.Path = "/home/apollitt/cryptbox"

	for {
		fmt.Println("Checking directory")
		if fs.HasChanged() {
			fs.Sync()
		}
		time.Sleep(SleepTimer)
	}
	os.Exit(0)
}
