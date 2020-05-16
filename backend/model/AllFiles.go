package model

type AllFiles struct {
	ParentFolderID   int64  `json:"parent_folder_id" column:"parent_folder_id"`
	ParentFolderPath string `json:"parent_folder_path" column:"folders.path_name"`
	ChildFileName    string `json:"name" column:"child_file_name"`
	ChildFileID      int64  `json:"id" column:"child_file_id"`
	PathName         string `json:"path_name" column:"files.path_name"`
	Type             string `json:"type" column:"files.type"`
}

func (allFiles *AllFiles) FileInFolderTable() string {
	return "file_in_folder"
}

func (allFiles *AllFiles) FoldersTable() string {
	return "folders"
}

func (allFiles *AllFiles) FilesTable() string {
	return "files"
}

func (allFiles *AllFiles) String() string {
	return Stringify(allFiles)
}
