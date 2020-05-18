import { Component, OnInit } from '@angular/core'
import { DataService } from '../../data.service'
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { response, fileInFolder, folderInFolder } from '../../models/model'
import {MatSnackBar} from '@angular/material/snack-bar';
@Component({
  selector: 'app-delete-dialog',
  templateUrl: './delete-dialog.component.html',
  styleUrls: ['./delete-dialog.component.css']
})
export class DeleteDialogComponent implements OnInit {

  constructor(
    private _snackBar: MatSnackBar,
    private dataService: DataService,
    public dialogRef: MatDialogRef<DeleteDialogComponent>,) { }


    selectedValue: number
  allUsers = []
  ngOnInit(): void {
    this.dataService.GetAllUser().subscribe((res: response) => {
      this.allUsers = res.data
      console.log(this.allUsers)
    })
  }

   
  openSnackBar(message, action) {
    this._snackBar.open(message, action, {
      duration: 3000,
    });
  }

  DeleteUser(userId){
    console.log(userId)
    this.dataService.DelteUser(userId).subscribe(res=>{
      console.log(res)
      this.openSnackBar("File Update", " 🎉")  
    })
    this.dialogRef.close()
  }
  close(){
    this.dialogRef.close()
  }
}
