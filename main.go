package main

import (
	"github.com/georgechieng-sc/interns-2022/folder"
	//"github.com/gofrs/uuid"
)

func main() {
	// Correct the variable name and ensure it is used correctly

	// Get the sample data (folders)
	res := folder.GetAllFolders()


	// Create the driver instance with the folder data
	folderDriver := folder.NewDriver(res)

	// Get child folders by orgID and folder name
	test, _ := folderDriver.MoveFolder("clear-arclight", "model-stargirl")

	// Print the child folders if needed (assuming PrettyPrint is implemented in folder package)
	folder.PrettyPrint(test)
}

