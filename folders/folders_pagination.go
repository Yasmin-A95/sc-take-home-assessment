package folders_pagination

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

	// for i := 0; i < len(p); i++ {
	// 	rff := &FetchFolderResponse{Folders: p[i]}
	// } //TODO output this with the p[i][i+3].Id as the tokem
	ffr := &FetchPagenatedResponse{Folders: p} 
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

/*
TODO if I had more time 

- res.Folders for each 
- check if it's index % 2 == 0 
- and for every one that is
- pair that index with that index + 1 
- then add the key
- check if the resulting container is half the length of the original
- if not append the last one in its own lil thingo


- the key should be the id of the first res.folder[index]Id in the next batch, 
- in this way the next lot is ALWAYS pointing to the lot after it
- so thats the index which % 2 is true for
- in this way the last page should have null as its next_page id

*/