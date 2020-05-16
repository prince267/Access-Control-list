import { Component, OnInit, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { DataService } from '../../data.service'
import { groupUsers, response } from '../../models/model'

@Component({
  selector: 'app-group-user-dialog',
  templateUrl: './group-user-dialog.component.html',
  styleUrls: ['./group-user-dialog.component.css']
})
export class GroupUserDialogComponent implements OnInit {

  constructor(
    private dataService: DataService,
    public dialogRef: MatDialogRef<GroupUserDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: number
  ) { }
  
  users: groupUsers[]

  ngOnInit() {
    this.GetGroupUsers(this.data)
  }

  GetGroupUsers(id: number) {
    this.dataService.GetGroupUsers(id).subscribe((res: response) => {
      this.users = res.data
    })
  }

  close(){
    this.dialogRef.close()
  }
}
