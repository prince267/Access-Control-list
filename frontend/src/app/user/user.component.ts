import { Component, OnInit } from '@angular/core';
import { LoginService } from '../authservice/login.service'
import { NestedTreeControl } from '@angular/cdk/tree';
import { MatTreeNestedDataSource } from '@angular/material/tree';
import { Router } from '@angular/router'
import { MatDialog } from '@angular/material/dialog';
import { DataService } from '../data.service'
import { FilesFolderRelation } from '../FilesFoldersRelation/FileFolderRelation'
import { GroupUserDialogComponent } from '../DialogComponent/group-user-dialog/group-user-dialog.component'
import { FileDataDialogComponent } from '../DialogComponent/file-data-dialog/file-data-dialog.component'
import { FileFolderOptionDialogComponent } from '../DialogComponent/file-folder-option-dialog/file-folder-option-dialog.component'
import * as model from '../models/model'



@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})

export class UserComponent implements OnInit {
  treeControl = new NestedTreeControl<model.FileFolderNode>(node => node.children);
  dataSource = new MatTreeNestedDataSource<model.FileFolderNode>();
  constructor(
    private loginService: LoginService,
    private dataService: DataService,
    private router: Router,
    public dialog: MatDialog,
  ) {

  }
  hasChild = (_: number, node: model.FileFolderNode) => !!node.children && node.children.length > 0;

  data: model.userData
  groups: model.userGroups[]
  UserFilesFolders = []

  async  ngOnInit() {
    this.data = JSON.parse(this.loginService.token())
    this.GetUserGroups(this.data.user_id)
    this.GetFIleFolderTree()
  }

  async GetFIleFolderTree(){
    var userFolders = await this.dataService.GetUserFolders(this.data.user_id)
    var userFiles = await this.dataService.GetUserFiles(this.data.user_id)
    this.UserFilesFolders = FilesFolderRelation(userFiles, userFolders)
    this.dataSource.data = this.UserFilesFolders;

  }
  GetUserGroups(id: number) {
    this.dataService.GetUserGroups(id).subscribe((res: model.response) => {
      this.groups = res.data
    })
  }

  openGroupUserDialog(group_id: Number) {

    const dialogRef = this.dialog.open(GroupUserDialogComponent, {
      data: group_id
    });
  }
  FileFolderOptionDialog(user_id: number, parent_path_name: string, parent_id: number) {
    // console.log([user_id,parent_path_name,parent_id])
    const dialogRef=this.dialog.open(FileFolderOptionDialogComponent, {
      data: { user_id, parent_path_name, parent_id }
    })
    dialogRef.afterClosed().subscribe(result => {
      this.GetFIleFolderTree();
    });
  }
  openFileDataDialog(path: string, type: string, name: string, permission_id: number) {
    if (type == "Folder") {
      return
    }
    this.dialog.open(FileDataDialogComponent, {
      data: { path, type, name, permission_id }
    })
  }

}
