import { HttpClient } from '@angular/common/http';
import { Component, ElementRef, ViewChild, OnInit } from '@angular/core';
import { EMPTY, Observable, throwError } from 'rxjs';
import { map } from 'rxjs/operators';
import { tap, catchError } from 'rxjs/operators';
import { ActivatedRoute } from '@angular/router';
import { Location } from '@angular/common';


@Component({
  selector: 'app-root',
  templateUrl: './messages.component.html',
  styleUrls: ['./messages.component.scss']
})
export class MessagesComponent implements OnInit{
  // currId: string ='';
  // otherId: string = ''

  chatInputMessage: string = "";
  searchInputMessage: string = "";

  chatMessages: {
    userId: number,
    recieverId: number,
    message: string,
    created_at: number,
    messageId: number,
  }[] =[]

  APIurl: string = `http://localhost:8080/api/messages`;

  constructor(private http: HttpClient, private route: ActivatedRoute, private location: Location) { }
  
  ngOnInit() {

    this.route.queryParams.subscribe(params => {
    this.currentUser.id  = params['id1'] ?? '0000';
    this.user1.id = params['id2'] ?? '0000';

    this.getMessages( this.currentUser.id , this.user1.id  ).subscribe((result) => {
    this.chatMessages = result.chatMessages;


    });

    });
  }

  @ViewChild('chatList', { static: true }) chatList!: ElementRef;

  currentUser = {
    name: 'moe ',
    id: '0001',
    profileImageUrl:
    'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTKQmFYe2KZvQcnKEfGNICxM4I4udEh_-uG90chKLlXMx2HDGPr_ODubOdkpUFdJVGSKs0&usqp=CAU',

  }

  user1= {
    name: 'Jane ',
    id: '0002',
    profileImageUrl:
    'https://ps.w.org/user-avatar-reloaded/assets/icon-256x256.png?rev=2540745',

  }

  goBack() {
    this.location.back();
  }

  goBackTwice() {
    this.location.back();
    this.location.back();
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
    const url = `http://localhost:8080/api/messages/${content}`;
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

  

  sendMessage(Ids: any): Observable<any> {
    const url = 'http://localhost:8080/api/messages';
    console.log('Send Message');
    console.log(Ids);

    return this.http.post(url, {
      sender_id: Ids.userId.toString(),
      receiver_id: Ids.recieverId.toString(),
      message: Ids.message,
    }).pipe(
      tap(response => {
        console.log('Message sent:', response);
        this.chatMessages.push(Ids);
      }),
      catchError(error => {
        console.error(error);
        return throwError(error);
      })
    ); 
  }


  title = 'chat-app';

  send() {
    const newMessage = {
      message: this.chatInputMessage,
      userId: this.currentUser.id,
      recieverId: this.user1.id,
      created_at: Date.now(),
    };
  
    this.sendMessage(newMessage).subscribe(
      (result) => {
        // Check if post was successful
        if (result.success) {
          // If post was successful, push newMessage to the chatMessages list
          this.chatMessages.push(result);
        } else {
          // If post was not successful, log the error
          console.error(result.message);
        }
      },
      (error) => {
        // If there was an error with the request, log the error
        console.error(error);
      }
    );
  
    this.chatInputMessage = "";
  }


search(message: string) {

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
