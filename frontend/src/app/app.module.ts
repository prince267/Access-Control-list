import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HttpClientModule } from '@angular/common/http';
import { UserComponent } from './user/user.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {MaterialModule} from './material-module';
import { LoginComponent } from './login/login.component';
import { GroupUserDialogComponent } from './DialogComponent/group-user-dialog/group-user-dialog.component';
import { FileDataDialogComponent } from './DialogComponent/file-data-dialog/file-data-dialog.component';
import { AdminDashboardComponent } from './admin-dashboard/admin-dashboard.component';
import { FileFolderOptionDialogComponent } from './DialogComponent/file-folder-option-dialog/file-folder-option-dialog.component';
import { ChangePermissionDialogComponent } from './DialogComponent/change-permission-dialog/change-permission-dialog.component';
import { RegisterDialogComponent } from './DialogComponent/register-dialog/register-dialog.component';
import { AssignUserDialogComponent } from './DialogComponent/assign-user-dialog/assign-user-dialog.component';
import { AdminFileFolderCreateDialogComponent } from './DialogComponent/admin-file-folder-create-dialog/admin-file-folder-create-dialog.component';

@NgModule({
  declarations: [
    AppComponent,
    UserComponent,
    LoginComponent,
    GroupUserDialogComponent,
    FileDataDialogComponent,
    AdminDashboardComponent,
    FileFolderOptionDialogComponent,
    ChangePermissionDialogComponent,
    RegisterDialogComponent,
    AssignUserDialogComponent,
    AdminFileFolderCreateDialogComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule,
    BrowserAnimationsModule,
    MaterialModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
