import {Component, OnInit} from '@angular/core';
import {CurrentUserService} from './_services';
import {UserModel} from './_models';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {

  isNavbarCollapsed = false;
  currentUser: UserModel;
  constructor(private currentUserService: CurrentUserService) {}

  ngOnInit() {
    this.currentUserService.get().subscribe(u => {
      this.currentUser = u;
    });
  }
}
