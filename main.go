package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: encfs /src/path /temp/mnt\n")
		return
	}

	src := os.Args[1]
	mnt := os.Args[2]

	if src == "" || mnt == "" {
		fmt.Printf("Usage: encfs /src/path /temp/mnt\n")
		return
	}

	root, err := fs.NewLoopbackRoot(src)
	if err != nil {
		log.Fatal(err)
	}

	server, err := fs.Mount(mnt, root, &fs.Options{
		MountOptions: fuse.MountOptions{Debug: true},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Mounted %s as loopback on %s\n", src, mnt)
	fmt.Printf("\n\nCAUTION:\nwrite operations on %s will also affect (%s)\n\n", mnt, src)
	fmt.Printf("Unmount by calling 'fusermount -u %s'\n", mnt)

	server.Wait()
}
