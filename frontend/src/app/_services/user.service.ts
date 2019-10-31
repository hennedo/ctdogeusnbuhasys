import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {UserModel} from '../_models';
import {Observable} from 'rxjs';
import {environment} from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpClient) { }

  getOne(id: string): Observable<UserModel> {
    return this.http.get<UserModel>(environment.apiBaseUrl + '/user/' + id);
  }
  getAll(): Observable<UserModel[]> {
    return this.http.get<UserModel[]>(environment.apiBaseUrl + '/user');
  }
  update(id: string, user: UserModel): Observable<UserModel> {
    return this.http.put<UserModel>(environment.apiBaseUrl + '/user/' + id, user);
  }
  create(user: UserModel): Observable<UserModel> {
    return this.http.post<UserModel>(environment.apiBaseUrl + '/user', user);
  }
}
