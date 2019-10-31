import { Injectable } from '@angular/core';
import {UserModel} from '../_models';
import {BehaviorSubject, Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class CurrentUserService {
  private subject: BehaviorSubject<UserModel>;
  private currentUser: UserModel;

  constructor() {
    this.subject = new BehaviorSubject(null);
    this.currentUser = null;
  }

  set(u: UserModel) {
    this.subject.next(u);
    this.currentUser = u;
  }

  get(): Observable<UserModel> {
    return this.subject;
  }
  getValue(): UserModel {
    return this.currentUser;
  }

  getHeader() {
    return {
      headers: {Authorization: 'Basic ' + btoa(this.currentUser.name + ':' + this.currentUser.pin)}
    };
  }
}
