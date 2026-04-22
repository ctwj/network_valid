package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPageCount(t *testing.T) {
	tests := []struct {
		name            string
		count           int64
		pageSize        int64
		page            int64
		expectedTotal   int64
		expectedOffset  int64
		expectedCurrent int64
	}{
		{
			name:            "normal pagination",
			count:           100,
			pageSize:        10,
			page:            1,
			expectedTotal:   10,
			expectedOffset:  0,
			expectedCurrent: 1,
		},
		{
			name:            "second page",
			count:           100,
			pageSize:        10,
			page:            2,
			expectedTotal:   10,
			expectedOffset:  10,
			expectedCurrent: 2,
		},
		{
			name:            "page exceeds total",
			count:           100,
			pageSize:        10,
			page:            20,
			expectedTotal:   10,
			expectedOffset:  90,
			expectedCurrent: 10,
		},
		{
			name:            "partial last page",
			count:           95,
			pageSize:        10,
			page:            10,
			expectedTotal:   10,
			expectedOffset:  90,
			expectedCurrent: 10,
		},
		{
			name:            "single item",
			count:           1,
			pageSize:        10,
			page:            1,
			expectedTotal:   1,
			expectedOffset:  0,
			expectedCurrent: 1,
		},
		{
			name:            "empty result",
			count:           0,
			pageSize:        10,
			page:            1,
			expectedTotal:   0,
			expectedOffset:  0,
			expectedCurrent: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			totalPage, offset, currentPage := PageCount(tt.count, tt.pageSize, tt.page)
			assert.Equal(t, tt.expectedTotal, totalPage, "total page mismatch")
			assert.Equal(t, tt.expectedOffset, offset, "offset mismatch")
			assert.Equal(t, tt.expectedCurrent, currentPage, "current page mismatch")
		})
	}
}
