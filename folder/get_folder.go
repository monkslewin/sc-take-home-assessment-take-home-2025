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

func GetParentPath(folders []Folder, name string) string {
	var parent_path string
    for _, folder := range folders {
        if folder.Name == name {
            parent_path = folder.Paths
            break
        }
    }

	return parent_path
	
}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
    folders_by_org := f.GetFoldersByOrgID(orgID)
    child_folders := []Folder{}

    parent_path := GetParentPath(folders_by_org, name)

    if parent_path == "" {
        return child_folders
    }


    for _, folder := range folders_by_org {
        if strings.HasPrefix(folder.Paths, parent_path+".") {
            child_folders = append(child_folders, folder)
        }
    }

    return child_folders
}
