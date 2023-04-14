import { Observable } from 'rxjs';
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
  @Input() fullname!:string;
  @Input() phone!:string;

  
  userIDCount: number = 2;
  userIDFinal: string=``;

  user: User = {
    username: this.username,
    password: this.password,
    email:this.email,
    full_name:this.fullname,
    phone_number:this.phone,
    current_conversations:[]
  }
  constructor(private userService: UserService, private router:Router){
    

  }
  ngOnInit(): void {
    //this.onCreateUser();
  
  }

  resetForm(form? :NgForm){
    if(form != null)
    form.reset();
    this.user = {
      username: '',
      password: '',
      email: '',
      user_id:'',
      phone_number:'',
      full_name: ''

    }
  }
  onCreateUser():void{
   
    
    this.userService.createUser(this.user).subscribe(
      
      (response) => {
        
        
        alert(`User created successfully!`)
        
        this.router.navigateByUrl('/login')
        this.resetForm()
      },
      (error: any) => {console.log(error)
      console.log(this.user)
      },
      () => console.log('Done creating user')
    );
  }
  zeroPad(num:number,count:number):string
{
  var numZeropad = num + '';
  while(numZeropad.length < count) {
  numZeropad = "0" + numZeropad;
  }
  console.log(numZeropad)
return numZeropad;
}

  
  userCount(): void{
    if(this.userIDCount != 9999){
      this.userIDCount++;
    }
  
  }
  userIDNext():string{
    
    this.userService.getNextID().subscribe(response =>{
      
      this.userIDFinal = response.toString();
      console.log(this.userIDFinal);
      
    });
    return this.userIDFinal;
  }
  
 
}
