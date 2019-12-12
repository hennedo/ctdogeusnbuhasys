import { Injectable } from '@angular/core';
import {ActivatedRouteSnapshot, Resolve, Router, RouterStateSnapshot} from '@angular/router';
import {UserService} from '@/_services';
import { UserModel } from '@/_models';
import {EMPTY, Observable, of} from 'rxjs';
import {mergeMap, take} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class UsersResolverService implements Resolve<UserModel[]> {

  constructor(private users: UserService, private router: Router) {}

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<UserModel[]> | Observable<never> {
    return this.users.getAll().pipe(
      take(1),
      mergeMap(usersData => {
        if (usersData) {
          return of(usersData);
        } else {
          this.router.navigate(['/']);
          return EMPTY;
        }
      })
    );
  }
}
