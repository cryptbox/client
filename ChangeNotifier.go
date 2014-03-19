package main

import (
	"os"
	"code.google.com/p/go.exp/fsnotify"
)

// inotify/readdirchange/kqueue
type FileSystemWatcher interface {
	Add(string) bool
	Wait() bool
}

type MonitorInterface interface {
	HasChanged() bool
	Sync() bool
	UpdatePathInfo() error
	Check() bool
}

type FileSysMonitor struct {
	Path    string
	Files   map[string]*SyncableFile
	fd      int // file descriptor
	epfd    int //epoll file descriptor
	HasDesc bool
	Scanned bool
	Watcher fsnotify.Watcher
}

type SyncableFile struct {
	Info      os.FileInfo
	Path      string
	Index     string
	WatchDesc int  // file descriptor
	HasDesc   bool // watch file descriptor exists
}
