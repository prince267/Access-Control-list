package model

// UserIModel prototype of db model
type UserIModel interface {
	UserTable() string
	String() string
}

type AllFoldersIModel interface {
	FolderInFolderTable() string
	FoldersTable() string
	String() string
}

type AllFilesIModel interface {
	FileInFolderTable() string
	FoldersTable() string
	FilesTable() string
	String() string
}

type NewFileInFolderIModel interface {
	NewFileInFolderTable() string
	String() string
}

type NewFolderInFolderIModel interface {
	NewFolderInFolderTable() string
	String() string
}

type FileInFolderIModel interface {
	FileInFolderTable() string
	FoldersTable() string
	FilesTable() string
	PermissionTable() string
	String() string
}

type FilesIModel interface {
	FilesTable() string
	String() string
}
type FolderInFolderIModel interface {
	FolderInFolderTable() string
	PermissionTable() string
	FoldersTable() string
	String() string
}
type FoldersIModel interface {
	FoldersTable() string
	String() string
}
type GroupsIModel interface {
	GroupsTable() string
	String() string
}
type PermissionIModel interface {
	PermissionTable() string
	String() string
}
type UserGroupIModel interface {
	UserGroupTable() string
	GroupTable() string
	String() string
}

type GroupUsersIModel interface {
	UserGroupTable() string
	UserTable() string
	String() string
}
