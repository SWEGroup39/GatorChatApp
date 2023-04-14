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

  ngOnInit(){

  }

  goBack() {
    this.location.back();
  }
}