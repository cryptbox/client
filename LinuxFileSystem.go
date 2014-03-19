package main

import (
	"code.google.com/p/go.exp/fsnotify"
	"os"
	"path/filepath"
)

func NewFileSysMonitor(path string, watcher *fsnotify.Watcher) *FileSysMonitor {
	fs := new(FileSysMonitor)
	fs.Path = path
	fs.Files = make(map[string]*SyncableFile)
	fs.Watcher = *watcher

	err := fs.Watcher.Watch(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	return fs
}

func (fs *FileSysMonitor) Check() bool {
	if fs.HasChanged() {
		fs.Sync()
	}
	return true
}

func (fs *FileSysMonitor) UpdateObjectList(index string, f os.FileInfo) bool {
	fs.Files[index].Info = f
	return true
}

func (fs *FileSysMonitor) HasChanged() bool {
	var err error
	if fs.Scanned {
		err = filepath.Walk(fs.Path, fs.Wait)
	} else {
		// first pass, manually walk over directory contents
		err = filepath.Walk(fs.Path, fs.InitPathInfo)
		fs.Scanned = true
	}

	if err != nil {
		return false
	}

	return true
}

func (fs *FileSysMonitor) InitPathInfo(path string, f os.FileInfo, err error) error {
	index := filepath.Join(path, f.Name())
	//fmt.Printf("Visiting: %v\n", index)
	file, ok := fs.Files[index]
	if ok && file.Info.Size() != f.Size() {
		log.Debug("File has changed: ", index)
		fs.Files[index].Info = f
	} else if !ok {
		log.Info("New file found: ", index)
		fs.Files[index] = new(SyncableFile)
		fs.Files[index].Info = f
	} else {
		//fmt.Println("Object has not changed.")
	}

	return nil
}

func (fs *FileSysMonitor) Wait(path string, f os.FileInfo, err error) error {
	go func() {
		for {
			select {
			case ev := <-fs.Watcher.Event:
				log.Info("event:", ev)
			case err := <-fs.Watcher.Error:
				log.Error("error:", err.Error())
			}
		}
	}()
	return nil
}

func (fs *FileSysMonitor) Sync() bool {
	return true
}
