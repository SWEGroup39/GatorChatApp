

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

  localID:string=''
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
    this.localID=sessionStorage.getItem('idLog')??''
    this.id = JSON.stringify(sessionStorage.getItem('currentUserI'+this.localID)).replace(/['"]/g, '');
    this.username = JSON.stringify(sessionStorage.getItem('currentUserU'+this.localID)).replace(/['"]/g, '');
    this.email = JSON.stringify(sessionStorage.getItem('currentUserE'+this.localID)).replace(/['"]/g, '');
    this.phone = JSON.stringify(sessionStorage.getItem('currentUserPh'+this.localID)).replace(/['"]/g, '');
    this.fullname = JSON.stringify(sessionStorage.getItem('currentUserF'+this.localID)).replace(/['"]/g, '');
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
    console.log(firstN[0].toUpperCase()+lastN[0].toUpperCase())
    return firstN[0].toUpperCase()+lastN[0].toUpperCase();
  }
  
  goBack1() {
    this.location.back();
  }
   
}
