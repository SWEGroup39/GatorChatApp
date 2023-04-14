import { Injectable } from '@angular/core';
import { HttpClient, HttpParams, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { catchError } from 'rxjs/operators';
import { of } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ConvoService {
  isLoggedIn: boolean = false;
  url: string = `http://localhost:8080/api/users`;

  constructor(private http: HttpClient) {
    
  }

  getConvoUserIds(username: string, password: string): Observable<string[]> {
    const requestBody: any = {
      username: username,
      password: password
    };
    return this.http.post<{ current_conversations: string[] }>(`${this.url}/UserInternal`, requestBody)
      .pipe(map(response => response.current_conversations));
  }

  getConvoUserName(id: string): Observable<{ username: string }> {
    return this.http.get<{ username: string }>(
      `${this.url}/${id}`
    ).pipe(
      map(response => ({ username: response.username })),
      catchError(() => of({ username: 'DOES NOT EXIST' }))
    );
  }
  
}



