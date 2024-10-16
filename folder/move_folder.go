package folder

import (
    "fmt"
	"strings"

  )

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	  // Find the source folder
	  src_folder, err := f.findFolderByName(name)
	  if err != nil {
		  return nil, fmt.Errorf("error: Source folder does not exist")
	  }
  
	  // Find the destination folder
	  dst_folder, err := f.findFolderByName(dst)
	  if err != nil {
		  return nil, fmt.Errorf("error: Destination folder does not exist")
	  }
  
	  // Prevent moving folders between different organizations
	  if src_folder.OrgId != dst_folder.OrgId {
		  return nil, fmt.Errorf("error: Cannot move a folder to a different organization")
	  }
  
	  // Prevent moving a folder into itself or into its own subtree
	  if src_folder.Paths == dst_folder.Paths {
		  return nil, fmt.Errorf("error: Cannot move a folder into itself")
	  }
  
	  if strings.HasPrefix(dst_folder.Paths, src_folder.Paths + ".") {
		  return nil, fmt.Errorf("error: Cannot move a folder into its own subtree")
	  }
  
	  child_folders := f.GetAllChildFolders(src_folder.OrgId, src_folder.Name)
  
	  // Store the old path before updating the source folder's path
	  old_path := src_folder.Paths
	  new_path := dst_folder.Paths + "." + src_folder.Name
  
	  src_folder.Paths = new_path
	  f.updateFolderInList(src_folder)
  
	  // Update paths for all child folders
	  for i, folder := range child_folders {
		  // Replace the old parent path with the new path in each child folder
		  child_folders[i].Paths = strings.Replace(folder.Paths, old_path, new_path, 1)
		  f.updateFolderInList(child_folders[i]) // Ensure the updated child folder is reflected in the folder list
	  }
  
	  return []Folder{}, nil
  }
  
  func (f *driver) updateFolderInList(updated_folder Folder) {
	  for i, folder := range f.folders {
		  if folder.Name == updated_folder.Name && folder.OrgId == updated_folder.OrgId {
			  f.folders[i] = updated_folder
			  break
		  }
	  }
  }
  
  func (f *driver) findFolderByName(name string) (Folder, error) {
	  for _, folder := range f.folders {
		  if folder.Name == name {
			  return folder, nil
		  }
	  }
	  return Folder{}, fmt.Errorf("error: Folder '%s' does not exist", name)
  }
  


