import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {IndexComponent} from './index/index.component';
import {BuyComponent} from './buy/buy.component';
import {CreateComponent} from './create/create.component';
import {AddfundsComponent} from './addfunds/addfunds.component';
import {LogComponent} from './log/log.component';
import {UpdateComponent} from './update/update.component';


const routes: Routes = [
  {
    path: '',
    component: IndexComponent
  },
  {
    path: ':id/buy',
    component: BuyComponent
  },
  {
    path: 'create',
    component: CreateComponent
  },
  {
    path: ':id/addfunds',
    component: AddfundsComponent
  },
  {
    path: ':id/log',
    component: LogComponent
  },
  {
    path: ':id/update',
    component: UpdateComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UserRoutingModule { }
