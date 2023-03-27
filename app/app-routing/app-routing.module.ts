import { AuthGuardGuard } from './../auth/auth.guard.guard';
import { NotificationComponent } from './../components/notification/notification.component';
import { SettingsComponent } from './../components/settings/settings.component';
import { ProfileComponent } from './../components/profile/profile.component';
import { ContactsComponent } from './../components/contacts/contacts.component';
import { MessagesComponent } from './../components/messages/messages.component';

import { DashboardComponent } from './../components/dashboard/dashboard.component';
import { SignUpComponent } from '../components/sign-up/sign-up.component';
import { LoginComponent } from '../components/login/login.component';
import { NgModule, Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import {RouterModule, Routes} from'@angular/router'
import { AboutComponent } from '../components/about/about.component';
import { HeaderComponent } from '../components/header/header.component';
import { HomeComponent } from '../components/home/home.component';
import { ChatListComponent } from '../components/chat-list/chat-list.component';
const appRoutes = [
  
  {path: 'home', component: HomeComponent},
  {path: 'login', component: LoginComponent},
  {path: 'signup', component: SignUpComponent},
  {path: 'about', component: AboutComponent},
  {path:'dashboard',component:DashboardComponent, canActivate:[AuthGuardGuard]},
  {path:'dashboard/:username/:password',component:DashboardComponent, canActivate:[AuthGuardGuard]},
  {path: 'messages', component:MessagesComponent, canActivate:[AuthGuardGuard]},
  {path:'contacts', component: ContactsComponent, canActivate:[AuthGuardGuard]},
  {path:'profile', component: ProfileComponent, canActivate:[AuthGuardGuard]},
  {path:'settings', component: SettingsComponent, canActivate:[AuthGuardGuard]},
  {path:'notification', component: NotificationComponent, canActivate:[AuthGuardGuard]},
  {path: 'chat-list', component: ChatListComponent, canActivate:[AuthGuardGuard]},
  {path: '**', redirectTo:'/home'}

];

@NgModule({
  
  imports: [
    RouterModule.forRoot(appRoutes)
  ],
  exports:[RouterModule],
  providers:[AuthGuardGuard]
})
export class AppRoutingModule { }
