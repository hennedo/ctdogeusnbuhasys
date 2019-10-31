import { Component, OnInit } from '@angular/core';
import {CurrentUserService} from '../../_services';

@Component({
  selector: 'app-buy',
  templateUrl: './buy.component.html',
  styleUrls: ['./buy.component.scss']
})
export class BuyComponent implements OnInit {

  constructor(private currentUser: CurrentUserService) { }

  ngOnInit() {
    this.currentUser.set({
      name: 'Henne',
      pin: '1234',
      verified: false,
      gravatar: false,
      imageUrl: 'lo',
      balance: 0
    });
  }

}
