import { Component, OnInit } from '@angular/core';
import {CurrentUserService} from '../../_services';
import {UserModel} from '@/_models';
import {UserModule} from '@/user/user.module';
import {ActivatedRoute} from '@angular/router';
import {pipe} from 'rxjs';
import {first, mergeMap, take} from 'rxjs/operators';

@Component({
  selector: 'app-index',
  templateUrl: './index.component.html',
  styleUrls: ['./index.component.scss']
})
export class IndexComponent implements OnInit {

  users: UserModel[];
  constructor(private currentUserService: CurrentUserService, private route: ActivatedRoute) { }
  ngOnInit() {
    this.currentUserService.set(null);
    this.route.data.subscribe(data => {
      this.users = data.users;
    });
  }

}
