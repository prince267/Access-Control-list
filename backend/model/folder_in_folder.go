package model

type FolderInFolder struct {
	UserID           int64  `json:"user_id" column:"user_id"`
	ParentFolderID   int64  `json:"parent_folder_id" column:"parent_folder_id"`
	ParentFolderPath string `json:"parent_folder_path" column:"folders1.path_name"`
	ChildFolderName  string `json:"name" column:"child_folder_name"`
	ChildFolderID    int64  `json:"id" column:"child_folder_id"`
	PermissionID     int64  `json:"permission_id" column:"folder_in_folder.permission_id"`
	Description      string `json:"permission_descrp" column:"descrp"`
	PathName         string `json:"path_name" column:"folders.path_name"`
	Type             string `json:"type" column:"folders.type"`
}

func (folderInFolder *FolderInFolder) FolderInFolderTable() string {
	return "folder_in_folder"
}

func (folderInFolder *FolderInFolder) PermissionTable() string {
	return "permission"
}

func (folderInFolder *FolderInFolder) FoldersTable() string {
	return "folders"
}
func (folderInFolder *FolderInFolder) String() string {
	return Stringify(folderInFolder)
}
