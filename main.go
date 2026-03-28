package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

var chownDir string

func main() {
	uid := os.Getuid()
	gid := os.Getgid()

	if chownDir == "" {
		log.Printf("Must be compiled with a valid directory to chown")
		return
	}

	dir := filepath.Clean(chownDir)
	absDir, err := filepath.Abs(dir)
	if err != nil {
		log.Println(err)
		return
	}

	if absDir != dir {
		log.Printf("Must use an absolute path not a relative path")
		return
	}

	log.Printf("Chown %d:%d %s", uid, gid, dir)
	err = ChownR(dir, uid, gid)
	if err != nil {
		log.Println(err)
	}
}

func ChownR(dir string, uid, gid int) error {
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err == nil {
			if d.Type()&fs.ModeSymlink == 0 {
				err = os.Chown(path, uid, gid)
			}
		}
		return err
	})
	return err
}
