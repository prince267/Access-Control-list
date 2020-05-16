import { Component, OnInit, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { DataService } from '../../data.service'
import {fileDataResponse} from '../../models/model'
import {MatSnackBar} from '@angular/material/snack-bar';
interface fileInfo {
  path: string,
  type: string,
  name:string,
  permission_id:number
}

@Component({
  selector: 'app-file-data-dialog',
  templateUrl: './file-data-dialog.component.html',
  styleUrls: ['./file-data-dialog.component.css']
})
export class FileDataDialogComponent implements OnInit {

  constructor(
    private _snackBar: MatSnackBar,
    private dataService: DataService,
    public dialogRef: MatDialogRef<FileDataDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: fileInfo) { }

  FileData: string
  ngOnInit(): void {
    this.GetFileData(this.data.path)
  }
  
  openSnackBar(message, action) {
    this._snackBar.open(message, action, {
      duration: 3000,
    });
  }
  
  GetFileData(path:string){
    let FileInfo={
      "path_name":path,
      "content":""
    }
    this.dataService.GetFileData(FileInfo).subscribe((res: fileDataResponse) => {
      this.FileData = res.data
    })
  }

  closeDialog(){
    this.dialogRef.close()
  }

  UpdateFileData(fileContent:string) {
    let FileInfo={
      "path_name":this.data.path,
      "content":fileContent
    }
    this.dataService.WriteIntoFile(FileInfo).subscribe((res:fileDataResponse)=>{
      this.openSnackBar("File Update", " 🎉")  
      console.log(res)
    })
    this.dialogRef.close()
    
  }
}
