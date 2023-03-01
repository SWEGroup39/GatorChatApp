import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Observable, map, tap } from 'rxjs';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  public getJsonValue: any;
  public postJsonValue: any;
  
  user= null;

  constructor(private http: HttpClient){

  }
  ngOnInit():void{
    

  }

  /*getUser(username: string, password: string): Observable<any> {
    const url = `http://localhost:8080/api/user/User/${username}/${password}`;
      
    return this.http.get<any>(url).pipe(
       map(response => {
     const user = {
          username: response.username,
           password: response.password,
            userId: response.userId,
         };
          return user;
        }),
        tap(response => {
        console.log('User found:', response);
         })
      );
    }
    */

 

  
}
