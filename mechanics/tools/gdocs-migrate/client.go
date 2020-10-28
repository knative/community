package main

import (
	"log"

	"google.golang.org/api/drive/v3"
)

const (
	FolderMimeType = "application/vnd.google-apps.folder"
)

func getDrive(srv *drive.Service, driveId string) (*File, error) {
	r, err := srv.Drives.Get(driveId).Do()
	if err != nil {
		return nil, err
	}
	return &File{
		id:       r.Id,
		name:     r.Name,
		driveID:  driveId,
		isFolder: true,
	}, nil
}

func loadAllFiles(forest *Forest, srv *drive.Service, drive File) {
	forest.Add(drive)
	var pageToken *string
	for pageToken == nil || *pageToken != "" {
		request := srv.Files.List().PageSize(1000).
			Corpora("drive").
			IncludeItemsFromAllDrives(true).
			SupportsAllDrives(true).
			DriveId(drive.id).
			Fields("nextPageToken, files(id, name, parents, mimeType, trashed)").
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
			if i.Trashed {
				// Ignore trashed files and folders.
				continue
			}
			forest.Add(File{
				id:       i.Id,
				name:     i.Name,
				parentID: i.Parents[0],
				driveID:  drive.id,
				isFolder: i.MimeType == FolderMimeType,
			})
		}
		// Print progress.
		log.Printf("Loaded %d files, %d folders\n", len(forest.GetFiles()), forest.FolderCount())
	}
}

func createDir(srv *drive.Service, parentDir File, name string) (*File, error) {
	f, err := srv.Files.Create(&drive.File{
		Name:     name,
		Parents:  []string{parentDir.id},
		DriveId:  parentDir.driveID,
		MimeType: FolderMimeType,
	}).SupportsAllDrives(true).Do()
	if err != nil {
		return nil, err
	}
	return &File{
		id:       f.Id,
		name:     f.Name,
		parentID: parentDir.id,
		driveID:  parentDir.driveID,
		isFolder: true,
	}, nil
}

func copyFile(srv *drive.Service, source File, dstDir File) (*File, error) {
	copier := drive.NewFilesService(srv)
	f, err := copier.Copy(source.id, &drive.File{
		Name:    source.name,
		Parents: []string{dstDir.id},
		DriveId: dstDir.driveID,
	}).SupportsAllDrives(true).Do()
	if err != nil {
		return nil, err
	}
	return &File{
		id:       f.Id,
		name:     f.Name,
		parentID: dstDir.id,
		driveID:  dstDir.id,
	}, nil
}
