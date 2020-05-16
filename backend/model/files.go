package model

type Files struct {
	FileID   int64  `json:"file_id,omitempty" key:"primary" autoincr:"1" column:"file_id"`
	FileName string `json:"file_name" column:"file_name"`
	PathName string `json:"path_name" column:"path_name"`
}

func (files *Files) FilesTable() string {
	return "files"
}

func (files *Files) String() string {
	return Stringify(files)
}
