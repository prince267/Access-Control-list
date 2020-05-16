import { Component, OnInit, Inject } from '@angular/core';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { DataService } from '../../data.service'
import { MatSnackBar } from '@angular/material/snack-bar';

interface ParentData {
  user_id: number,
  parent_path_name: string,
  parent_id: number
}

@Component({
  selector: 'app-file-folder-option-dialog',
  templateUrl: './file-folder-option-dialog.component.html',
  styleUrls: ['./file-folder-option-dialog.component.css']
})
export class FileFolderOptionDialogComponent implements OnInit {

  Entity: string;
  Options: string[] = ['Folder', 'File'];

  constructor(
    private _snackBar: MatSnackBar,
    public dialog: MatDialog,
    public dialogRef: MatDialogRef<FileFolderOptionDialogComponent>,
    private dataService: DataService,
    @Inject(MAT_DIALOG_DATA) public data: ParentData) { }

  ngOnInit(): void {
  }

  closeDialog() {
    this.dialogRef.close()
  }
  openSnackBar(message, action) {
    this._snackBar.open(message, action, {
      duration: 3000,
    });
  }

  async createEntity(EntityName: string, EntityType: string) {

    if (EntityType == undefined || EntityName == "") {
      this.openSnackBar("Please Choose Options and Path", " 😓")
      return
    }
    let EntityPathName = this.data.parent_path_name + "/" + EntityName
    if (EntityType == "Folder") {
      let FolderData = {
        "folder_name": EntityName,
        "path_name": EntityPathName
      }
      var FolderResponse = await this.dataService.CreateFolder(FolderData)
      if (FolderResponse.status == 200) {
        let UserFolderData = {
          "user_id": this.data.user_id,
          "parent_folder_id": this.data.parent_id,
          "child_folder_name": EntityName,
          "child_folder_id": FolderResponse.data,
          "permission_id": 2
        }
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
          "user_id": this.data.user_id,
          "parent_folder_id": this.data.parent_id,
          "child_file_name": EntityName,
          "child_file_id": FileResponse.data,
          "permission_id": 2
        }
        this.dataService.NewUserFile(UserFileData).subscribe(res => {
          this.openSnackBar("New File Created", " 🎉")
          this.dialogRef.close()
        })
      }
    }
    // console.log("insert into file/folder", [EntityName, EntityPathName])
    // console.log("insert into file/folder_infolder", [this.data.user_id, this.data.parent_id, EntityName, 212, 1])
  }
}
