import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ConvoService } from 'src/app/service/convo.service';
import { Location } from '@angular/common';
@Component({
  selector: 'app-chat-list-component',
  templateUrl: './chat-list-component.component.html',
  styleUrls: ['./chat-list-component.component.css']
})
export class ChatListComponentComponent {
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
      this.id = params['id'] ?? 'failed';
      this.username = params['username'] ?? 'failed';
      this.password = params['password'] ?? 'failed';
  
      this.convoService.getConvoUserIds(this.username, this.password).subscribe(conversations => {
        this.convos = conversations.map(user => ({
          id: user,
          username: ``
        }));
        
        for (const convo of this.convos) {
          this.convoService.getConvoUserName(convo.id).subscribe(item=> {
            convo.username = item.username;
            console.log(convo.username)
          })
        }
        console.log(this.convos)

      });

    });
  }

  goBack() {
    this.location.back();
  }
}
