package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type ChangeNotifier interface {
	HasChanged() bool
	Sync() bool
}

type GenericFileSystem struct {
	Path        string
	Objects map[string]*Object
}

type Object struct {
	Info os.FileInfo
	Path string
	Key  string
}

func NewGenericFileSystem() *GenericFileSystem {
	fs := new(GenericFileSystem)
	fs.Path = ""
	fs.Objects = make(map[string]*Object)
	return fs
}

func (fs GenericFileSystem) HasChanged() bool {
	err := filepath.Walk(fs.Path, fs.UpdatePathInfo)

	if err != nil {
		return false
	}

	return true
}

func (fs *GenericFileSystem) UpdatePathInfo(path string, f os.FileInfo, err error) error {
	index := filepath.Join(path, f.Name())
	//fmt.Printf("Visiting: %v\n", index)
	object, ok := fs.Objects[index]
	if ok && object.Info.Size() != f.Size() {
		fmt.Println("Object has changed: ", index)
		fs.Objects[index].Info = f
	} else if !ok {
		fmt.Println("New object found: ", index)
		fs.Objects[index] = new(Object)
		fs.Objects[index].Info = f
	} else {
		//fmt.Println("Object has not changed.")
	}

	return nil
}

func (fs *GenericFileSystem) Sync() bool {

	return true
}
