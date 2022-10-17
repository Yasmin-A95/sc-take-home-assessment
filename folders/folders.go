package folders

import (
	"github.com/gofrs/uuid"
)

func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) { // called by main and gets passed the OrgID via FFR type
	// var ( // TODO use or discard these unuser variables
	// 	err error
	// 	f1  Folder
	// 	fs  []*Folder
	// )
	f := []Folder{} // an instance of the folder struct 
	r, _ := FetchAllFoldersByOrgID(req.OrgID) // calling the func, passing the OrgID and storing the output
	for _, v := range r { // for all the key val pairs in r
		f = append(f, *v) // put a pointer to each val in f
	}
	var fp []*Folder // init pointer folder
	for _, v1 := range f { // for each key val pair in f
		fp = append(fp, &v1) // append the address of val to fp
	}
	var ffr *FetchFolderResponse // TODO merge down to one line using := // ffr is assigned to the pointer of FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: fp} // then it's assigned to the address of FetchFolderResponse with something that looks similar to object destructuring
	return ffr, nil // return the object containing the list of folders 
}

func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) { // takes the orgID
	folders := GetSampleData() // calling the function that reads "sample.json" storing the results in a var

	resFolder := []*Folder{} // init a folder with pointer
	for _, folder := range folders { // map over the sample data and check to see if the orgID matches the one we passed to this function
		if folder.OrgId == orgID { 
			resFolder = append(resFolder, folder) // append those that match to resFolder
		} // TODO test functionality when there are varied OrgId's might be a good idea
	}
	
	return resFolder, nil // return resFolder // resFolder == sample.json as all OrgId are the same 
}
