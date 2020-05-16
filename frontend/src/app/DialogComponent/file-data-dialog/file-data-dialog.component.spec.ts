import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FileDataDialogComponent } from './file-data-dialog.component';

describe('FileDataDialogComponent', () => {
  let component: FileDataDialogComponent;
  let fixture: ComponentFixture<FileDataDialogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FileDataDialogComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FileDataDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
