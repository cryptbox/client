package main

import (
	"syscall"
	"code.google.com/p/go.exp/fsnotify"
)

type LinuxWatcher struct {
	Path    string
	fd      int // file descriptor
	HasFd   bool
	epfd    int // epoll file descriptor
	HasEpfd bool
	Mask	uint32
	wfd		int // watcher file descriptor
	Watcher fsnotify.Watcher
}

func NewLinuxWatcher() *LinuxWatcher {
	watcher := new(LinuxWatcher)
	watcher.Path = "";
	watcher.HasFd = false;
	watcher.HasEpfd = false;
	watcher.Mask = syscall.IN_MODIFY

	return watcher;
}

func (watcher *LinuxWatcher) Add(path string) bool {
	log.Debug("Adding a new linux watcher")
	if !watcher.HasFd {
		watcher.fd, _ = syscall.InotifyInit()
		watcher.HasFd = true
	}

	if !watcher.HasEpfd {
		watcher.epfd, _ = syscall.EpollCreate(1) // 1?
		watcher.HasEpfd = true
	}
	watcher.Path = path
	var errAdding error
	watcher.wfd, errAdding = syscall.InotifyAddWatch(watcher.fd, path, watcher.Mask)
	// TODO: missing call to epoll_ctl
	// link file descriptor to epoll watcher
	var event syscall.EpollEvent
	log.Debug("event: ", event)
	event.Events = syscall.EPOLLIN
	event.Fd = int32(watcher.fd)
	event.Pad = event.Fd

	syscall.EpollCtl(watcher.epfd, syscall.EPOLL_CTL_ADD, watcher.fd, &event)

	if errAdding != nil {
		log.Error("Error adding Inotify Watcher")
	}

	log.Debug("Added Inotify Watcher", watcher.epfd)

	return true
}

func (watcher *LinuxWatcher) Wait() bool {
	log.Debug("Waiting for changes")
	events := make([]syscall.EpollEvent, 1)
	syscall.EpollWait(watcher.epfd, events, -1)
	log.Debug("!! Changes detected: ", events)
	// TODO: look at events
	return true
}

