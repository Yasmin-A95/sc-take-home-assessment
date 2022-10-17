package folders_test

import (
	"testing"
	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/stretchr/testify/assert"
	"github.com/gofrs/uuid"
)

func Test_GetAllFolders(t *testing.T) {
	req := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
	}
	assert.NotNil(t, req)
	// t.Run("test", func(t *testing.T) {
	// 	assert.NotNil(folders.GetAllFolders(req))
	// })
}
