import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

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
  constructor(private route: ActivatedRoute) {}

  ngOnInit(): void {
    // this.route.queryParams.subscribe(params => {
    //   this.username = params['username'] ?? 'failed';
    //   this.password = params['password'] ?? 'failed';
    //   this.id = params['id'] ?? 'failed';
    //   this.email = params['email']??'failed';
    //   console.log(this.username+ ' ' + this.password + ' ' + this.id + ' ' + this.email)
    // });
    this.username = JSON.stringify(sessionStorage.getItem('currentUserU')).replace(/['"]/g, '');
    this.password = JSON.stringify(sessionStorage.getItem('currentUserP')).replace(/['"]/g, '');
    this.email = JSON.stringify(sessionStorage.getItem('currentUserE')).replace(/['"]/g, '');
    this.id = JSON.stringify(sessionStorage.getItem('currentUserI')).replace(/['"]/g, '');
  }
  logOut():void{
    sessionStorage.setItem('userLoggedIn','false')
  }

}
