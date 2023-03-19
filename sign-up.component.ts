import { Router } from '@angular/router';
import { Component, OnInit, Input } from '@angular/core';
import { NgForm } from '@angular/forms';
import { User } from 'src/app/interface/user';
import { UserService } from './../../service/user.service';

@Component({
  selector: 'app-sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css']
})
export class SignUpComponent implements OnInit{
  @Input() username!: string;
  @Input() password!:string;
  @Input() email!:string;
  
  
  constructor(private userService: UserService, private router:Router){}
  ngOnInit(): void {
    //this.onCreateUser();
   
  }
   user: User = {
    username :this.username,
    password: this.password,
    email:this.email,
    'user_id':"6754"
   
  }

  resetForm(form? :NgForm){
    if(form != null)
    form.reset();
    this.user = {
      username: '',
      password: '',
      email: '',
      user_id:''

    }
  }
  onCreateUser():void{
    this.userService.createUser(this.user).subscribe(
      (response) => {
        this.resetForm()
        alert(`User created successfully!`)
        this.router.navigateByUrl('/login')
      },
      (error: any) => console.log(error),
      () => console.log('Done creating user')
    );
  }

}
