import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ChatListComponentComponent } from './chat-list-component.component';

describe('ChatListComponentComponent', () => {
  let component: ChatListComponentComponent;
  let fixture: ComponentFixture<ChatListComponentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ChatListComponentComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ChatListComponentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
