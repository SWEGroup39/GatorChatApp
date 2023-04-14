import { ResourceLoader } from '@angular/compiler';
import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ConvoService } from './../../service/convo.service';
import { Location } from '@angular/common';

@Component({
  selector: 'app-chat-list',
  templateUrl: './chat-list.component.html',
  styleUrls: ['./chat-list.component.css']
})
export class ChatListComponent{
  convos: {
    id: string,
    username: string,
  }[] =[];


  id: string = '';
  username: string ='';
  password: string = '';
 
  constructor(private convoService: ConvoService, private route: ActivatedRoute, private location: Location) {}


  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      this.id = JSON.stringify(localStorage.getItem('currentUserI')).replace(/['"]/g, '');
      this.username = JSON.stringify(localStorage.getItem('currentUserU')).replace(/['"]/g, '');
      this.password = JSON.stringify(localStorage.getItem('currentUserP')).replace(/['"]/g, '');
      let currentConvString = localStorage.getItem('currentUserC');
      
      //debug purposes 
      console.log(currentConvString)
      console.log(this.id);
      console.log(this.username);
      console.log(this.password);

     
        for (const friendId of JSON.parse(currentConvString?? '' )) {
          console.log(friendId)
          this.convoService.getConvoUserName(friendId).subscribe(item => {
            this.convos.push({id: friendId, username: item.username})
          });
        }
        console.log(this.convos);
      });
  }

  goBack() {
    this.location.back();
  }
}