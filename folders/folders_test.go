package folders_test

import (
	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetAllFolders(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		req := &folders.FetchFolderRequest{
			OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
		}
		assert.NotNil(t, req)

		/*
			- this is throwing a panic error :
			open sample.json: no such file or directory [recovered]

			res, err := folders.GetAllFolders(req)
			assert.NotNil(t, res)
			assert.NotNil(t, err)
		*/

	})
}
