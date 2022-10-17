package folders

import (
	"github.com/gofrs/uuid"
)

// TODO move files back to correct spots in the tree .
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	f := []Folder{}
	r, _ := FetchAllFoldersByOrgID(req.OrgID)
	for _, v := range r {
		f = append(f, *v)
	}
	var fp []*Folder
	for _, v1 := range f {
		fp = append(fp, &v1)
	}
	var p [][]*Folder
	for i := 0; i < (len(fp) /2); i++ {
		p = append(p, fp[i: i+2])
	}
	PrettyPrint(p)
	ffr := &FetchFolderResponse{Folders: fp}
	return ffr, nil
}

func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil // return resFolder // resFolder == sample.json as all OrgId are the same
}

// Copy over the `GetFolders` and `FetchAllFoldersByOrgID` to get started
