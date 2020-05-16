import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GroupUserDialogComponent } from './group-user-dialog.component';

describe('GroupUserDialogComponent', () => {
  let component: GroupUserDialogComponent;
  let fixture: ComponentFixture<GroupUserDialogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GroupUserDialogComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GroupUserDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
