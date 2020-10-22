package main

import (
	"fmt"
	"log"
	"strings"

	"google.golang.org/api/drive/v3"
)

var (
	knativeCommunity = File{
		id:   "0APnJ_hRs30R2Uk9PVA",
		name: "Knative Community Drive",
	}
	knativeTeam = File{
		id:   "0AM-QGZJ-HUA8Uk9PVA",
		name: "Knative Team Drive",
	}
)

const (
	FolderMimeType = "application/vnd.google-apps.folder"
)

func loadAllFiles(forest *Forest, srv *drive.Service, drive File) {
	forest.Add(File{
		id:       drive.id,
		name:     drive.name,
		isFolder: true,
	})
	var pageToken *string
	for pageToken == nil || *pageToken != "" {
		// List files from Knative Community Drive.
		request := srv.Files.List().PageSize(1000).
			Corpora("drive").
			IncludeItemsFromAllDrives(true).
			SupportsAllDrives(true).
			DriveId(drive.id).
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

	forest := NewForest()
	for _, drive := range []File{knativeTeam, knativeCommunity} {
		fmt.Println("Loading directory from ... ", drive.name)
		loadAllFiles(&forest, srv, drive)
	}

	for _, f := range forest.GetFiles() {
		fmt.Printf("[%s] %s\n", strings.Join(forest.GetPath(f), "/"), f.name)
	}
}
