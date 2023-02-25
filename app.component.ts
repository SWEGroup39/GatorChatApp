import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { EMPTY, Observable, throwError } from 'rxjs';
import { map } from 'rxjs/operators';
import { tap, catchError } from 'rxjs/operators';



@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {

  constructor(private http: HttpClient) { }

  chatInputMessage: string = "";
  currentUser = {
    name: 'John ',
    id: 1234,
    profileImageUrl:
    'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTKQmFYe2KZvQcnKEfGNICxM4I4udEh_-uG90chKLlXMx2HDGPr_ODubOdkpUFdJVGSKs0&usqp=CAU',

  }

  user1= {
    name: 'Jane ',
    id: 4321,
    profileImageUrl:
    'https://ps.w.org/user-avatar-reloaded/assets/icon-256x256.png?rev=2540745',

  }

  user2= {
    name: 'Jill ',
    id: 3,
    profileImageUrl:
    'https://e7.pngegg.com/pngimages/348/800/png-clipart-man-wearing-blue-shirt-illustration-computer-icons-avatar-user-login-avatar-blue-child.png',
    

  }

  
  // getMessages() {
  //   this.http.get('http://localhost:8080/api/messages').subscribe(response => {
  //     console.log(response);
  //   });
  // }

  public getMessages(): Observable<{ chatMessages: { userId: number, recieverId: number,messageId: number, message: string, created_at: number, updated_at: number, deleted_at: number}[] }> {
    return this.http.get<{ ID: number, CreatedAt: number, UpdatedAt: number, DeletedAt: number, message: string, SenderId: number, RecieverId: number, messageId: number }[]>('http://localhost:8080/api/messages')
      .pipe(
        map((response: any[]) => {
          console.log('Response:', response);
           const chatMessages = response.map(item => ({
            userId: item.sender_id,
            recieverId: item.RecieverId,
            messageId: item.messageId,
            message: item.message,
            created_at: new Date(item.CreatedAt).getTime(),
            updated_at: new Date(item.UpdatedAt).getTime(),
            deleted_at: new Date(item.DeletedAt).getTime(),
          }));
          return { chatMessages };
        })
      );
  }

  // public sendMessage(message: { userId: number, recieverId: number, message: string, created_at: number, updated_at: number, deleted_at: number}): Observable<any> {
  //   return this.http.post<any>('http://localhost:8080/api/messages', message)
  //     .pipe(
  //       tap(response => {
  //         console.log('Message sent:', response);
  //       }),
  //     );
  // }
  // public sendMessage(newMessage: any): Observable<any> {
  //   console.log('send Messagesa is called');
  //   const url = 'http://localhost:8080/api/messages';
  //   return this.http.post(url, {
  //     Sender_ID: newMessage.userId.toString(),
  //     Receiver_ID: newMessage.recieverId.toString(),
  //     Message: newMessage.message,
  //     CreatedAt: newMessage.created_at,
  //     UpdatedAt: newMessage.updated_at,
  //     DeletedAt: newMessage.deleted_at,
  //   })
  // }

  sendMessage(newMessage: any): Observable<any> {
    const url = 'http://localhost:8080/api/messages';
    console.log('Send Message');
    console.log(newMessage);
    return this.http.post(url, {
      sender_id: newMessage.userId.toString(),
      receiver_id: newMessage.recieverId.toString(),
      message: newMessage.message,
    })
  }

  ngOnInit() {
    this.getMessages().subscribe((result) => {
      this.chatMessages = result.chatMessages;
    });
  }
  

  chatMessages: {
    userId: number,
    recieverId: number,
    message: string,
    created_at: number,
  }[] =[]

  title = 'chat-app';

  send() {
    const newMessage = {
      message: this.chatInputMessage,
      userId: this.currentUser.id,
      recieverId: this.user1.id,
      created_at: Date.now(),
    };


    console.log('Inside send function');
    
    this.sendMessage(newMessage).subscribe(result => this.chatMessages.push(result));
    //this.chatMessages.push(newMessage);
    this.chatInputMessage = "";
  }
}
