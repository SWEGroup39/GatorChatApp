import { Router } from '@angular/router';
import { Component } from '@angular/core';
import { MatIconRegistry } from '@angular/material/icon';
import { DomSanitizer } from '@angular/platform-browser';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'chatG-app';
  constructor(
    private matIconRegistry: MatIconRegistry,
    private domSanitizer: DomSanitizer,
    public router:Router
  ) {
  this.matIconRegistry.addSvgIcon(
  'gator',
    this.domSanitizer.bypassSecurityTrustResourceUrl("../assets/alligator-crocodile-icon.svg")
  );
  this.matIconRegistry.addSvgIcon(
    'message',
      this.domSanitizer.bypassSecurityTrustResourceUrl("../assets/message-square-lines-svgrepo-com.svg")
    );
    this.matIconRegistry.addSvgIcon(
      'contacts',
        this.domSanitizer.bypassSecurityTrustResourceUrl("../assets/contacts-alt-svgrepo-com.svg")
      );
      this.matIconRegistry.addSvgIcon(
        'profile',
          this.domSanitizer.bypassSecurityTrustResourceUrl("../assets/profile-svgrepo-com.svg")
        );

        this.matIconRegistry.addSvgIcon(
          'settings',
            this.domSanitizer.bypassSecurityTrustResourceUrl("../assets/settings-svgrepo-com.svg")
          );
          this.matIconRegistry.addSvgIcon(
            'notification',
              this.domSanitizer.bypassSecurityTrustResourceUrl("../assets/notifications-svgrepo-com.svg")
            );
            this.matIconRegistry.addSvgIcon(
              'dashboard',
                this.domSanitizer.bypassSecurityTrustResourceUrl("../assets/dashboard-svgrepo-com.svg")
              );
              this.matIconRegistry.addSvgIcon(
                'pencil',
                this.domSanitizer.bypassSecurityTrustResourceUrl("../assets/pencil-edit-button-svgrepo-com.svg")
              );
}
  
}
