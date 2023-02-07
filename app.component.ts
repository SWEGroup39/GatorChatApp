import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Gator Chat';
  message= '';
  messages: string[]=[
    
  ]
  addMessage(newMessage: string){
    this.messages.push(newMessage);
  }
}
