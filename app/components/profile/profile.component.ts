
import { Component } from '@angular/core';
import { Location } from '@angular/common';
import { ActivatedRoute } from '@angular/router';
import { User } from 'src/app/interface/user';


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
  emailAddress: string = 'test@ufl.edu';
  phone:string = '';
  user:User = {
    username: this.username,
    password: this.password,
    user_id: this.id,
    email: this.emailAddress
  }
  editing: boolean = false;
  

  constructor(private route: ActivatedRoute, private location: Location) {}

  ngOnInit() {
    
    this.route.queryParams.subscribe(params => {
      this.id = params['id'] ?? 'failed';
      this.username = params['username'] ?? 'failed';
      this.password = params['password'] ?? 'failed';
      
      console.log(this.id + ' ' + this.username + ' ' + this.password);
  
    });
    const value = sessionStorage.getItem(this.username);
    
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
