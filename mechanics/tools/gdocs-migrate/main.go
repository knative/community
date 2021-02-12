package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/api/drive/v3"
	"knative.dev/pkg/pool"
)

func mirror(srv *drive.Service, forest *Forest, src *File, dst *File, w pool.Interface) error {
	srcPath, dstPath := forest.GetPath(src), forest.GetPath(dst)
	for _, child := range forest.Children(src) {
		// File copying operations are done in Go routines, thus it is
		// necessary to capture range variable here.
		child := child
		childName := strings.Join(append(srcPath, src.name, child.name), "/")
		cloneName := strings.Join(append(dstPath, dst.name, child.name), "/")
		clone, exist := forest.Children(dst)[child.name]
		if child.isFolder {
			if !exist { // Destination folder not exist
				log.Printf("Creating directory %q", cloneName)
				created, err := createDir(srv, dst, child.name)
				if err != nil {
					return err
				}
				forest.Add(created)
				clone = *created
			}
			if err := mirror(srv, forest, &child, &clone, w); err != nil {
				return err
			}
		}
		if !exist && !child.isFolder {
			// We won't recopy a file on a rerun. That means we will need to
			// do the diligence of deleting a partially copied file when exceptions happened.
			w.Go(func() error { // File copying can be done in parallel.
				log.Printf("Copying %q to %q", childName, cloneName)
				dstFile, err := copyFile(srv, &child, dst)
				log.Printf("Copied %q[%s] -> %q[%s]", childName, child.id, cloneName, dstFile.id)
				return errors.Wrap(err, fmt.Sprintf("Cannot copy %q to %q", childName, cloneName))
			})
		}
	}
	return nil
}

// flag.StringVars
var sourceDriveID, destDriveID string

func main() {
	flag.StringVar(&sourceDriveID, "sourceDriveID", "", "The source drive ID")
	flag.StringVar(&destDriveID, "destDriveID", "", "The dest drive ID")
	flag.Parse()

	if sourceDriveID == "" || destDriveID == "" {
		log.Fatal("Both --sourceDriveID and --destDriveID are required")
	}
	// You need to put credentials into a file named credentials.json.
	// TODO(evankanderson): Write some better docs.
	client, err := clientFromFile("credentials.json")
	if err != nil {
		log.Fatal("Unable to initialize client: ", err)
	}

	srv, err := drive.New(client)
	if err != nil {
		log.Fatal("Unable to retrieve Drive client: ", err)
	}

	source, err := driveAsFile(srv, sourceDriveID)
	if err != nil {
		log.Fatal("Failed to look up source drive: ", sourceDriveID, err)
	}
	dest, err := driveAsFile(srv, destDriveID)
	if err != nil {
		log.Fatal("Failed to look up dest drive: ", destDriveID, err)
	}

	forest := NewForest()
	for _, drive := range []*File{source, dest} {
		fmt.Println("Loading directory from", drive.name)
		loadAllFiles(&forest, srv, drive)
	}
	// Display some statistics about different file types
	countByTypes := make(map[string]int)
	for _, f := range forest.GetFiles() {
		countByTypes[f.mimeType]++
	}
	for kind, count := range countByTypes {
		fmt.Printf("%s:\t%d\n", kind, count)
	}
	// Create a thread pool for file copying operations. Don't be too
	// ambitious: 429s await.
	threads := pool.New(5)
	if err := mirror(srv, &forest, source, dest, threads); err != nil {
		log.Fatal("Error mirroring data ", err)
	}
	if err := threads.Wait(); err != nil {
		log.Fatal("Error mirroring data ", err)
	}
}
