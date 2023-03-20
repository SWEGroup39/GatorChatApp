import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { LoginComponent } from './components/login/login.component';
import { FlexLayoutModule } from '@angular/flex-layout';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatIconModule} from '@angular/material/icon';
import {HttpClientModule} from '@angular/common/http';
import { SignUpComponent } from './components/sign-up/sign-up.component';
import { AppRoutingModule } from './app-routing/app-routing.module';
import { HeaderComponent } from './components/header/header.component';
import { AboutComponent } from './components/about/about.component';
import { HomeComponent } from './components/home/home.component';
import {AccordionModule} from 'primeng/accordion';     //accordion and accordion tab
import {StyleClassModule} from 'primeng/styleclass';
import { FooterComponent } from './components/footer/footer.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { MessagesComponent } from './components/messages/messages.component';
import { ContactsComponent } from './components/contacts/contacts.component';
import { ProfileComponent } from './components/profile/profile.component';
import { SettingsComponent } from './components/settings/settings.component';
import { NotificationComponent } from './components/notification/notification.component';


@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    SignUpComponent,
    HeaderComponent,
    AboutComponent,
    HomeComponent,
    FooterComponent,
    DashboardComponent,
    MessagesComponent,
    ContactsComponent,
    ProfileComponent,
    SettingsComponent,
    NotificationComponent,
    
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    FlexLayoutModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatCardModule,
    MatToolbarModule,
    MatIconModule,
    HttpClientModule,
    AccordionModule,
    StyleClassModule, 
    FormsModule,
    
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {
  
 }
