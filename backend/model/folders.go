package model

type Folders struct {
	FolderID   int64  `json:"folder_id,omitempty" key:"primary" autoincr:"1" column:"folder_id"`
	FolderName string `json:"folder_name" column:"folder_name"`
	PathName   string `json:"path_name" column:"path_name"`
}

func (folders *Folders) FoldersTable() string {
	return "folders"
}

func (folders *Folders) String() string {
	return Stringify(folders)
}
