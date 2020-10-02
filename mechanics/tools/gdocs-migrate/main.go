package main

import (
	"fmt"
	"log"

	"google.golang.org/api/drive/v3"
)

func main() {
	// You need to put credentials into a file named credentials.json.
	// TODO(evankanderson): Write some better docs.
	client, err := clientFromFile("credentials.json")
	if err != nil {
		log.Fatal("Unable to initialize client:", err)
	}

	srv, err := drive.New(client)
	if err != nil {
		log.Fatal("Unable to retrieve Drive client:", err)
	}

	// Copied from https://developers.google.com/drive/api/v3/quickstart/go
	// TODO: replace this
	r, err := srv.Files.List().PageSize(10).
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		log.Fatal("Unable to retrieve files:", err)
	}
	fmt.Println("Files:")
	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range r.Files {
			fmt.Printf("%s (%s)\n", i.Name, i.Id)
		}
	}
}
