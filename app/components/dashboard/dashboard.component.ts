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
  fullname:string=''
  constructor(private route: ActivatedRoute) {}

  ngOnInit(): void {
    // this.route.queryParams.subscribe(params => {
    //   this.username = params['username'] ?? 'failed';
    //   this.password = params['password'] ?? 'failed';
    //   this.id = params['id'] ?? 'failed';
    //   this.email = params['email']??'failed';
    //   console.log(this.username+ ' ' + this.password + ' ' + this.id + ' ' + this.email)
    // });
    this.username = JSON.stringify(localStorage.getItem('currentUser')).replace(/['"]/g, '');
    this.password = JSON.stringify(localStorage.getItem('currentUserP')).replace(/['"]/g, '');
    this.email = JSON.stringify(localStorage.getItem('currentUserE')).replace(/['"]/g, '');
    this.id = JSON.stringify(localStorage.getItem('currentUserI')).replace(/['"]/g, '');
    this.fullname = JSON.stringify(localStorage.getItem('currentUserF')).replace(/['"]/g, '');
    console.log(JSON.stringify(localStorage));
  }
  logOut():void{
    sessionStorage.setItem('userLoggedIn','false')
  }

}
