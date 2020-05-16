import { Component } from '@angular/core';
import {LoginService} from './authservice/login.service'
import {Router} from '@angular/router'
import { MatDialog } from '@angular/material/dialog';
import {RegisterDialogComponent} from '../app/DialogComponent/register-dialog/register-dialog.component'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'ACL';
  constructor(public loginService:LoginService,
    private router:Router,
    public dialog: MatDialog,
    ){}

  logout(){
      this.loginService.logout()
    this.router.navigate(["/"])
    }
    
    openRegisterDialog(){
      this.dialog.open(RegisterDialogComponent)
    }
  
}
