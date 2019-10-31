import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {CurrentUserService} from './current-user.service';
import {Observable} from 'rxjs';
import {ResponseMessageModel} from '../_models';
import {environment} from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class BuyService {

  constructor(private http: HttpClient, private currentUser: CurrentUserService) { }

  buy(id: string): Observable<ResponseMessageModel> {
    return this.http.post<ResponseMessageModel>(environment.apiBaseUrl + '/buy', { product: id }, this.currentUser.getHeader());
  }
}
