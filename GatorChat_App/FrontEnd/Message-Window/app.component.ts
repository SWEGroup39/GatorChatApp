import { HttpClient } from '@angular/common/http';
import { Component, ElementRef, ViewChild } from '@angular/core';
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

  @ViewChild('chatList', { static: true }) chatList!: ElementRef;


  title = 'chat-app';
   APIurl: string = `http://localhost:8080/api/messages`;

  chatMessages: {
    userId: number,
    recieverId: number,
    message: string,
    created_at: number,
    messageId: number,
  }[] =[]

  ngOnInit() {
    this.getMessages('5678', '1234').subscribe((result) => {
      this.chatMessages = result.chatMessages;
    });
  }

  chatInputMessage: string = "";
  searchInputMessage: string = "";
  currentUser = {
    name: 'John ',
    id: 5678,
    profileImageUrl:
    'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTKQmFYe2KZvQcnKEfGNICxM4I4udEh_-uG90chKLlXMx2HDGPr_ODubOdkpUFdJVGSKs0&usqp=CAU',

  }

  user1= {
    name: 'Jane ',
    id: 1234,
    profileImageUrl:
    'https://ps.w.org/user-avatar-reloaded/assets/icon-256x256.png?rev=2540745',

  }

  user2= {
    name: 'Jill ',
    id: 3,
    profileImageUrl:
    'https://e7.pngegg.com/pngimages/348/800/png-clipart-man-wearing-blue-shirt-illustration-computer-icons-avatar-user-login-avatar-blue-child.png',
    

  }


  public getMessages(Id1: string, Id2: string ): Observable<{ chatMessages: { userId: number, recieverId: number,messageId: number, message: string, created_at: number, updated_at: number, deleted_at: number}[] }> {
    const url = `${this.APIurl}/${Id1}/${Id2}`;
    return this.http.get<{ ID: number, CreatedAt: number, UpdatedAt: number, DeletedAt: number, message: string, SenderId: number, RecieverId: number, messageId: number }[]>(url)
      .pipe(
        map((response: any[]) => {
          console.log('Response:', response);
           const chatMessages = response.map(item => ({
            userId: item.sender_id,
            recieverId: item.RecieverId,
            messageId: item.ID,
            message: item.message,
            created_at: new Date(item.CreatedAt).getTime(),
            updated_at: new Date(item.UpdatedAt).getTime(),
            deleted_at: new Date(item.DeletedAt).getTime(),
          }));
          return { chatMessages };
        }),
        tap((response: any) => {
          console.log('body: ' + response.body);
        })
      );
  }
  

  searchMessages(content: string): Observable<{ Messages: { messageId: number }[] }> {
    const url = `${this.APIurl}/${content}`;
    return this.http.get<{ ID: number }[]>(url)
      .pipe(
        map((response: { ID: number }[]) => {
          console.log('Messages found:', response);
          const Messages = response.map(item => ({ messageId: item.ID }));
          return { Messages };
        }),
        tap((response: any) => {
          console.log('body:', response.body);
        })
      );
  }

  
  sendMessage(Message: any): Observable<any> {
    const url = `${this.APIurl}`;
    console.log('Send Message');
    console.log(Message);

    return this.http.post(url, {
      sender_id: Message.userId.toString(),
      receiver_id: Message.recieverId.toString(),
      message: Message.message,
    }).pipe(
      tap(response => {
        console.log('Message sent:', response);
        this.chatMessages.push(Message);
      }),
      catchError(error => {
        console.error(error);
        return throwError(error);
      })
    ); 
  }


  send() {
    const newMessage = {
      message: this.chatInputMessage,
      userId: this.currentUser.id,
      recieverId: this.user1.id,
      created_at: Date.now(),
    };
  
    this.sendMessage(newMessage).subscribe(
      (result) => {
          console.log(result.message);
      },
      (error) => {
        // If there was an error with the request, log the error
        console.error(error);
      }
    );
  
    this.chatInputMessage = "";
  }


search(message: string) {
  console.log('search: ' + message)

  let  Ids: {messageId: number}[] = [];
  this.searchMessages(message).subscribe(result => {
  Ids = result.Messages;

  Ids.forEach(id => console.log('Id found:' + id.messageId))

  const messageIndex = this.chatMessages.findIndex(ms => ms.messageId == Ids[0].messageId);
  console.log(messageIndex)
   
    console.log(this.chatList.nativeElement.children[messageIndex]);

      if (messageIndex >= 0) {
        const messageElement = this.chatList.nativeElement.children[messageIndex];
        messageElement.scrollIntoView({ behavior: 'smooth' });
      }
  
      this.searchInputMessage = "";

    });
  }
}
