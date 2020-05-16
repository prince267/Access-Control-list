// function FilesMapping(files) {
//     var map = {}
//     for (let i = 0; i < files.length; i++) {
//         let obj = files[i]
//         if (!(obj.parent_folder_id in map)) {
//             map[obj.parent_folder_id] = {}
//             map[obj.parent_folder_id].files = []
//         }

//     }
//     for (let i = 0; i < files.length; i++) {
//         let obj = files[i]
//         map[obj.parent_folder_id].files.push(obj)
//     }

//     return map;
// }

// function FolderFileMapping(Folders, FilesMap) {
//     for (let i = 0; i < Folders.length; i++) {
//         if (Folders[i].child_folder_id in FilesMap) {
//             Folders[i].files = FilesMap[Folders[i].child_folder_id].files
//         }
//         else {
//             Folders[i].files = []
//         }
//     }
//     return Folders
// }

function FolderInFolderMapping(FileFolder, HeadParent) {
    var map = {}
    for (var i = 0; i < FileFolder.length; i++) {
        var obj = FileFolder[i]
        if (!(obj.id in map)) {
            map[obj.id] = obj
            map[obj.id].children = []
        }

        if (typeof map[obj.id].name == 'undefined') {
            map[obj.id].user_id = obj.user_id
            map[obj.id].parent_folder_id = obj.parent_folder_id
            map[obj.id].parent_folder_path = obj.parent_folder_path
            map[obj.id].name = obj.name
            map[obj.id].id = obj.id
            map[obj.id].permission_id = obj.permission_id
            map[obj.id].permission_descrp = obj.permission_descrp
            map[obj.id].path_name = obj.path_name
            map[obj.id].type = obj.type
        }

        if (HeadParent.indexOf(obj.parent_folder_id) > -1) {
            var parent:string = '-'
        }
        else {
            var parent:string = obj.parent_folder_id

        }

        if (!(parent in map)) {
            map[parent] = {}
            map[parent].children = []
        }

        map[parent].children.push(map[obj.id])
    }
    return map['-']
}

function FilesFolderRelation(files, folder) {
    var FileFolder=files.data.concat(folder.data)
    if(FileFolder.length==0){
        return[]
    }
    // var FilesMap = FilesMapping(files.data)
    // var FileFolderMap = FolderFileMapping(folder.data, FilesMap)
    // // FolderInFolderMapping(FileFolderMap,FilesMap)
    var ChildFolderId = []
    var ParentFolderId = []
    for (let i = 0; i < FileFolder.length; i++) {
        ChildFolderId = ChildFolderId.concat(FileFolder[i].id)
        ParentFolderId = ParentFolderId.concat(FileFolder[i].parent_folder_id)
        // if (!(FileFolder[i].parent_folder_id in ParentFolderId)) {
        //     ParentFolderId = ParentFolderId.concat(FileFolder[i].parent_folder_id)
        // }
    }
    var HeadParent = []
    for (let i = 0; i < ParentFolderId.length; i++) {
        if (ChildFolderId.indexOf(ParentFolderId[i]) == -1) {
            HeadParent = HeadParent.concat(ParentFolderId[i])
        }
    }
    var FileFolderMap = FolderInFolderMapping(FileFolder, HeadParent)

    return FileFolderMap.children
    // console.log("***  ",JSON.stringify(FileFolder))
}

export { FilesFolderRelation }