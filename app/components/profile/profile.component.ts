

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
  fullname:string ='';
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
    this.id = JSON.stringify(localStorage.getItem('currentUserI')).replace(/['"]/g, '');
    this.username = JSON.stringify(localStorage.getItem('currentUserU')).replace(/['"]/g, '');
    this.email = JSON.stringify(localStorage.getItem('currentUserE')).replace(/['"]/g, '');
    this.phone = JSON.stringify(localStorage.getItem('currentUserPh')).replace(/['"]/g, '');
    console.log(this.phone)
    this.fullname = JSON.stringify(localStorage.getItem('currentUserF')).replace(/['"]/g, '');
  }

  goBack() {
    this.location.back();
  }
  edit(element:any){
    this.editing = true;
    setTimeout(()=>{
      element.focus();
    })
 
    
  }

  getInitials(fullName: string){
    let firstN = fullName.substring(0, fullName.indexOf(' ')).toString()
    let lastN = fullName.substring(fullName.indexOf(' ') + 1).toString()
    console.log(firstN[0]?.toUpperCase()+lastN[0]?.toUpperCase())
    return firstN[0]?.toUpperCase()+lastN[0]?.toUpperCase();
  }
  
  goBack1() {
    this.location.back();
  }
   
}
