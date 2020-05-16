import { Component, OnInit } from '@angular/core';
import { DataService } from '../data.service';
import { response, FileFolderNode, AllFileFolderNode } from '../models/model'
import { NestedTreeControl } from '@angular/cdk/tree';
import { MatTreeNestedDataSource } from '@angular/material/tree';
import { MatDialog } from '@angular/material/dialog';
import { FilesFolderRelation } from '../FilesFoldersRelation/FileFolderRelation'
import { MatSnackBar } from '@angular/material/snack-bar';
import { AllFilesFolderRelation } from '../FilesFoldersRelation/AllFilesFoldersRelation'
import { ChangePermissionDialogComponent } from '../DialogComponent/change-permission-dialog/change-permission-dialog.component'
import { AssignUserDialogComponent } from '../DialogComponent/assign-user-dialog/assign-user-dialog.component'
import {AdminFileFolderCreateDialogComponent} from '../DialogComponent/admin-file-folder-create-dialog/admin-file-folder-create-dialog.component'
interface user {
  user_id: number,
  first_name: string,
  last_name: string,
  password: string
}

@Component({
  selector: 'app-admin-dashboard',
  templateUrl: './admin-dashboard.component.html',
  styleUrls: ['./admin-dashboard.component.css']
})

export class AdminDashboardComponent implements OnInit {
  treeControl = new NestedTreeControl<FileFolderNode>(node => node.children);
  dataSource = new MatTreeNestedDataSource<FileFolderNode>();

  FileFolderTreeControl = new NestedTreeControl<AllFileFolderNode>(node => node.children)
  FileFolderDataSource = new MatTreeNestedDataSource<AllFileFolderNode>();

  constructor(
    private _snackBar: MatSnackBar,
    public dialog: MatDialog,
    private dataService: DataService
  ) { }

  Id: number
  hasChild = (_: number, node: FileFolderNode) => !!node.children && node.children.length > 0;
  FileFolderhasChild = (_: number, node: AllFileFolderNode) => !!node.children && node.children.length > 0;

  UserFilesFolders = []
  FoldersFiles = []
  users: user[]
  selectedValue: number

  ngOnInit() {
    this.dataService.GetAllUser().subscribe((res: response) => {
      this.users = res.data
    })
    this.GetAllFileAndFolders()
  }

  openSnackBar(message, action) {
    this._snackBar.open(message, action, {
      duration: 3000,
    });
  }

  async GetUserFileFolderTree(UserId: number) {
    this.Id = UserId
    if (UserId == undefined) {
      this.dataSource.data = []
      return
    }
    var userFolders = await this.dataService.GetUserFolders(UserId)
    var userFiles = await this.dataService.GetUserFiles(UserId)
    this.UserFilesFolders = FilesFolderRelation(userFiles, userFolders)
    console.log(this.UserFilesFolders)
    this.dataSource.data = this.UserFilesFolders;
  }

  PermissionDialog(node1) {
    const dialogRef = this.dialog.open(ChangePermissionDialogComponent, {
      data: node1
    })
    dialogRef.afterClosed().subscribe(result => {
      this.GetUserFileFolderTree(this.Id)
    });
  }

  async GetAllFileAndFolders() {
    var Folders = await this.dataService.GetAllFolders()
    var Files = await this.dataService.GetAllFiles()
    this.FoldersFiles = FilesFolderRelation(Folders, Files)
    this.FileFolderDataSource.data = this.FoldersFiles
  }


  DeleteInFolder(node) {
    for (var i in node) {
      if (!!node[i] && typeof (node[i]) == "object") {
        if (i !== "children") {
          if (node[i].type == "Folder") {
            this.dataService.DeleteFolderInFolderById(node[i].id).subscribe(res => {
              console.log(res)
            })
          }
          if (node[i].type == "File") {
            this.dataService.DeleteFileInFolderById(node[i].id).subscribe(res => {
              console.log(res)
            })
          }
        }
        this.DeleteInFolder(node[i]);
      }
    }
  }


  DeleteFilesFolder(node) {
    for (var i in node) {
      if (!!node[i] && typeof (node[i]) == "object") {
        if (i !== "children") {
          if (node[i].type == "Folder") {
            this.dataService.DeleteFolderById(node[i].id).subscribe(res => {
              console.log(res)
            })
          }
          if (node[i].type == "File") {
            this.dataService.DeleteFileById(node[i].id).subscribe(res => {
              console.log(res)
            })
          }
        }
        this.DeleteFilesFolder(node[i]);
      }
    }
  }

  Delete(node) {

    if (node.type == "File") {
      this.dataService.DeleteFileInFolderById(node.id).subscribe(res => {
        console.log(res)
      })
      this.dataService.DeleteFileById(node.id).subscribe(res => {
        console.log(res)
      })
      this.dataService.DeleteEntity(node.path_name).subscribe(res => {
        console.log(res)
      })
    }
    else {
      // Delete FolderInFolder and FileInFolder data
      this.dataService.DeleteFolderInFolderById(node.id).subscribe(res => {
        console.log(res)
      })
      if (node.children.length != 0) {
        this.DeleteInFolder(node)
      }
      // Delete Folders and Files Data
      this.dataService.DeleteFolderById(node.id).subscribe(res => {
        console.log(res)
      })
      if (node.children.length != 0) {
        this.DeleteFilesFolder(node)
      }

      this.dataService.DeleteEntity(node.path_name).subscribe(res => {
        console.log(res)
      })

    }
    this.GetAllFileAndFolders();
    this.GetUserFileFolderTree(this.Id)


    this.openSnackBar(node.type == "Folder" ?
      "Folder Deleted" : "File Deleted", " 🎉")
  }

  assignDialog(node) {
    if (node.id == 0) {
      return
    }
    let data = {
      entityData: node,
      id: node.id,
      type: node.type
    }
    const dialogRef = this.dialog.open(AssignUserDialogComponent, {
      data: data
    })
    dialogRef.afterClosed().subscribe(res => {
      this.GetUserFileFolderTree(this.Id)
      this.GetAllFileAndFolders()
    })
  }
  NewEntity(node){
    let folderData={
      id:node.id,
      path_name:node.path_name
    }
  const dialogRef=  this.dialog.open(AdminFileFolderCreateDialogComponent,{
      data:folderData
    })
    dialogRef.afterClosed().subscribe(res=>{
      this.GetUserFileFolderTree(this.Id)
      this.GetAllFileAndFolders()
      
    })
  }
}

