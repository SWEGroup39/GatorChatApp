import { Component ,OnInit, Input} from '@angular/core';
import { Location } from '@angular/common';
import { UserService } from 'src/app/service/user.service';
import { Router } from '@angular/router';

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
  firstN:string = ''
  lastN:string = ''
  userID: string=''

  goBack() {
    this.location.back();
  }

  ngOnInit() {
    this.fullName = JSON.stringify(localStorage.getItem('currentUserF')).replace(/['"]/g, '');
    this.firstN = this.fullName.substring(0,this.fullName.indexOf(' '))
    this.lastN = this.fullName.substring(this.fullName.indexOf(' ') + 1)
    this.phone = JSON.stringify(localStorage.getItem('currentUserPh')).replace(/['"]/g, '');
    this.email = JSON.stringify(localStorage.getItem('currentUserE')).replace(/['"]/g, '');
    this.userID = JSON.stringify(localStorage.getItem('currentUserI')).replace(/['"]/g, '');
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
        alert('Updated Full Name Successfully')
        this.fullName = JSON.stringify(localStorage.setItem('currentUserF',this.fullName)).replace(/['"]/g, '');

      },
      (error)=>console.log(error),
      ()=>console.log('Updated Full Name!')
    );
  }
  updatePhoneNum():void{
    
    this.user.updatePhoneNum(this.userID,this.phone).subscribe(
      (response)=>{
        alert('Updated Phone Number Successfully')
        this.fullName = JSON.stringify(localStorage.setItem('currentUserPh',this.phone)).replace(/['"]/g, '');

      },
      (error)=>console.log(error),
      ()=>console.log('Updated Phone Number!')
    );
  }

  cancelChanges():void{
    this.email = this.email;
    this.fullName = this.fullName;
    this.phone = this.phone;
  }
}
