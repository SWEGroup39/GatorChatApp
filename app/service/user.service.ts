import { Injectable } from '@angular/core';
import { HttpClient, HttpParams, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { User } from '../interface/user';
@Injectable({
  providedIn: 'root'
})
export class UserService {
  isLoggedIn: boolean = false;
  emailAddress: string = ``;

  constructor(private http: HttpClient) {
    
  }

  getUsers(): Observable<User[]>{
    return this.http.get<User[]>(`http://localhost:8080/api/users`);
  }

  createUser(user: User): Observable<User>{
    return this.http.post<User>(`http://localhost:8080/api/users`, user);
  }
  getUser(password: string, username: string ): Observable<any>{
    const requestBody: any = {
      email: username,
      password: password
    };
    
    return this.http.post<any>('http://localhost:8080/api/users/User',requestBody);
  }

  getNextID():Observable<string>{
    return this.http.get<any>('http://localhost:8080/api/users/nextID');
  }

  deleteUser(userID:string):Observable<string>{
    return this.http.delete<string>(`http://localhost:8080/api/users/${userID}`)
  }

  updateFN(userID: string, newFN:string):Observable<string>{
    const requestBody:any = {
      full_name: newFN
    }
    console.log(requestBody)
    return this.http.put<string>(`http://localhost:8080/api/users/updateFN/${userID}`,requestBody)
  }

  updatePhoneNum(userID:string, newPN:string):Observable<string>{
    const requestBody:any = {
      phone_number:newPN
    }
    return this.http.put<string>(`http://localhost:8080/api/users/updatePN/${userID}`,requestBody)
  }
  updatePassword(userID:string, newPass:string,oldPass:string):Observable<string>{
    const requestBody:any = {
      password:newPass,
      original_pass: oldPass
    }
    console.log(oldPass)
    console.log(newPass)
    return this.http.put<string>(`http://localhost:8080/api/users/updateP/${userID}`,requestBody)
  }
  
  updateUsername(userID:string, newUsername:string):Observable<string>{
    const requestBody:any = {
      username: newUsername
    }
    return this.http.put<string>(`http://localhost:8080/api/users/updateN/${userID}`,requestBody)
  }

  searchContact(searchVal:string):Observable<any>{
    const requestBody:any = {
      username:searchVal
    }
    return this.http.post<any>(`http://localhost:8080/api/users/search`,requestBody)
  }

  addConversationID(sender_id:string, receiver_id:string):Observable<string>{
    
    return this.http.get<string>(`http://localhost:8080/api/users/${sender_id}/${receiver_id}`)
  }
}





// //   searchMessages(content: string): Observable<{ Messages: { messageId: number }[] }> {
//   const url = `${this.APIurl}/${content}`;
//   return this.http.get<{ ID: number }[]>(url)
//     .pipe(
//       map((response: { ID: number }[]) => {
//         console.log('Messages found:', response);
//         const Messages = response.map(item => ({ messageId: item.ID }));
//         return { Messages };
//       }),
//       tap((response: any) => {
//         console.log('body:', response.body);
//       })
//     );
// }