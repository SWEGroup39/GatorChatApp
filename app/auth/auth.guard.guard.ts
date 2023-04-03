import { UserService } from './../service/user.service';
import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree } from '@angular/router';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthGuardGuard implements CanActivate {

  constructor(private userService:UserService, private router: Router){}
  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): boolean {
      const isAuthenticated = JSON.parse(sessionStorage.getItem('userLoggedIn')||'{}');
      if(isAuthenticated === true){
        return true;
      }
      else{
        this.router.navigateByUrl('/login');
        return false;
      }
  }

  
  
}
