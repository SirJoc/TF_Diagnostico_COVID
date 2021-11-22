import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import {MatCardModule} from "@angular/material/card";
import {MatIconModule} from "@angular/material/icon";
import {MatFormFieldModule} from "@angular/material/form-field";
import {UsersApiService} from "../service/users_api.service";
import {User} from "../models/data";
import {FormBuilder, FormGroup} from "@angular/forms";


@Component({
  selector: 'app-log-in',
  templateUrl: './log-in.component.html',
  styleUrls: ['./log-in.component.css']
})
export class LogInComponent implements OnInit {
  toppings: FormGroup;
  flag!: boolean;
  constructor(private usersApi: UsersApiService, private router: Router, fb:FormBuilder) {
    this.toppings = fb.group(
      {
        username: "",
        password: ""
      }
    )
  }
  userData: User = {} as User;
  leng_users = 1;

  ngOnInit(): void {
    this.usersApi.getAll().subscribe(response=> {
      this.leng_users = response.length;
    })
  }

  navigateToHome(): void {
    this.router.navigate(['/home'])
      .then(() => console.log('Navigated to Home'));
  }

  identifyUser() : void {
    this.flag = false;
    for(let i = 1; i<= this.leng_users;i++)
    {
      this.usersApi.getById(i).subscribe(response=>{
        console.log(response);
        if(response.username === this.toppings.value.username){
          if(response.password === this.toppings.value.password){
            this.navigateToHome();
            this.flag = true;
          }
        }
      })
    }
  }

}
