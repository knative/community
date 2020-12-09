package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/api/drive/v3"
)

const (
	folderMimeType = "application/vnd.google-apps.folder"
)

func driveAsFile(srv *drive.Service, driveId string) (*File, error) {
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

func loadAllFiles(forest *Forest, srv *drive.Service, drive *File) {
	forest.Add(drive)
	var pageToken *string
	for pageToken == nil || *pageToken != "" {
		request := srv.Files.List().PageSize(1000).
			Corpora("drive").
			IncludeItemsFromAllDrives(true).
			SupportsAllDrives(true).
			DriveId(drive.id).
			Fields("nextPageToken", "files(id, name, parents, mimeType, trashed)").
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
			forest.Add(&File{
				id:           i.Id,
				name:         i.Name,
				parentID:     i.Parents[0],
				driveID:      drive.id,
				isFolder:     i.MimeType == folderMimeType,
				mimeType:     i.MimeType,
				modifiedTime: i.ModifiedTime,
			})
		}
		// Print progress.
		log.Printf("Loaded %d files, %d folders\n", len(forest.GetFiles()), forest.FolderCount())
	}
}

func createDir(srv *drive.Service, parentDir *File, name string) (*File, error) {
	f, err := srv.Files.Create(&drive.File{
		Name:     name,
		Parents:  []string{parentDir.id},
		DriveId:  parentDir.driveID,
		MimeType: folderMimeType,
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

func deleteFile(srv *drive.Service, fileId string) error {
	return srv.Files.Delete(fileId).SupportsAllDrives(true).Do()
}

func prependAuthor(author *drive.User, content string) string {
	return author.DisplayName + " wrote: " + content
}

func sanitizeEmail(content string) string {
	// Use a homoglyph of @ to make sure people aren't getting emails
	return strings.ReplaceAll(content, "@", "\uff20")
}

// reformat the time from RFC3339 for brevity
func decorateTime(rfc3339 string) string {
	t, _ := time.Parse(time.RFC3339, rfc3339)
	return t.UTC().Format("2006-Jan-02 23:04")
}

func decorateComment(input *drive.Comment) *drive.Comment {
	// Drive API expects only the content, anchor, and quotedFileContent fields to be set on create.
	return &drive.Comment{
		Anchor: input.Anchor,
		Content: fmt.Sprintf("[%s] %s",
			decorateTime(input.CreatedTime),
			prependAuthor(input.Author, sanitizeEmail(input.Content))),
		QuotedFileContent: input.QuotedFileContent,
	}
}

func decorateReply(input *drive.Reply) *drive.Reply {
	return &drive.Reply{
		Action: input.Action,
		Content: fmt.Sprintf("[%s] %s",
			decorateTime(input.CreatedTime),
			prependAuthor(input.Author, sanitizeEmail(input.Content))),
	}
}

func copyFile(srv *drive.Service, source *File, dstDir *File) (*File, error) {
	f, err := srv.Files.Copy(source.id, &drive.File{
		Name:         source.name,
		ModifiedTime: source.modifiedTime,
		Parents:      []string{dstDir.id},
		DriveId:      dstDir.driveID,
	}).SupportsAllDrives(true).Do()
	if err != nil {
		return nil, err
	}
	// Copy all the comments.
	// TODO: page them appropriately.
	comments, err := srv.Comments.List(source.id).Fields("*").Do()
	if err != nil {
		deleteFile(srv, f.Id)
		return nil, errors.Wrap(err, "CommentsList failed")
	}
	for _, comment := range comments.Comments {
		newComment, err := srv.Comments.Create(f.Id, decorateComment(comment)).Fields("*").Do()
		if err != nil {
			deleteFile(srv, f.Id)
			return nil, errors.Wrap(err, "CommentCreate failed")
		}
		for _, reply := range comment.Replies {
			if _, err := srv.Replies.Create(f.Id, newComment.Id, decorateReply(reply)).Fields("*").Do(); err != nil {
				deleteFile(srv, f.Id)
				return nil, errors.Wrap(err, "ReplyCreate failed")
			}
		}
	}
	return &File{
		id:       f.Id,
		name:     f.Name,
		parentID: dstDir.id,
		driveID:  dstDir.id,
	}, nil
}
