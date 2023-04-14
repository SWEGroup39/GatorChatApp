import { Component ,OnInit} from '@angular/core';
import { Location } from '@angular/common';

@Component({
  selector: 'app-contacts',
  templateUrl: './contacts.component.html',
  styleUrls: ['./contacts.component.css']
})
export class ContactsComponent {

  constructor( private location: Location) {}
  currentConvo = localStorage.getItem('currentUserC');
  contactList: any[] = [];
  size!: number;
  ngOnInit(){

  }

  contactListArray(){
    if(this.contactList.length == 0){
      this.size = 0;
    }
    else{
      this.size = this.contactList.length;
    }
  }

  goBack() {
    this.location.back();
  }
}