import { Component, OnInit, Inject, PLATFORM_ID } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { DataService } from '../../data.service'
import { response, fileInFolder, folderInFolder } from '../../models/model'
interface user {
  user_id: number,
  first_name: string,
  last_name: string,
  password: string
}

interface node {
  entityData: any,
  id: number,
  type: string,
}
@Component({
  selector: 'app-assign-user-dialog',
  templateUrl: './assign-user-dialog.component.html',
  styleUrls: ['./assign-user-dialog.component.css']
})
export class AssignUserDialogComponent implements OnInit {
  constructor(
    private dataService: DataService,
    public dialogRef: MatDialogRef<AssignUserDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: node
  ) { }

  Entity: number;
  Options: number[] = [1, 2];

  users = []
  allUsers = []
  selectedValue: number
  ngOnInit(): void {
    console.log(this.data.entityData)
    this.dataService.GetAllUser().subscribe((res: response) => {
      this.allUsers = res.data
      console.log(this.allUsers)
    })


    if (this.data.type == "Folder") {
      this.dataService.GetFolderUser(this.data.id).subscribe((res: {
        status: number,
        data: folderInFolder[]
      }
      ) => {
        this.users = res.data
        console.log("this. users",this.users)
      })
    }
    if (this.data.type == "File") {
      this.dataService.GetFileUser(this.data.id).subscribe((res: {
        status: number,
        data: fileInFolder[]
      }) => {
        this.users = res.data

      })
    }

    // this.unassignedUsers(this.allUsers,this.users)
  }

  traverse(node, userId, pID) {
    for (var i in node) {
      if (!!node[i] && typeof (node[i]) == "object") {
        if (i !== "children") {

          if (node[i].type == "Folder") {
            let temp1 = node[i]
            this.dataService.CheckIsFolderUser(userId, node[i].id).subscribe((res: {
              status: number,
              data: any,
              message?: string
            }) => {
              if (res.data == "") {
                let UserFolderData = {
                  "user_id": userId,
                  "parent_folder_id": temp1.parent_folder_id,
                  "child_folder_name": temp1.name,
                  "child_folder_id": temp1.id,
                  "permission_id": pID
                }
                // console.log("userFolderData **",UserFolderData)
                this.dataService.NewUserFolder(UserFolderData).subscribe(res => {
                  console.log(res)
                })
              }
            })

          }

          if (node[i].type == "File") {
            let temp = node[i]
            this.dataService.CheckIsFileUser(userId, node[i].id).subscribe((res: {
              status: number,
              data: any,
              message?: string
            }) => {
              if (res.data == "") {
                let UserFileData = {
                  "user_id": userId,
                  "parent_folder_id": temp.parent_folder_id,
                  "child_file_name": temp.name,
                  "child_file_id": temp.id,
                  "permission_id": pID
                }
                this.dataService.NewUserFile(UserFileData).subscribe(res => {
                  console.log(res)
                })
              }
            })
          }
        }
        this.traverse(node[i], userId, pID);
      }
    }
  }

  assigne(userId, pID) {
    console.log(userId, pID)
    if (this.data.entityData.type == "File") {
      this.dataService.CheckIsFileUser(userId, this.data.entityData.id).subscribe((res: {
        status: number,
        data: any,
        message?: string
      }) => {
        console.log(res)
        if (res.data == "") {
          let UserFileData = {
            "user_id": userId,
            "parent_folder_id": this.data.entityData.parent_folder_id,
            "child_file_name": this.data.entityData.name,
            "child_file_id": this.data.entityData.id,
            "permission_id": pID
          }
          this.dataService.NewUserFile(UserFileData).subscribe(res => {
            console.log(res)
          })
        }
      })
    }
    else {
      this.dataService.CheckIsFolderUser(userId, this.data.entityData.id).subscribe((res: {
        status: number,
        data: any,
        message?: string
      }) => {
        console.log(res)
        if (res.data == "") {
          let UserFolderData = {
            "user_id": userId,
            "parent_folder_id": this.data.entityData.parent_folder_id,
            "child_folder_name": this.data.entityData.name,
            "child_folder_id": this.data.entityData.id,
            "permission_id": pID
          }
          this.dataService.NewUserFolder(UserFolderData).subscribe(res => {
            console.log(res)
          })
        }
      })
      if (this.data.entityData.children.length > 0) {
        this.traverse(this.data.entityData, userId, pID)
      }
    }
this.dialogRef.close()
  }

  close() {
    this.dialogRef.close()
  }
}
