

import { Component } from '@angular/core';
import { Location } from '@angular/common';
import { ActivatedRoute } from '@angular/router';
import { User } from 'src/app/interface/user';
import { LoginComponent } from '../login/login.component';
import { UserService } from 'src/app/service/user.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent {
  
  id: string = '';
  username: string ='';
  password: string = '';
  birthday: Date = new Date("01/01/1900");
  firstName: string ='Moe';
  lastName: string = 'Mama';
  email: string = ``;
  phone:string = '';
  
  editing: boolean = false;


  constructor(private route: ActivatedRoute, private location: Location, private userService:UserService) {}
  
  ngOnInit() {
   
    // this.route.queryParams.subscribe(params => {
    //   this.id = params['id'] ?? 'failed';
    //   this.username = params['username'] ?? 'failed';
    //   this.password = params['password'] ?? 'failed';
    //   this.email = params['email'] ?? 'failed';
    //   console.log(params)
    //   console.log(this.id + ' ' + this.username + ' ' + this.password);
  
    // });
    this.id = JSON.stringify(sessionStorage.getItem('currentUserI')).replace(/['"]/g, '');
    this.username = JSON.stringify(sessionStorage.getItem('currentUserU')).replace(/['"]/g, '');
    this.email = JSON.stringify(sessionStorage.getItem('currentUserE')).replace(/['"]/g, '');
    
    
  }

  goBack() {
    this.location.back();
  }
  edit(element:any){
    this.editing = true;
    setTimeout(()=>{
      element.focus();
    })
    this.storePhone();
    
  }
  public storePhone(){
    const value = this.phone;
    sessionStorage.setItem(this.username,value);
  }
  goBack1() {
    this.location.back();
  }
   
}
