package main

import (
	//"fmt"
	"os"
)

func checkFolder(dir string) bool {
	_, err := os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModeDir)
		if err == nil {
			return true
		}
	}
	return false
}

func (inbox *Inbox) Store() {

}

func (inbox *Inbox) Load() {

}