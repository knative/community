package main

import (
	"fmt"
	"log"

	"google.golang.org/api/drive/v3"
)

const (
	KnativeCommunityDriveId = "0APnJ_hRs30R2Uk9PVA"
	KnativeTeamDriveId      = "0AM-QGZJ-HUA8Uk9PVA"
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

	drives := map[string]string{
		"Knative Community [old]": KnativeCommunityDriveId,
		"Knative Team":            KnativeTeamDriveId,
	}

	for name, id := range drives {
		// List files from Knative Community Drive.
		r, err := srv.Files.List().PageSize(10).
			Corpora("drive").
			IncludeItemsFromAllDrives(true).
			SupportsAllDrives(true).
			DriveId(id).
			Fields("nextPageToken, files(id, name)").Do()
		if err != nil {
			log.Fatal("Unable to retrieve files:", err)
		}
		fmt.Printf("---------------------------\n%s files:\n---------------------------\n", name)
		if len(r.Files) == 0 {
			fmt.Println("No files found.")
		} else {
			for _, i := range r.Files {
				fmt.Printf("%s (%s)\n", i.Name, i.Id)
			}
		}
	}
}
