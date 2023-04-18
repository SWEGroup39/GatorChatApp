import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ConvoService } from 'src/app/service/convo.service';
import { Location } from '@angular/common';
import { UserService } from 'src/app/service/user.service';
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
 localID:string=''
  constructor(private convoService: ConvoService, private route: ActivatedRoute, private location: Location, private user:UserService) {}

  // ngOnInit() {
  //   this.route.queryParams.subscribe(params => {
  //     // this.id = params['id'] ?? 'failed';
  //     // this.username = params['username'] ?? 'failed';
  //     // this.password = params['password'] ?? 'failed';
  //     this.id = JSON.stringify(localStorage.getItem('currentUserI')).replace(/['"]/g, '');
  //     this.username = JSON.stringify(localStorage.getItem('currentUserU')).replace(/['"]/g, '');
  //     this.password = JSON.stringify(localStorage.getItem('currentUserP')).replace(/['"]/g, '');
  //     let currentConvString = localStorage.getItem('currentUserC');
  //     console.log(currentConvString)
  //     console.log(this.id);
  //     console.log(this.username);
  //     console.log(this.password);
    
  
  //     this.convoService.getConvoUserIds(this.username, this.password).subscribe(conversations => {
  //       // this.convos = conversations.map(user => ({
  //       //   id: user,
  //       //   username: ``
  //       // }
  //       //)
  //       //);
        
  //       for (const friendId of JSON.parse(:)) {
  //         this.convoService.getConvoUserName( currentConvString?.friendId).subscribe(item=> {
  //          // convo.username = item.username;
  //          this.convos.push({id: friendId, username: item.username})
  //          // console.log(convo.username)
  //         })
  //       }
  //       console.log(this.convos)

  //     });

  //   });
  // }
  ngOnInit() {
    this.localID = sessionStorage.getItem('idLog')??''
    this.route.queryParams.subscribe(params => {
      this.id = JSON.stringify(sessionStorage.getItem('currentUserI'+this.localID)).replace(/['"]/g, '');
      this.username = JSON.stringify(sessionStorage.getItem('currentUserU'+this.localID)).replace(/['"]/g, '');
      this.password = JSON.stringify(sessionStorage.getItem('currentUserP'+this.localID)).replace(/['"]/g, '');
      let currentConvString = sessionStorage.getItem('currentUserC'+this.localID);
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