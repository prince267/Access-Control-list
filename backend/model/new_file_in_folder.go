package model

type NewFileInFolder struct {
	UserID         int64  `json:"user_id" key:"primary" column:"user_id"`
	ParentFolderID int64  `json:"parent_folder_id" key:"primary" column:"parent_folder_id"`
	ChildFileName  string `json:"child_file_name"  column:"child_file_name"`
	ChildFileID    int64  `json:"child_file_id" key:"primary" column:"child_file_id"`
	PermissionID   int64  `json:"permission_id" column:"permission_id"`
}

func (newFileInFolder *NewFileInFolder) NewFileInFolderTable() string {
	return "file_in_folder"
}

func (newFileInFolder *NewFileInFolder) String() string {
	return Stringify(newFileInFolder)
}
