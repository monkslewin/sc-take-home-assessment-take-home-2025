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
		// Test 1
		{
			name:  "No folders for given OrgID", // Test when no folders match the orgID
			orgID: uuid.Must(uuid.NewV4()),      // A random orgID that won't match
			folders: []folder.Folder{ 
				{Name: "Folder1", OrgId: uuid.Must(uuid.NewV4()), Paths: "Folder1"},
				{Name: "Folder2", OrgId: uuid.Must(uuid.NewV4()), Paths: "Folder2"},
			},
			want: []folder.Folder{}, // Expect no folders to match
		},
		// Test 2
		{
			name:  "Folders with matching orgID", // Test when matches exist 
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID), 
			folders: []folder.Folder{ 
				{
					Name:  "Folder1",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), 
					Paths: "path/to/folder1",
				},
				{
					Name:  "Folder2",
					OrgId: uuid.Must(uuid.NewV4()), 
					Paths: "path/to/folder2",
				},
				{
					Name:  "Folder3",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), 
					Paths: "path/to/folder3",
				},
			},
			want: []folder.Folder{ 
				{
					Name:  "Folder1",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "path/to/folder1",
				},
				{
					Name:  "Folder3",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "path/to/folder3",
				},
			},
		},
		// Test 3
		{
			name: "No folder exists",
			orgID: uuid.Must(uuid.NewV4()), // Random orgID
			folders: []folder.Folder{},     // No folders present
			want:    []folder.Folder{},
		},
		// Test 4
		{
			name:  "All folders match orgID", 
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID), // DefaultOrgID for testing
			folders: []folder.Folder{
				{
					Name:  "Folder1",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "path/to/folder1",
				},
				{
					Name:  "Folder2",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "path/to/folder2",
				},
			},
			want: []folder.Folder{
				{
					Name:  "Folder1",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "path/to/folder1",
				},
				{
					Name:  "Folder2",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "path/to/folder2",
				},
			},
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // t.Run creates subsets of tests
			f := folder.NewDriver(tt.folders)
			got := f.GetFoldersByOrgID(tt.orgID)
			assert.Equal(t, tt.want, got)

		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel() // Handles tests simultaneously
	tests := [...]struct {
		name         string
		orgID        uuid.UUID
		folders      []folder.Folder
		parentFolder string
		want         []folder.Folder
	}{
		// Test 1
		{ 
			name: "No folders exist",
			orgID: uuid.Must(uuid.NewV4()), // Random orgID
			folders: []folder.Folder{},     // No folders present
			parentFolder: "parent",
			want:    []folder.Folder{},  // Expecting an empty list
		},
		// Test 2
		{ 
			name: "Child folders exist",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{Name: "parent", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "parent"},
				{Name: "child1", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "parent.child1"},
				{Name: "child2", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "parent.child1.child2"},
				{Name: "child3", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "random_folder.child3"}, // Should not include this one
			},
			parentFolder: "parent",
			want: []folder.Folder{
				{Name: "child1", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "parent.child1"},
				{Name: "child2", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "parent.child1.child2"},
			},
		},
		// Test 3
		{
			name: "Parent exists but no child folders",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{Name: "parent", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "parent"},
				{Name: "child1", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "random_folder.child1"}, // Not child folders of the parent folder
				{Name: "child2", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "random_folder.child1.child2"},
			},
			parentFolder: "parent",
			want: []folder.Folder{},  // Expecting an empty list
		},
		// Test 4
		{
			name: "Folders exist but none are children",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{Name: "parent", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "parent"},
				{Name: "child1", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "child1"},
				{Name: "child2", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "child1.child2"},
				{Name: "child3", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "child1.child2.child3"}, // Should not include this one
			},
			parentFolder: "parent",
			want: []folder.Folder{},  // Expecting an empty list
		},
		// Test 5
		{
			name: "Mid-level children exist",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{Name: "parent1", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "parent1"},
				{Name: "parent2", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "parent1.parent2"}, 
				{Name: "child1", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "parent1.parent2.child1"},
				{Name: "child2", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "parent1.parent2.child1.child2"},
			},
			parentFolder: "parent2", // Parent is not at the top level
			want: []folder.Folder{
				{Name: "child1", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "parent1.parent2.child1"},
				{Name: "child2", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "parent1.parent2.child1.child2"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // t.Run creates subsets of tests
			f := folder.NewDriver(tt.folders)
			got := f.GetAllChildFolders(tt.orgID, tt.parentFolder)

			// Validate the returned folders
			assert.Equal(t, tt.want, got)
		})
	}
}


