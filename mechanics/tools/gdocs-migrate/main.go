package main

import (
	"fmt"
	"log"
	"strings"

	"google.golang.org/api/drive/v3"
)

const (
	KnativeCommunityDriveID = "0APnJ_hRs30R2Uk9PVA"
	KnativeTeamDriveID      = "0AM-QGZJ-HUA8Uk9PVA"
	FolderMimeType          = "application/vnd.google-apps.folder"
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
		"Knative Community Drive": KnativeCommunityDriveID,
		"Knative Team Drive":      KnativeTeamDriveID,
	}

	forest := NewForest()
	for name, id := range drives {
		fmt.Println("Loading directory from ... ", name)
		forest.Add(File{
			id:       id,
			name:     name,
			isFolder: true,
		})
		var pageToken *string
		for pageToken == nil || *pageToken != "" {
			// List files from Knative Community Drive.
			request := srv.Files.List().PageSize(1000).
				Corpora("drive").
				IncludeItemsFromAllDrives(true).
				SupportsAllDrives(true).
				DriveId(id).
				Fields("nextPageToken, files(id, name, parents, mimeType)").
				OrderBy("folder,modifiedTime desc,name")
			if pageToken != nil {
				request.PageToken(*pageToken)
			}
			r, err := request.Do()
			if err != nil {
				log.Fatal("Unable to retrieve files:", err)
			}
			pageToken = &r.NextPageToken
			for _, i := range r.Files {
				forest.Add(File{
					id:       i.Id,
					name:     i.Name,
					parentID: i.Parents[0],
					isFolder: i.MimeType == FolderMimeType,
				})
			}
			fmt.Printf("Loaded %d files, %d folders\n", len(forest.GetFiles()), forest.FolderCount())
		}
	}

	for _, f := range forest.GetFiles() {
		fmt.Printf("[%s] %s\n", strings.Join(forest.GetPath(f), "/"), f.name)
	}
}
