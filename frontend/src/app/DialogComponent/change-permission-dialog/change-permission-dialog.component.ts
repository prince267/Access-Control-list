import { Component, OnInit, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { FileFolderNode } from '../../models/model'
import { DataService } from '../../data.service'
import {MatSnackBar} from '@angular/material/snack-bar';

@Component({
  selector: 'app-change-permission-dialog',
  templateUrl: './change-permission-dialog.component.html',
  styleUrls: ['./change-permission-dialog.component.css']
})
export class ChangePermissionDialogComponent implements OnInit {

  constructor(
    private _snackBar: MatSnackBar,
    private dataService: DataService,
    public dialogRef: MatDialogRef<ChangePermissionDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: FileFolderNode
  ) { }

  ngOnInit(): void {
  }

  openSnackBar(message, action) {
    this._snackBar.open(message, action, {
      duration: 3000,
    });
  }

  traverse(node, id) {
    for (var i in node) {
      if (!!node[i] && typeof (node[i]) == "object") {
        if (i !== "children") {
          if (node[i].type == "Folder") {
            let FolderData = {
              "user_id": node[i].user_id,
              "parent_folder_id": node[i].parent_folder_id,
              "child_folder_name": node[i].name,
              "child_folder_id": node[i].id,
              "permission_id": id
            }
            this.dataService.updateUserFolderPermission(FolderData).subscribe(res => {

            })
          }
          if (node[i].type == "File") {
            let FileData = {
              "user_id": node[i].user_id,
              "parent_folder_id": node[i].parent_folder_id,
              "child_file_name": node[i].name,
              "child_file_id": node[i].id,
              "permission_id": id
            }
            this.dataService.updateUserFilePermission(FileData).subscribe(res => {
            })
          }
        }
        this.traverse(node[i], id);
      }
    }
  }
  ChangePermission(id: number) {
    if (this.data.type == "File") {
      let FileData = {
        "user_id": this.data.user_id,
        "parent_folder_id": this.data.parent_folder_id,
        "child_file_name": this.data.name,
        "child_file_id": this.data.id,
        "permission_id": id
      }
      this.dataService.updateUserFilePermission(FileData).subscribe(res => {
      })
    }
    else {
      let FolderData = {
        "user_id": this.data.user_id,
        "parent_folder_id": this.data.parent_folder_id,
        "child_folder_name": this.data.name,
        "child_folder_id": this.data.id,
        "permission_id": id
      }
      this.dataService.updateUserFolderPermission(FolderData).subscribe(res => {
      })

      if (this.data.children.length != 0) {
        this.traverse(this.data, id)
      }
    }
    this.dialogRef.close()
    if(id==1){
      this.openSnackBar("Permission Change to READ", " 🎉")
    }
    else{
      this.openSnackBar("Permission change to READ/WRITE", " 🎉")  
    }
    }
}
