import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { UserService } from 'src/app/service/user.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit{
  username: string ='';
  password: string = '';
  id: string = ''
  email:string=''
  fullname:string=''
  localID:string=''
  constructor(private route: ActivatedRoute, private userService:UserService) {}

  ngOnInit(): void {
    // this.route.queryParams.subscribe(params => {
    //   this.username = params['username'] ?? 'failed';
    //   this.password = params['password'] ?? 'failed';
    //   this.id = params['id'] ?? 'failed';
    //   this.email = params['email']??'failed';
    //   console.log(this.username+ ' ' + this.password + ' ' + this.id + ' ' + this.email)
    // });
    this.localID = sessionStorage.getItem('idLog')??''
    this.username = JSON.stringify(sessionStorage.getItem(`currentUserU`+this.localID)).replace(/['"]/g, '');
    this.password = JSON.stringify(sessionStorage.getItem('currentUserP'+this.localID)).replace(/['"]/g, '');
    this.email = JSON.stringify(sessionStorage.getItem('currentUserE'+this.localID)).replace(/['"]/g, '');
    this.id = JSON.stringify(sessionStorage.getItem('currentUserI'+this.localID)).replace(/['"]/g, '');
    this.fullname = JSON.stringify(sessionStorage.getItem('currentUserF'+this.localID)).replace(/['"]/g, '');
    console.log(JSON.stringify(sessionStorage));
    // localStorage.setItem('currentUserU'+this.id,this.username)
    // localStorage.setItem('currentUserP'+this.id,this.password)
    // localStorage.setItem('currentUserE'+this.id,this.email)
    // localStorage.setItem('currentUserF'+this.id,this.fullname)
    // localStorage.setItem('currentUserI'+this.id,this.id)
  }
  logOut():void{
    sessionStorage.setItem('userLoggedIn','false')
  }

}
