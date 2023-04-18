import { Component ,OnInit, Input} from '@angular/core';
import { Location } from '@angular/common';
import { UserService } from 'src/app/service/user.service';
import { Router } from '@angular/router';
import { Observable, catchError } from 'rxjs';


@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.css']
})
export class SettingsComponent{

  constructor(private location: Location, private user:UserService, private router: Router) {}
  @Input() fullName:string = ''
  @Input() email:string = ''
  @Input() phone:string = ''
  @Input() username:string=''
  firstN:string = ''
  lastN:string = ''
  userID: string=''
  @Input() newPassword:string=''
  @Input() password:string=''
  errorMessage:string=''
  editing:boolean =false;
  localID:string=''
  goBack() {
    this.location.back();
  }

  ngOnInit() {
    this.localID = sessionStorage.getItem('idLog')??''
    this.fullName = JSON.stringify(sessionStorage.getItem('currentUserF'+this.localID)).slice(1,-1);
    this.firstN = this.fullName.substring(0,this.fullName.indexOf(' '))
    this.lastN = this.fullName.substring(this.fullName.indexOf(' ') + 1)
    this.phone = JSON.stringify(sessionStorage.getItem('currentUserPh'+this.localID)).slice(1,-1);
    this.email = JSON.stringify(sessionStorage.getItem('currentUserE'+this.localID)).slice(1,-1);
    this.userID = JSON.stringify(sessionStorage.getItem('currentUserI'+this.localID)).slice(1,-1);
    this.username = JSON.stringify(sessionStorage.getItem('currentUserU'+this.localID)).slice(1,-1);
    // this.password = JSON.stringify(localStorage.getItem('currentUserP')).slice(1,-1);
    
  }

  editFieldFN(element:any){
    this.editing = true;
    
    setTimeout(()=>{
      element.focus()
    })
    
    
  }
  saveFieldFN(){
    this.updateFullName()
    sessionStorage.setItem('currentUserF'+this.localID,this.fullName);
  }
  editFieldPN(element:any){
    this.editing = true;
    
    setTimeout(()=>{
      element.focus()
    })
    
  }
  saveFieldPN(){
    this.updatePhoneNum()
    sessionStorage.setItem('currentUserPh'+this.localID,this.phone);
  }

  editFieldU(element:any){
    this.editing = true;
    
    setTimeout(()=>{
      element.focus()
    })
    
    
  }
  saveFieldU(){
    this.updateUsername()
    sessionStorage.setItem('currentUserU'+this.localID,this.username);
  }

  deleteUser():void{
    this.user.deleteUser(this.userID).subscribe(
      (response)=>{
        alert('User Account Deleted Successfully')
        sessionStorage.removeItem('userLoggedIn')
        this.router.navigateByUrl('/login')

      },
      (error) => console.log(error),
      ()=> console.log('Deleted!')
    );
  }

  updateFullName():void{
    console.log(this.fullName)
    this.user.updateFN(this.userID,this.fullName).subscribe(
      (response)=>{
        console.log(response)
        
        console.log(this.fullName)
      },
      (error)=>{console.log(error)},
      ()=>console.log('Updated Full Name!')
    );
  }
  updatePhoneNum():void{
    
    this.user.updatePhoneNum(this.userID,this.phone).subscribe(
      (response)=>{
        console.log(response)
        
        console.log(this.phone)
      },
      (error)=>{console.log(error)},
      ()=>console.log('Updated Phone Number!')
    );
  }
  updatePassword():void{
    console.log('Old password '+this.password)
    this.user.updatePassword(this.userID,this.newPassword, this.password).subscribe(
      (response)=>{
        console.log(response)
        this.password = this.newPassword;
        
        console.log('Updated password: ' + this.password)
      },
      (error)=>{
        console.log(error)
      },
      ()=>console.log('Updated Password!')
    );
    
  }
  updateUsername():void{
    console.log('Old username: ' + this.username)
    this.user.updateUsername(this.userID,this.username).subscribe(
      (response)=>{
        console.log(response)
        
        console.log('New username: ' + this.username)

      },
      (error)=>{},
      ()=>console.log('Updated Username!')
    );
  }
  saveAs():void{
    this.updatePassword();
    sessionStorage.setItem('currentUserP'+this.localID,this.password);
    window.location.reload()
  }

  cancelChanges():void{
    this.email = this.email;
    this.fullName = this.fullName;
    this.phone = this.phone;
    this.password = this.password;
  }
}
