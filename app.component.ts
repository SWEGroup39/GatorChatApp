import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  chatInputMessage: string = "";
  currentUser = {
    name: 'John ',
    id: 1,
    profileImageUrl:
    'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTKQmFYe2KZvQcnKEfGNICxM4I4udEh_-uG90chKLlXMx2HDGPr_ODubOdkpUFdJVGSKs0&usqp=CAU',

  }

  user1= {
    name: 'Jane ',
    id: 2,
    profileImageUrl:
    'https://ps.w.org/user-avatar-reloaded/assets/icon-256x256.png?rev=2540745',

  }

  user2= {
    name: 'Jill ',
    id: 3,
    profileImageUrl:
    'https://e7.pngegg.com/pngimages/348/800/png-clipart-man-wearing-blue-shirt-illustration-computer-icons-avatar-user-login-avatar-blue-child.png',
    

  }

  chatMessages: {
    user: any,
    message: string,
    created_at: number
  }[] =[
    {
      user: this.currentUser,
      message: "Yo whats up?",
      created_at: Date.now()
    },
    {
      user: this.user1,
      message: "Not much, hbu?",
      created_at: Date.now()
    },
    {
      user: this.user2,
      message: "We should go on a vacay or smt",
      created_at: Date.now()
    },
    
  ]

  title = 'chat-app';

  send()
  {
   //var button = document.getElementById("send-button");

    // if(this.chatInputMessage.length <= 0)
    // {
    //   button?.style.backgroundColor == "grey";
    // }
    
      this.chatMessages.push({
        message: this.chatInputMessage,
        user: this.currentUser,
        created_at: Date.now()
      })
      this.chatInputMessage = "";
    
    
  }
}
