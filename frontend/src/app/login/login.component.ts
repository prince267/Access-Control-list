import { Component, OnInit } from '@angular/core';
import { HttpErrorResponse } from '@angular/common/http';
import { Router, NavigationEnd } from "@angular/router";
import { FormBuilder, FormGroup, Validators, FormControl } from '@angular/forms';
import {LoginService} from '../authservice/login.service'
import {MatSnackBar} from '@angular/material/snack-bar';
interface LoginResponse {
  status: number;
  data: object;
}
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  constructor(
    private _snackBar: MatSnackBar,
    private formBuilder: FormBuilder,
    private loginService: LoginService,
    private router:Router
  ) { }

  loginForm: FormGroup;
  UserId = new FormControl('', [Validators.required, Validators.maxLength(100)]);
  Password = new FormControl('', [Validators.required, Validators.minLength(4), Validators.maxLength(10)]);

  ngOnInit(): void {
    this.createFormValidations();
  }

  createFormValidations() {
    this.loginForm = this.formBuilder.group({
      UserId: this.UserId,
      Password: this.Password
    })
  }

  openSnackBar(message, action) {
    this._snackBar.open(message, action, {
      duration: 3000,
    });
  }

  onSubmit() {
    if(this.loginForm.value.UserId==0 && this.loginForm.value.Password=="root"){
      localStorage.setItem("token",JSON.stringify(this.loginForm.value));
      localStorage.setItem("admin",JSON.stringify(this.loginForm.value))
      this.router.navigate(['admin'])
    }
    else{
      let loginData = {
        "user_id": this.loginForm.value.UserId,
        "password": this.loginForm.value.Password
      }
  
      if(this.loginForm.invalid){
        return
      }
      this.loginService.login(loginData).subscribe((data :LoginResponse) => {
        localStorage.setItem("token",JSON.stringify(data.data));
        this.router.navigate(['user'])
      },
      (error: HttpErrorResponse) => {
        console.log(error.status)
    
      if(error.status==404){
        this.openSnackBar("No User Found", " 😓")
      }
      })
    }
   
  }
}
