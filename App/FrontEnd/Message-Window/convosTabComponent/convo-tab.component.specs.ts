import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ConvoTabComponent } from './convo-tab.component';

describe('ConvoTabComponent', () => {
  let component: ConvoTabComponent;
  let fixture: ComponentFixture<ConvoTabComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ConvoTabComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ConvoTabComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
