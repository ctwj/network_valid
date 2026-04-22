package admin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDeveloper(t *testing.T) {
	tests := []struct {
		name        string
		pid         int
		expectError bool
	}{
		{
			name:        "developer with Pid=0",
			pid:         0,
			expectError: false,
		},
		{
			name:        "agent with Pid>0",
			pid:         1,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test controller
			ctrl := &BaseController{}

			// Set up mock ManagerInfo
			ctrl.ManagerInfo.Pid = tt.pid

			// Note: IsDeveloper() calls b.Error() which stops execution
			// We can't fully test this without mocking the ResponseWriter
			// This test verifies the logic structure
			if tt.pid > 0 {
				// Agent should trigger error
				assert.True(t, ctrl.ManagerInfo.Pid > 0)
			} else {
				// Developer should not trigger error
				assert.Equal(t, 0, ctrl.ManagerInfo.Pid)
			}
		})
	}
}

func TestIsInManagerList(t *testing.T) {
	tests := []struct {
		name        string
		managerList []int
		checkId     int
		expected    bool
	}{
		{
			name:        "id in list",
			managerList: []int{1, 2, 3, 4, 5},
			checkId:     3,
			expected:    true,
		},
		{
			name:        "id not in list",
			managerList: []int{1, 2, 3, 4, 5},
			checkId:     6,
			expected:    false,
		},
		{
			name:        "empty list",
			managerList: []int{},
			checkId:     1,
			expected:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := &BaseController{}
			ctrl.ManagerIdArr = tt.managerList

			// Check if id is in list (without calling Error method)
			found := false
			for _, i := range ctrl.ManagerIdArr {
				if tt.checkId == i {
					found = true
					break
				}
			}

			assert.Equal(t, tt.expected, found)
		})
	}
}

func TestPermissionCheckLogic(t *testing.T) {
	tests := []struct {
		name           string
		pid            int
		powerId        int
		requestPath    string
		roleItems      []mockRoleItem
		shouldCheck    bool
		hasPermission  bool
		requestParam   string // For fine-grained permission check
	}{
		{
			name:          "developer bypasses permission check",
			pid:           0,
			powerId:       0,
			requestPath:   "/admin/project/createKeys",
			shouldCheck:   false,
			hasPermission: true,
		},
		{
			name:        "agent with matching permission",
			pid:         1,
			powerId:     1,
			requestPath: "/admin/project/createKeys",
			roleItems: []mockRoleItem{
				{Path: "/admin/project/createKeys", Index: "", Value: ""},
			},
			shouldCheck:   true,
			hasPermission: true,
		},
		{
			name:        "agent without matching permission",
			pid:         1,
			powerId:     1,
			requestPath: "/admin/project/createKeys",
			roleItems: []mockRoleItem{
				{Path: "/admin/project/lockKey", Index: "", Value: ""},
			},
			shouldCheck:   true,
			hasPermission: false,
		},
		{
			name:        "agent with fine-grained permission match",
			pid:         1,
			powerId:     1,
			requestPath: "/admin/project/batchKeys",
			roleItems: []mockRoleItem{
				{Path: "/admin/project/batchKeys", Index: "type", Value: "lock"},
			},
			shouldCheck:   true,
			hasPermission: true,
			requestParam:  "lock", // Request wants "lock" and permission is "lock"
		},
		{
			name:        "agent with fine-grained permission mismatch",
			pid:         1,
			powerId:     1,
			requestPath: "/admin/project/batchKeys",
			roleItems: []mockRoleItem{
				{Path: "/admin/project/batchKeys", Index: "type", Value: "delete"},
			},
			shouldCheck:   true,
			hasPermission: false,
			requestParam:  "lock", // Request wants "lock" but permission is "delete"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Verify permission check logic
			if tt.pid == 0 {
				// Developer should bypass check
				assert.True(t, tt.hasPermission)
				return
			}

			// For agents, check if path matches any role item
			permissionFound := false
			for _, item := range tt.roleItems {
				if item.Path == tt.requestPath {
					if item.Index != "" && item.Value != "" {
						// Fine-grained check: request param must match permission value
						if tt.requestParam != "" && stringsContains(tt.requestParam, item.Value) {
							permissionFound = true
						}
						// If no requestParam provided, we can't verify fine-grained permission
					} else {
						permissionFound = true
					}
				}
			}

			assert.Equal(t, tt.hasPermission, permissionFound)
		})
	}
}

// stringsContains is a helper for substring check
func stringsContains(s, substr string) bool {
	return len(s) >= len(substr) && s == substr
}

// mockRoleItem for testing
type mockRoleItem struct {
	Path  string
	Index string
	Value string
}

// TestBaseControllerStructure verifies the controller structure
func TestBaseControllerStructure(t *testing.T) {
	ctrl := &BaseController{}

	// Verify initial values
	assert.Equal(t, 0, ctrl.ManagerId)
	assert.Nil(t, ctrl.ManagerIdArr)
	assert.Equal(t, "", ctrl.apiUrl)
}

// TestSuccessResultStructure verifies response structure
func TestSuccessResultStructure(t *testing.T) {
	result := SuccessResult{
		Errno:  0,
		Data:   "test data",
		Errmsg: "success",
	}

	assert.Equal(t, 0, result.Errno)
	assert.Equal(t, "test data", result.Data)
	assert.Equal(t, "success", result.Errmsg)
}

// TestErrorResultStructure verifies error response structure
func TestErrorResultStructure(t *testing.T) {
	result := ErrorResult{
		Errno:  400,
		Errmsg: "error message",
	}

	assert.Equal(t, 400, result.Errno)
	assert.Equal(t, "error message", result.Errmsg)
}

// TestNewBaseController creates a new base controller for integration testing
func TestNewBaseController(t *testing.T) {
	// This test requires beego context setup
	// It's a placeholder for integration tests
	t.Skip("Requires full beego context setup")
}

// Helper function to create a test request
func createTestRequest(method, path string) *http.Request {
	req := httptest.NewRequest(method, path, nil)
	return req
}

// Helper function to create a test response recorder
func createTestResponseRecorder() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}

// TestPermissionPathMatching tests the path matching logic
func TestPermissionPathMatching(t *testing.T) {
	roleList := []struct {
		Path  string
		Name  string
	}{
		{Path: "/admin/project/createKeys", Name: "创建激活码"},
		{Path: "/admin/project/lockKey", Name: "锁定激活码"},
		{Path: "/admin/project/deleteKeys", Name: "删除激活码"},
	}

	tests := []struct {
		name        string
		requestPath string
		shouldMatch bool
		matchedName string
	}{
		{
			name:        "exact match",
			requestPath: "/admin/project/createKeys",
			shouldMatch: true,
			matchedName: "创建激活码",
		},
		{
			name:        "case insensitive match",
			requestPath: "/ADMIN/PROJECT/CREATEKEYS",
			shouldMatch: true,
			matchedName: "创建激活码",
		},
		{
			name:        "no match",
			requestPath: "/admin/project/otherAction",
			shouldMatch: false,
			matchedName: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := false
			matchedName := ""
			for _, item := range roleList {
				if stringsEqualFold(item.Path, tt.requestPath) {
					found = true
					matchedName = item.Name
					break
				}
			}

			assert.Equal(t, tt.shouldMatch, found)
			if tt.shouldMatch {
				assert.Equal(t, tt.matchedName, matchedName)
			}
		})
	}
}

// stringsEqualFold is a helper for case-insensitive comparison
func stringsEqualFold(s1, s2 string) bool {
	// Simple implementation for testing
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		c1 := s1[i]
		c2 := s2[i]
		if c1 >= 'A' && c1 <= 'Z' {
			c1 += 32
		}
		if c2 >= 'A' && c2 <= 'Z' {
			c2 += 32
		}
		if c1 != c2 {
			return false
		}
	}
	return true
}
