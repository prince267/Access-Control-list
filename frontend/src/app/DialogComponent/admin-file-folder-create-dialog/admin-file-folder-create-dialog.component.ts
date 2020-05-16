import { Component, OnInit, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { DataService } from '../../data.service'
import { response } from '../../models/model'
import { MatSnackBar } from '@angular/material/snack-bar';

interface folderData {
  id: number,
  path_name: string
}

@Component({
  selector: 'app-admin-file-folder-create-dialog',
  templateUrl: './admin-file-folder-create-dialog.component.html',
  styleUrls: ['./admin-file-folder-create-dialog.component.css']
})
export class AdminFileFolderCreateDialogComponent implements OnInit {

  constructor(
    private _snackBar: MatSnackBar,
    private dataService: DataService,
    public dialogRef: MatDialogRef<AdminFileFolderCreateDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: folderData
  ) { }

  allUsers = []
  Entity: string;
  Options: string[] = ['Folder', 'File'];
  selectedValue: number
  permission: number;
  permissions: number[] = [1, 2];

  ngOnInit(): void {
    this.dataService.GetAllUser().subscribe((res: response) => {
      this.allUsers = res.data
      console.log(this.allUsers)
    })

    console.log(this.data)
  }

  openSnackBar(message, action) {
    this._snackBar.open(message, action, {
      duration: 3000,
    });
  }
  
  async createEntity(UserId,EntityType,EntityName,permission){

    let EntityPathName = this.data.path_name + "/" + EntityName
    if (EntityType == "Folder") {
      let FolderData = {
        "folder_name": EntityName,
        "path_name": EntityPathName
      }
      var FolderResponse = await this.dataService.CreateFolder(FolderData)
      if (FolderResponse.status == 200) {
        let UserFolderData = {
          "user_id": UserId,
          "parent_folder_id": this.data.id,
          "child_folder_name": EntityName,
          "child_folder_id": FolderResponse.data,
          "permission_id": permission
        }
        console.log(UserFolderData)
        this.dataService.NewUserFolder(UserFolderData).subscribe(res => {
          this.openSnackBar("New Folder Created", " 🎉")
          this.dialogRef.close()
        })
      }
    }
    if (EntityType == "File") {
      let FileData = {
        "file_name": EntityName,
        "path_name": EntityPathName
      }
      var FileResponse = await this.dataService.CreateFile(FileData)
      if (FileResponse.status == 200) {
        let UserFileData = {
          "user_id": UserId,
          "parent_folder_id": this.data.id,
          "child_file_name": EntityName,
          "child_file_id": FileResponse.data,
          "permission_id": permission
        }
        console.log(UserFileData)
        this.dataService.NewUserFile(UserFileData).subscribe(res => {
          this.openSnackBar("New File Created", " 🎉")
          this.dialogRef.close()
        })
      }
    }
  }
}
