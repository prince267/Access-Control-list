package model

type AllFolders struct {
	ParentFolderID   int64  `json:"parent_folder_id" column:"parent_folder_id"`
	ParentFolderPath string `json:"parent_folder_path" column:"folders1.path_name"`
	ChildFolderName  string `json:"name" column:"child_folder_name"`
	ChildFolderID    int64  `json:"id" column:"child_folder_id"`
	PathName         string `json:"path_name" column:"folders.path_name"`
	Type             string `json:"type" column:"folders.type"`
}

func (allFolders *AllFolders) FolderInFolderTable() string {
	return "folder_in_folder"
}

func (allFolders *AllFolders) FoldersTable() string {
	return "folders"
}
func (allFolders *AllFolders) String() string {
	return Stringify(allFolders)
}
