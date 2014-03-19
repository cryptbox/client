package main

import (
	"fmt"
	"os"
	"time"
	"github.com/jcelliott/lumber"
	"code.google.com/p/go.exp/fsnotify"
)

var log *lumber.FileLogger

func main() {

	// global logging
	var logerr error
	log, logerr = lumber.NewRotateLogger("cryptbox.log", 5000, 9)
	log.Level(lumber.DEBUG);

	if logerr != nil {
		fmt.Print("Unable to create new logger.")
	}


	//go StartWebServer()

	// timer
	var SleepTimer time.Duration
	SleepTimer = 2000000000 // 2 seconds

	// watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error(err.Error())
	}
	fs := NewFileSysMonitor("/home/apollitt/cryptbox", watcher)

	log.Debug("Starting client")

	for {
		log.Debug("Checking file system")
		fs.Check()
		time.Sleep(SleepTimer)
	}
	os.Exit(0)
}
