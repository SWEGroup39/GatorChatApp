import { UserService } from 'src/app/service/user.service';
import { DashboardComponent } from './../dashboard/dashboard.component';

import { HttpClient } from '@angular/common/http';
import { Component, OnInit, Input, ViewChild, Injectable } from '@angular/core';
import { Observable, map, tap } from 'rxjs';
import { User } from '../../interface/user';
import { NgForm } from '@angular/forms';
import { Router } from '@angular/router';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})

export class LoginComponent implements OnInit {
  @Input() email: string =``;
  @Input() password: string = ``;
  
  submitSuccess:boolean=false;
  user: User = {
    username :this.email,
    password: this.password,
  
  }
  idLog:string=''
  constructor(private userService: UserService, private router: Router){}

  onGetUsers():void{
    
    this.userService.getUsers().subscribe(
      (response) => console.log(response),
      (error:any) => alert('Incorrect username and password please try again'),
      () => console.log('Done getting users')
    );
  }


  onGetUser():void{

    this.userService.getUser(this.user.password, this.user.username).subscribe(
      (response) => {
        this.resetForm()
        this.submitSuccess = true;
        this.isloggedIn(this.submitSuccess);
        console.log('Logged In')
        this.idLog = sessionStorage.getItem('idLog')??''
       const { user_id, username, password,email,current_conversations,phone_number,full_name } = response;
       sessionStorage.setItem(`currentUserU`+this.idLog, username)
       sessionStorage.setItem(`currentUserP`+this.idLog, password)
       sessionStorage.setItem(`currentUserE`+this.idLog,email)
       sessionStorage.setItem(`currentUserI`+this.idLog, user_id)
       sessionStorage.setItem(`currentUserPh`+this.idLog,phone_number)
       sessionStorage.setItem('currentUserF'+this.idLog,full_name)
       sessionStorage.setItem('currentUserC'+this.idLog,JSON.stringify(current_conversations))
       this.router.navigate(['/dashboard']);

      },
      (error:any) => {alert(`Username or Password is incorrect! Please try again`)
      this.resetForm()
    },
      () => console.log('Done getting user and it exists')
    );
  }

   isloggedIn(isLogged:boolean){
    this.userService.isLoggedIn = isLogged;
    
    sessionStorage.setItem('userLoggedIn','true');
    sessionStorage.setItem('idLog',this.userService.loggedInUser)
    
  }
  get loggedIn():boolean{
    return this.userService.isLoggedIn;
  }

  // onCreateUser():void{
  //   this.userService.createUser(this.user).subscribe(
  //     (response) => console.log(response),
  //     (error: any) => console.log(error),
  //     () => console.log('Done creating user')
  //   );
  // }

 
  ngOnInit():void{
    //this.onGetUsers();
   //this.onCreateUser();
    //this.onGetUser();
    

  }

  resetForm(form? :NgForm){
    if(form != null)
    form.reset();
    this.user = {
      username: '',
      password: ''

    }
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