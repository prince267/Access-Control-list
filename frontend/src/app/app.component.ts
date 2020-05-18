import { Component } from '@angular/core';
import { LoginService } from './authservice/login.service'
import { Router } from '@angular/router'
import { MatDialog } from '@angular/material/dialog';
import { RegisterDialogComponent } from '../app/DialogComponent/register-dialog/register-dialog.component'
import { DeleteDialogComponent } from '../app/DialogComponent/delete-dialog/delete-dialog.component'
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'ACL';
  constructor(public loginService: LoginService,
    private router: Router,
    public dialog: MatDialog,
  ) { }

  logout() {
    this.loginService.logout()
    this.router.navigate(["/"])
  }

  openRegisterDialog() {
    const dialogRef=this.dialog.open(RegisterDialogComponent)
    dialogRef.afterClosed().subscribe(res=>{
      window.location.reload();
    })
  }

  openDeleteDialog(){
  const dialogRef=this.dialog.open(DeleteDialogComponent) 
  dialogRef.afterClosed().subscribe(res=>{
    window.location.reload();
   })
}
}