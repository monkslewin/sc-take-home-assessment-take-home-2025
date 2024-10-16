package folder

import (
    "fmt"
	"strings"

  )

  func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	srcFolder, err := f.findFolderByName(name)
	if err != nil {
		return nil, err
	}

	dstFolder, err := f.findFolderByName(dst)
	if err != nil {
		return nil, err
	}

	// Prevent moving a folder into its own subtree
	if strings.HasPrefix(dstFolder.Paths, srcFolder.Paths + ".") {
		return nil, fmt.Errorf("error: Cannot move a folder into its own subtree")
	}

	// Get all child folders
	fmt.Println("Source folder:", srcFolder)
	fmt.Println("Destination folder:", dstFolder)
	childFolders := f.GetAllChildFolders(srcFolder.OrgId, srcFolder.Name)
	
	// Update paths

	newPath := dstFolder.Paths + "." + srcFolder.Name
	srcFolder.Paths = newPath

	oldPath := srcFolder.Paths

	for i, folder := range childFolders {
		// Replace the old parent path with the new path in each child folder
		childFolders[i].Paths = strings.Replace(folder.Paths, oldPath, newPath, 1)
	}

	return []Folder{}, nil
}


func (f *driver) findFolderByName(name string) (Folder, error) {
	for _, folder := range f.folders {
		if folder.Name == name {
			return folder, nil
		}
	}
	return Folder{}, fmt.Errorf("Folder '%s' not found", name)
}
