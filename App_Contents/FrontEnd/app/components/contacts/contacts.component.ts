

import { Component ,OnInit,Input} from '@angular/core';
import { Location } from '@angular/common';
import { UserService } from 'src/app/service/user.service';
import { ConvoService } from 'src/app/service/convo.service';


@Component({
  selector: 'app-contacts',
  templateUrl: './contacts.component.html',
  styleUrls: ['./contacts.component.css']
})
export class ContactsComponent {
  ConvoService: any;

  constructor( private location: Location, private userService:UserService,  private convoService:ConvoService) {}
  // currentConvo = localStorage.getItem('currentUserC');
 // contactList: any[] = [];

  users: {
    id: string,
    username: string,
  }[] =[];

  @Input() searchValue:string=``;
  size!: number;
  id:string=``
  username:string=``
  currentID:string=''
  localID:string=''
  currentUserUsername:string=''
  odd:boolean = true;

  ngOnInit(){

    this.localID = sessionStorage.getItem('idLog')??''
    this.currentID = sessionStorage.getItem('currentUserI'+this.localID)??''
    this.currentUserUsername = sessionStorage.getItem('currentUserU'+this.localID)??''
    let currentConvString = sessionStorage.getItem('currentUserC'+this.localID);
   
    for (const friendId of JSON.parse(currentConvString?? '' )) {
      this.convoService.getConvoUserName(friendId).subscribe(item => {
        this.users.push({id: friendId, username: item.username})

      console.log(friendId + '--> ' + item.username)
      });
    }

    console.log(this.users);

    console.log(this.users.length)
    this.contactListArray()
  }

  contactListArray(){
    if(this.users.length % 2 == 0){
      this.odd = false
    }
    else{
      this.odd = true
    }
    return this.odd
  }

  modifyContactsList(add:boolean, id:string)
  {
    let currentConvString = sessionStorage.getItem('currentUserC' + this.localID);
    let currentConvArray = JSON.parse(currentConvString ?? '');

    if(add)
      currentConvArray.push(id);
    else
      currentConvArray = currentConvArray.filter((friendId: string) => friendId !== id);

    sessionStorage.setItem('currentUserC' + this.localID, JSON.stringify(currentConvArray));
    console.log( currentConvArray)

  }
  

  searchContact():void{
    this.userService.searchContact(this.searchValue).subscribe(
      (response)=>{
        console.log('User found!')
        this.username = this.searchValue.substring(0, this.searchValue.indexOf('#')).toString()
        this.id = this.searchValue.substring(this.searchValue.indexOf('#')+1).toString()

        this.users.push({username: this.username, id:this.id});
        console.log(sessionStorage)

        this.modifyContactsList(true, this.id);

        this.addConversationID()

      },
      (error)=>{
        console.log(error)
        alert(this.searchValue+' not found!')
        window.location.reload()
      },
      ()=>{
   
      }
      
    );
  }

  deleteContact(deleteID:string, deleteUsername:string):void{
    
    // Update users list 
    this.deleteContactsID(deleteID);
    let contactListN = this.users.filter(item =>  item.id != deleteID && item.username != deleteUsername);
    console.log(contactListN);
    this.users = contactListN;  

    this.modifyContactsList(false, deleteID);
    
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
