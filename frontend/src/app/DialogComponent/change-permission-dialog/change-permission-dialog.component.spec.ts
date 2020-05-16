import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ChangePermissionDialogComponent } from './change-permission-dialog.component';

describe('ChangePermissionDialogComponent', () => {
  let component: ChangePermissionDialogComponent;
  let fixture: ComponentFixture<ChangePermissionDialogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ChangePermissionDialogComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ChangePermissionDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
