package folders

import (
	"github.com/gofrs/uuid"
)

// TODO move files back to correct spots in the tree .

func GetAllFolders(req *FetchFolderRequest) (*FetchPagenatedResponse, error) {
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
	for i := 0; i < (len(fp) / 2); i++ {
		p = append(p, fp[i: i+2])
	}

	for i := 0; i < len(p); i++ {
		rff := &FetchFolderResponse{Folders: p[i]}
		PrettyPrint(rff)
	}
	ffr := &FetchPagenatedResponse{Folders: p} 
	PrettyPrint(ffr)
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
	return resFolder, nil 
}

