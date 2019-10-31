import { Component, OnInit } from '@angular/core';
import {FormControl, FormGroup} from '@angular/forms';
import {CurrentUserService, UserService} from '../../_services';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.scss']
})
export class CreateComponent implements OnInit {
  formdata: FormGroup;
  constructor(private userService: UserService, private currentUserService: CurrentUserService) { }

  ngOnInit() {
    this.formdata = new FormGroup({
      name: new FormControl(''),
      email: new FormControl(''),
      balance: new FormControl('0,00'),
      gravatar: new FormControl(false),
      pin: new FormControl('')
    });
  }

  onSubmit() {
    this.userService.create(this.formdata.value).subscribe(u => {
      this.currentUserService.set(u);
    }, err => {
      console.log(err);
    });
  }
}
