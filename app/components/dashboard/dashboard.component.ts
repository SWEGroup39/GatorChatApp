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
  constructor(private route: ActivatedRoute) {}

  ngOnInit(): void {
    this.route.queryParams.subscribe(params => {
      this.username = params['username'] ?? 'failed';
      this.password = params['password'] ?? 'failed';
      this.id = params['id'] ?? 'failed';
    });
  }

}
