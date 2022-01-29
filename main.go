package main

import (
	"io/fs"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
)

var chownDir string

func main() {
	u, err := user.Current()
	if err != nil {
		log.Println(err)
		return
	}

	iUid, err := strconv.Atoi(u.Uid)
	if err != nil {
		log.Println(err)
		return
	}

	iGid, err := strconv.Atoi(u.Gid)
	if err != nil {
		log.Println(err)
		return
	}

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

	log.Printf("Chown %d:%d %s", iUid, iGid, dir)
	err = ChownR(dir, iUid, iGid)
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
