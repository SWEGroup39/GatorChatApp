import { Component ,OnInit,Input} from '@angular/core';
import { Location } from '@angular/common';
import { UserService } from 'src/app/service/user.service';

@Component({
  selector: 'app-contacts',
  templateUrl: './contacts.component.html',
  styleUrls: ['./contacts.component.css']
})
export class ContactsComponent {

  constructor( private location: Location, private userService:UserService) {}
  // currentConvo = localStorage.getItem('currentUserC');
  contactList: any[] = [];
  @Input() searchValue:string=``;
  size!: number;
  id:string=``
  username:string=``
  currentID:string=''
  localID:string=''
  currentUserUsername:string=''
  ngOnInit(){
    // for(let i = 0; i < this.contactList.length; i++){
    //   console.log(this.contactList.at(i))
    // }
    // this.contactList = JSON.parse(localStorage.getItem("this.username")??'')

    this.localID = sessionStorage.getItem('idLog')??''
    this.currentID = sessionStorage.getItem('currentUserI'+this.localID)??''
    this.currentUserUsername = sessionStorage.getItem('currentUserU'+this.localID)??''
    this.contactList = JSON.parse(localStorage.getItem("contact"+this.currentID)??'[]')
    console.log(this.contactList.length)
  }

  contactListArray(){
    if(this.contactList.length == 0){
      this.size = 0;
    }
    else{
      this.size = this.contactList.length;
    }
    return this.size
  }

  searchContact():void{
    // this.username = this.searchValue.substring(0, this.searchValue.indexOf('#')).toString()
    // this.id = this.searchValue.substring(this.searchValue.indexOf('#')+1).toString()
    this.userService.searchContact(this.searchValue).subscribe(
      (response)=>{
        console.log('User found!')
        this.username = this.searchValue.substring(0, this.searchValue.indexOf('#')).toString()
        this.id = this.searchValue.substring(this.searchValue.indexOf('#')+1).toString()
        this.contactList.push({username: this.username, id:this.id})
        localStorage.setItem("contact"+this.currentID,JSON.stringify(this.contactList))
        console.log(sessionStorage)
        this.addConversationID()
      },
      (error)=>{
        console.log(error)
      },
      ()=>{
        // this.contactList = JSON.parse(localStorage.getItem("contact"+this.currentID)??' ')
      }
      
    );
  }

  deleteContact(deleteID:string, deleteUsername:string):void{
    
    let newContactList = this.contactList.filter(function(item){
      return item.id != deleteID && item.username != deleteUsername
    })
    this.contactList = newContactList
    localStorage.setItem("contact"+this.currentID,JSON.stringify(newContactList))
    this.deleteContactsID(deleteID)
  }

  addConversationID():void{
    this.userService.addConversationID(this.currentID,this.id).subscribe(
      (response)=>{
        console.log(response)
      },
      (error)=>{
        console.log(error)
      }
      
    );
  }

  deleteContactsID(deleteID:string):void{
    this.userService.deleteIDCovo(this.currentID, deleteID).subscribe(
      (response)=>{
        console.log(response)
      },
      (error)=>{
        console.log(error)
      }
    );
  }



  goBack() {
    this.location.back();
  }
}