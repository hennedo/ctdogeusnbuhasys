import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { UserRoutingModule } from './user-routing.module';
import { IndexComponent } from './index/index.component';
import { CreateComponent } from './create/create.component';
import { UpdateComponent } from './update/update.component';
import { BuyComponent } from './buy/buy.component';
import { LogComponent } from './log/log.component';
import { AddfundsComponent } from './addfunds/addfunds.component';


@NgModule({
  declarations: [IndexComponent, CreateComponent, UpdateComponent, BuyComponent, LogComponent, AddfundsComponent],
  imports: [
    CommonModule,
    UserRoutingModule
  ]
})
export class UserModule { }
