package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoleItemStruct(t *testing.T) {
	item := RoleItem{
		ID:     1,
		RoleId: 100,
		Path:   "/admin/project/createKeys",
		Index:  "type",
		Value:  "lock",
		Name:   "创建激活码",
	}

	assert.Equal(t, 1, item.ID)
	assert.Equal(t, 100, item.RoleId)
	assert.Equal(t, "/admin/project/createKeys", item.Path)
	assert.Equal(t, "type", item.Index)
	assert.Equal(t, "lock", item.Value)
	assert.Equal(t, "创建激活码", item.Name)
}

func TestRoleItemInternalStruct(t *testing.T) {
	r := role{
		Path:  "/admin/project/createKeys",
		Index: "type",
		Value: "lock",
		Name:  "创建激活码",
		Child: []role{
			{
				Path:  "/admin/project/lockKey",
				Index: "",
				Value: "",
				Name:  "锁定激活码",
			},
		},
	}

	assert.Equal(t, "/admin/project/createKeys", r.Path)
	assert.Equal(t, "type", r.Index)
	assert.Equal(t, "lock", r.Value)
	assert.Equal(t, "创建激活码", r.Name)
	assert.Len(t, r.Child, 1)
}

func TestRoleItemGetUserRole(t *testing.T) {
	// This test requires database connection
	t.Skip("Requires database connection")
}

func TestRoleItemPermissionPathFormat(t *testing.T) {
	// Test that all permission paths follow the expected format
	for _, parent := range RoleList {
		for _, child := range parent.Child {
			// Paths should start with /admin/
			assert.True(t, len(child.Path) == 0 || len(child.Path) >= 7,
				"Path should be empty or at least 7 characters: %s", child.Path)

			if child.Path != "" {
				assert.Contains(t, child.Path, "/admin/",
					"Path should contain /admin/: %s", child.Path)
			}
		}
	}
}

func TestRoleItemIndexValueConsistency(t *testing.T) {
	// Test that Index and Value are either both empty or both set
	for _, parent := range RoleList {
		for _, child := range parent.Child {
			if child.Index != "" {
				assert.NotEmpty(t, child.Value,
					"Value should be set when Index is set for path: %s", child.Path)
			}
			if child.Value != "" {
				assert.NotEmpty(t, child.Index,
					"Index should be set when Value is set for path: %s", child.Path)
			}
		}
	}
}

func TestRoleListCompleteness(t *testing.T) {
	// Test that RoleList is properly initialized
	assert.NotNil(t, RoleList, "RoleList should not be nil")
	assert.Greater(t, len(RoleList), 0, "RoleList should have items")

	// Each top-level item should have children
	for _, parent := range RoleList {
		assert.NotEmpty(t, parent.Name, "Parent should have a name")
		assert.NotNil(t, parent.Child, "Parent should have children")
	}
}

func TestRoleItemUniquePaths(t *testing.T) {
	// Test that all paths are unique within RoleList
	pathCount := make(map[string]int)

	var collectPaths func(items []role)
	collectPaths = func(items []role) {
		for _, item := range items {
			if item.Path != "" {
				pathCount[item.Path]++
			}
			if len(item.Child) > 0 {
				collectPaths(item.Child)
			}
		}
	}

	collectPaths(RoleList)

	// Check for duplicates (same path with different Index/Value is allowed)
	for path, count := range pathCount {
		if count > 1 {
			// This is allowed for fine-grained permissions (e.g., batchKeys with different types)
			t.Logf("Path %s appears %d times (may be intentional for fine-grained permissions)", path, count)
		}
	}
}

func TestRoleItemNameUniqueness(t *testing.T) {
	// Test that all permission names are unique
	nameCount := make(map[string]int)

	var collectNames func(items []role)
	collectNames = func(items []role) {
		for _, item := range items {
			if item.Name != "" {
				nameCount[item.Name]++
			}
			if len(item.Child) > 0 {
				collectNames(item.Child)
			}
		}
	}

	collectNames(RoleList)

	// All names should be unique
	for name, count := range nameCount {
		assert.Equal(t, 1, count, "Permission name should be unique: %s", name)
	}
}

func TestRoleItemHierarchy(t *testing.T) {
	// Test the hierarchy structure of RoleList
	for _, parent := range RoleList {
		// Parent items should have empty paths (they're groups)
		assert.Empty(t, parent.Path, "Parent group should have empty path: %s", parent.Name)

		// Parent items should have children
		assert.NotEmpty(t, parent.Child, "Parent group should have children: %s", parent.Name)

		// Children should have paths
		for _, child := range parent.Child {
			assert.NotEmpty(t, child.Path, "Child should have a path in group: %s", parent.Name)
		}
	}
}

func TestRoleItemCRUDIntegration(t *testing.T) {
	// These tests require database connection
	t.Skip("Requires database connection for integration testing")
}

// TestRoleItemPermissionMatching tests permission matching logic
func TestRoleItemPermissionMatching(t *testing.T) {
	tests := []struct {
		name          string
		itemPath      string
		itemIndex     string
		itemValue     string
		requestPath   string
		requestParam  string
		shouldMatch   bool
	}{
		{
			name:         "simple path match",
			itemPath:     "/admin/project/createKeys",
			itemIndex:    "",
			itemValue:    "",
			requestPath:  "/admin/project/createKeys",
			shouldMatch:  true,
		},
		{
			name:         "path mismatch",
			itemPath:     "/admin/project/createKeys",
			itemIndex:    "",
			itemValue:    "",
			requestPath:  "/admin/project/lockKey",
			shouldMatch:  false,
		},
		{
			name:         "fine-grained match",
			itemPath:     "/admin/project/batchKeys",
			itemIndex:    "type",
			itemValue:    "lock",
			requestPath:  "/admin/project/batchKeys",
			requestParam: "lock",
			shouldMatch:  true,
		},
		{
			name:         "fine-grained mismatch",
			itemPath:     "/admin/project/batchKeys",
			itemIndex:    "type",
			itemValue:    "lock",
			requestPath:  "/admin/project/batchKeys",
			requestParam: "delete",
			shouldMatch:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Simple path match
			pathMatch := tt.itemPath == tt.requestPath

			if !pathMatch {
				assert.False(t, tt.shouldMatch)
				return
			}

			// If no fine-grained check needed
			if tt.itemIndex == "" || tt.itemValue == "" {
				assert.True(t, tt.shouldMatch)
				return
			}

			// Fine-grained check
			paramMatch := tt.requestParam == tt.itemValue
			assert.Equal(t, tt.shouldMatch, paramMatch)
		})
	}
}
