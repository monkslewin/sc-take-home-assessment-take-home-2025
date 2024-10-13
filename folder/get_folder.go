package folder

import "github.com/gofrs/uuid"

import (
	"strings"
  )

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res
}

func GetParentPath(folders []Folder) string {
	parent_path := 
}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
    folders_by_org := f.GetFoldersByOrgID(orgID)
    child_folders := []Folder{}

    var parentPath string
    for _, folder := range folders_by_org {
        if folder.Name == name {
            parentPath = folder.Paths
            break
        }
    }

    if parentPath == "" {
        return child_folders
    }


    for _, folder := range folders_by_org {
        if strings.HasPrefix(folder.Paths, parentPath+".") {
            child_folders = append(child_folders, folder)
        }
    }

    return child_folders
}
