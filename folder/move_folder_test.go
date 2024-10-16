package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel() 
	tests := [...]struct {
		name        string
		src         string
		dst         string
		orgID       uuid.UUID
		folders     []folder.Folder
		wantFolders []folder.Folder
		wantError   string
	}{
		// Test 1: Moving "bravo" to "delta"
		{
			name:  "Move bravo to delta",
			src:   "bravo",
			dst:   "delta",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "golf", Paths: "golf", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
			},
			wantFolders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "bravo", Paths: "alpha.delta.bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)}, // Correct full path
				{Name: "charlie", Paths: "alpha.delta.bravo.charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)}, // Correct full path
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "golf", Paths: "golf", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
			},
			wantError: "",
		},

		{
			name:  "Move bravo to golf",
			src:   "bravo",
			dst:   "golf",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "golf", Paths: "golf", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
			},
			wantFolders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "bravo", Paths: "golf.bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)}, // Moved to golf
				{Name: "charlie", Paths: "golf.bravo.charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)}, // Path updated accordingly
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
				{Name: "golf", Paths: "golf", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)},
			},
			wantError: "",
		},
		

		
		// Test 2: Error when moving a folder to a child of itself
		{
			name:      "Move bravo to charlie should return an error",
			src:       "bravo",
			dst:       "charlie",
			orgID:     uuid.FromStringOrNil(folder.DefaultOrgID),
			folders:   []folder.Folder{{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)}, {Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)}},
			wantError: "error: Cannot move a folder into its own subtree",
		},
		// Test 3: Error when moving a folder to itself
		{
			name:      "Move bravo to itself should return an error",
			src:       "bravo",
			dst:       "bravo",
			orgID:     uuid.FromStringOrNil(folder.DefaultOrgID),
			folders:   []folder.Folder{{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),}},
			wantError: "error: Cannot move a folder into itself",
		},
		// Test 4: Error when moving a folder between different organizations
		{
			name:      "Move bravo to a different organization should return an error",
			src:       "bravo",
			dst:       "foxtrot",
			orgID:     uuid.Must(uuid.NewV4()), // Different orgID
			folders:   []folder.Folder{{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)}, {Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.Must(uuid.NewV4())}},
			wantError: "error: Cannot move a folder to a different organization",
		},
		// Test 5: Error when source folder does not exist
		{
			name:      "Move invalid folder to delta should return an error",
			src:       "invalid_folder",
			dst:       "delta",
			orgID:     uuid.FromStringOrNil(folder.DefaultOrgID),
			folders:   []folder.Folder{{Name: "delta", Paths: "alpha.delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)}},
			wantError: "error: Source folder does not exist",
		},
		// Test 6: Error when destination folder does not exist
		{
			name:      "Move bravo to invalid folder should return an error",
			src:       "bravo",
			dst:       "invalid_folder",
			orgID:     uuid.FromStringOrNil(folder.DefaultOrgID),
			folders:   []folder.Folder{{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID)}},
			wantError: "error: Destination folder does not exist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // t.Run creates subsets of tests
			f := folder.NewDriver(tt.folders)
			_, err := f.MoveFolder(tt.src, tt.dst)

			if tt.wantError != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.wantError, err.Error())
			} else {
				assert.NoError(t, err)
				got := f.GetFoldersByOrgID(tt.orgID) 
				assert.Equal(t, tt.wantFolders, got)
			}
		})
	}
}
