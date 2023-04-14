import { DashboardComponent } from './../dashboard/dashboard.component';


import { UserService } from './../../service/user.service';
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
    email : this.email,
    password: this.password,
  
  }

  constructor(private userService: UserService, private router: Router){}

  onGetUsers():void{
    
    this.userService.getUsers().subscribe(
      (response) => console.log(response),
      (error:any) => alert('Incorrect username and password please try again'),
      () => console.log('Done getting users')
    );
  }


  onGetUser():void{

    this.userService.getUser(this.user.password, this.user.email).subscribe(
      (response) => {
        this.resetForm()
        this.submitSuccess = true;
        this.isloggedIn(this.submitSuccess);
        console.log('Logged In')
       const { user_id, username, password,email,current_conversations,phone_number,full_name } = response;
       localStorage.setItem(`currentUserU`, username)
       localStorage.setItem(`currentUserP`,password)
       localStorage.setItem(`currentUserE`,email)
       localStorage.setItem(`currentUserI`, user_id)
       localStorage.setItem(`currentUserPh`,phone_number)
       localStorage.setItem('currentUserF',full_name)
       localStorage.setItem('currentUserC',JSON.stringify(current_conversations))
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
    
  }
  get loggedIn():boolean{
    return this.userService.isLoggedIn;
  }


  ngOnInit():void{
 

  }

  resetForm(form? :NgForm){
    if(form != null)
    form.reset();
    this.user = {
      email: '',
      password: ''

    }
  }


 

  
}