import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AdminFileFolderCreateDialogComponent } from './admin-file-folder-create-dialog.component';

describe('AdminFileFolderCreateDialogComponent', () => {
  let component: AdminFileFolderCreateDialogComponent;
  let fixture: ComponentFixture<AdminFileFolderCreateDialogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AdminFileFolderCreateDialogComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AdminFileFolderCreateDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
