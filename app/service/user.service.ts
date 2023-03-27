import { Injectable } from '@angular/core';
import { HttpClient, HttpParams, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { User } from '../interface/user';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  isLoggedIn: boolean = false;

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
      username: username,
      password: password
    };
    return this.http.post<any>('http://localhost:8080/api/users/User',requestBody);
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