package main

import "testing"

func TestHasChanged(t *testing.T) {
	fs = NewGenericFileSystem()
	fs.Path = "/tmp/cryptbox"

}
