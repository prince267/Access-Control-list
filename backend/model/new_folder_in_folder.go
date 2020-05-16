package model

type NewFolderInFolder struct {
	UserID          int64  `json:"user_id" key:"primary" column:"user_id"`
	ParentFolderID  int64  `json:"parent_folder_id" key:"primary" column:"parent_folder_id"`
	ChildFolderName string `json:"child_folder_name" key:"primary" column:"child_folder_name"`
	ChildFolderID   int64  `json:"child_folder_id" key:"primary" column:"child_folder_id"`
	PermissionID    int64  `json:"permission_id" column:"permission_id"`
}

func (newFolderInFolder *NewFolderInFolder) NewFolderInFolderTable() string {
	return "folder_in_folder"
}

func (newFolderInFolder *NewFolderInFolder) String() string {
	return Stringify(newFolderInFolder)
}
