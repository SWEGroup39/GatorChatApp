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
 
    this.username = JSON.stringify(localStorage.getItem('currentUserU')).replace(/['"]/g, '');
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
