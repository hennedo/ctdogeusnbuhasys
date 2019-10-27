import { Component, OnInit } from '@angular/core';
import {FormControl, FormGroup} from '@angular/forms';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.scss']
})
export class CreateComponent implements OnInit {
  formdata: FormGroup;
  constructor() { }

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
    console.log(this.formdata.value);
  }
}
