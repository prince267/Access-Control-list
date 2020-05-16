package model

type FileInFolder struct {
	UserID           int64  `json:"user_id" column:"user_id"`
	ParentFolderID   int64  `json:"parent_folder_id" column:"parent_folder_id"`
	ParentFolderPath string `json:"parent_folder_path" column:"folders.path_name"`
	ChildFileName    string `json:"name" column:"child_file_name"`
	ChildFileID      int64  `json:"id" column:"child_file_id"`
	PermissionID     int64  `json:"permission_id" column:"file_in_folder.permission_id"`
	Description      string `json:"permission_descrp" column:"descrp"`
	PathName         string `json:"path_name" column:"files.path_name"`
	Type             string `json:"type" column:"files.type"`
}

func (fileInFolder *FileInFolder) FileInFolderTable() string {
	return "file_in_folder"
}

func (fileInFolder *FileInFolder) FoldersTable() string {
	return "folders"
}

func (fileInFolder *FileInFolder) FilesTable() string {
	return "files"
}

func (fileInFolder *FileInFolder) PermissionTable() string {
	return "permission"
}

func (fileInFolder *FileInFolder) String() string {
	return Stringify(fileInFolder)
}
