import { HttpClient } from '@angular/common/http';
import { Component, ElementRef, ViewChild, OnInit, ChangeDetectorRef } from '@angular/core';
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
  title = 'chat-app';

  // variables 
  showMenu: boolean = false;
  editMode: boolean = false;
  currentMessage: any;
  editedMessage: string = "";

  chatInputMessage: string = "";
  searchInputMessage: string = "";

  // Messages List
  chatMessages: {
    userId: number,
    recieverId: number,
    message: string,
    created_at: number,
    messageId: number,
  }[] =[]

  APIurl: string = `http://localhost:8080/api/messages`;

  currentUser = {
    name: 'null',
    id: 'null',
  }

  user1= {
    name: 'null',
    id: 'null',
  }

  constructor(private http: HttpClient, private route: ActivatedRoute, private location: Location, private cd: ChangeDetectorRef) { }
  
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

  
  goBack() {
    this.location.back();
  }

  goBackTwice() {
    this.location.back();
    this.location.back();
  }

  showDate(item: any, index: number): boolean {
    if (index === 0) {
      // always show the date for the first item
      return true;
    }
    const prevItem = this.chatMessages[index - 1];
    const prevDate = new Date(prevItem.created_at);
    const currDate = new Date(item.created_at);
  
    // show the date if it's different from the previous message's date
    return prevDate.toDateString() !== currDate.toDateString();
  }


isSameDate(time1: number, time2: number) {
    return this.convertTimestampToDay(time1).getTime() === this.convertTimestampToDay(time2).getTime();
}

 clickMenu(event: MouseEvent)
  {
    if (event.detail === 1 && event.button === 0) {
      console.log('clicked');
    }
  }

  closeToggle() {
    this.currentMessage = null;
  }


  showToggleMenu(event: MouseEvent, item: any) {
    event.preventDefault();
    this.currentMessage = item;
    
  }

  convertTimestampToDay(timestamp: number): Date {
    const date = new Date(timestamp);
    date.setHours(0, 0, 0, 0);
    return date;
  }

  editMessage() {
    // event.stopPropagation(); 
    // console.log("Editing message:", this.currentMessage); // prevent event from bubbling up
    // this.currentMessage = message;
    // this.editMode = true;
    // this.editedMessage = message.message;
  }



  delete(id: number) {
    console.log("Deleting message:", this.currentMessage);
    this.deleteMessage(id).subscribe(
      (result) => {
          this.getMessages( this.currentUser.id , this.user1.id  ).subscribe((result) => {
            this.chatMessages = result.chatMessages;
        
              });
        }
      
    )

    console.log(this.chatMessages);
    this.cd.detectChanges();

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


  // API functions (Request)

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
            day: new Date(item.CreatedAt)
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
//   sendMessage(Ids: any): Observable<any> {
//     const url = 'http://localhost:8080/api/messages';
//     console.log('Send Message');
//     console.log(Ids);

//     return this.http.post(url, {
//       sender_id: Ids.userId.toString(),
//       receiver_id: Ids.recieverId.toString(),
//       message: Ids.message,
//     }).pipe(
//       map(response => {
//         console.log('Message sent:', response);
//         const newMessage = {
//           userId: Ids.userId,
//           recieverId: Ids.recieverId,
//           message: Ids.message,
//           created_at: response.,
//           messageId: response.ID
//         };
//         this.chatMessages.push(newMessage);
//         return newMessage;
//       }),
//       catchError(error => {
//         console.error(error);
//         return throwError(error);
//       })
//     ); 
// }


  deleteMessage(id: number): Observable<any> {
    const url = `http://localhost:8080/api/messages/${id}`;
    console.log('Deleting Message: ' + id);
  
    return this.http.delete(url).pipe(
      map((response: any) => {
        console.log('Deleted sucessfully')
        return response;
      }),
      catchError((error: any) => {
        console.log('Delete failed')
        return throwError(error);
      })
    );
  }


}
}
