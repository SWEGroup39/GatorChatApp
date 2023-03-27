import { Component } from '@angular/core';
import { Location } from '@angular/common';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent {
  id: string = '';
  username: string ='';
  password: string = '';
  birthday: Date = new Date("01/01/1900");
  firstName: string ='Moe';
  lastName: string = 'Mama';

  constructor(private route: ActivatedRoute, private location: Location) {}

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      this.id = params['id'] ?? 'failed';
      this.username = params['username'] ?? 'failed';
      this.password = params['password'] ?? 'failed';

      console.log(this.id + ' ' + this.username + ' ' + this.password);
    });
  }

  goBack() {
    this.location.back();
  }
}
