package main

type File struct {
	id           string
	name         string
	parentID     string
	driveID      string
	isFolder     bool
	mimeType     string
	modifiedTime string
}

type Forest struct {
	fileByID map[string]File
	children map[string]map[string]File
	leaves   []File
}

func NewForest() Forest {
	return Forest{
		fileByID: make(map[string]File),
		children: make(map[string]map[string]File),
	}
}

func (fo *Forest) Add(f *File) {
	fo.fileByID[f.id] = *f
	if fo.children[f.parentID] == nil {
		fo.children[f.parentID] = make(map[string]File)
	}
	fo.children[f.parentID][f.name] = *f
	if !f.isFolder {
		fo.leaves = append(fo.leaves, *f)
	}
}

func (fo *Forest) GetFiles() []File {
	return fo.leaves
}

func (fo *Forest) Size() int {
	return len(fo.fileByID)
}

func (fo *Forest) FolderCount() int {
	return fo.Size() - len(fo.GetFiles())
}

// GetPath returns all the ancestor of a given File in decreasing
// seniority.
//
// This should only be called after the whole forest has ben loaded,
// i.e. Add(...) has been called for all the Files.
func (fo *Forest) GetPath(f *File) []string {
	parents := []string{}
	cur := *f
	for cur.parentID != "" {
		cur = fo.fileByID[cur.parentID]
		parents = append(parents, cur.name)
	}
	return reverse(parents)
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (fo *Forest) Children(dir *File) map[string]File {
	return fo.children[dir.id]
}
