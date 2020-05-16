import { Component, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { FormBuilder, FormGroup, Validators, FormControl } from '@angular/forms';
import { Router, NavigationEnd } from "@angular/router";
import { DataService } from "../../data.service"
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-register-dialog',
  templateUrl: './register-dialog.component.html',
  styleUrls: ['./register-dialog.component.css']
})
export class RegisterDialogComponent implements OnInit {

  constructor(
    private _snackBar: MatSnackBar,
    private dataService: DataService,
    private formBuilder: FormBuilder,
    private router: Router,
    public dialogRef: MatDialogRef<RegisterDialogComponent>
  ) { }

  registrationForm: FormGroup;
  first_name = new FormControl('', [Validators.required, Validators.minLength(1), Validators.maxLength(100)]);
  last_name = new FormControl('', [Validators.required, Validators.minLength(1), Validators.maxLength(100)]);
  Password = new FormControl('', [Validators.required, Validators.minLength(4), Validators.maxLength(10)]);

  ngOnInit(): void {
    this.createFormValidations()
  }

  openSnackBar(message, action) {
    this._snackBar.open(message, action, {
      duration: 3000,
    });
  }

  createFormValidations() {
    this.registrationForm = this.formBuilder.group({
      first_name: this.first_name,
      last_name: this.last_name,
      Password: this.Password
    })
  }
  CreateUser() {
    let userData = {
      "first_name": this.registrationForm.value.first_name,
      "last_name": this.registrationForm.value.last_name,
      "password": this.registrationForm.value.Password
    }
    console.log(userData)
    this.dataService.NewUser(userData).subscribe(res => {
      this.openSnackBar("New User Created", " 🎉")
          
      this.dialogRef.close()
      this.router.navigate['/admin']
    })
    // console.log(this.registrationForm.value)
  }
}
