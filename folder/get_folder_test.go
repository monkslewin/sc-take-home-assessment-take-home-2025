package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel() // // Handles tests simultaneously 
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name:  "No folders for given OrgID", // Test when no folders match the orgID
			orgID: uuid.Must(uuid.NewV4()),      // A random orgID that won't match
			folders: []folder.Folder{ // Define test data (no folders will match)
				{Name: "Folder1", OrgId: uuid.Must(uuid.NewV4()), Paths: "Folder1"},
				{Name: "Folder2", OrgId: uuid.Must(uuid.NewV4()), Paths: "Folder2"},
			},
			want: []folder.Folder{}, // Expect no folders to match
		},
		// TODO: your tests here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // t.Run creates subsets of tests
			f := folder.NewDriver(tt.folders)
			got := f.GetFoldersByOrgID(tt.orgID)
			assert.Equal(t, tt.want, got)


		})
	}
}
