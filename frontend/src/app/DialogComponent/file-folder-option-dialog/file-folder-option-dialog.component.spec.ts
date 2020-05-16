import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FileFolderOptionDialogComponent } from './file-folder-option-dialog.component';

describe('FileFolderOptionDialogComponent', () => {
  let component: FileFolderOptionDialogComponent;
  let fixture: ComponentFixture<FileFolderOptionDialogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FileFolderOptionDialogComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FileFolderOptionDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
