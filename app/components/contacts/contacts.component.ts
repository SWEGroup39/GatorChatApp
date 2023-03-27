import { Component } from '@angular/core';
import { Location } from '@angular/common';

@Component({
  selector: 'app-contacts',
  templateUrl: './contacts.component.html',
  styleUrls: ['./contacts.component.css']
})
export class ContactsComponent {

  constructor( private location: Location) {}


  goBack() {
    this.location.back();
  }
}
