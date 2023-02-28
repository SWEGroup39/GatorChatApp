import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  public getJsonValue: any;
  public postJsonValue: any;
  
  user= null;

  constructor(private http: HttpClient){

  }
  ngOnInit():void{
    this.getMethod();

  }

  // public getMethod(){
  //   this.http.get('http://localhost:8080/api/users').subscribe((data) => {
  //     console.log(data);
  //     this.getJsonValue = data;
  //   });
    
  // }

  public getMethod(){
    this.http.get('http://localhost:8080/api/users').subscribe((data) => {
      console.log(data);
      this.getJsonValue = data;
    });
    
  }

  public postMethod(){
    
    this.http.post('http://localhost:8080/api/users',{}).subscribe((data) => {
      console.log(data);
      this.postJsonValue = data;              

    });
    
  }

  
}
