import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import {MatCardModule} from "@angular/material/card";
import {MatIconModule} from "@angular/material/icon";
import {MatFormFieldModule} from "@angular/material/form-field";
import {UsersApiService} from "../service/users_api.service";
import {User} from "../models/data";


@Component({
  selector: 'app-log-in',
  templateUrl: './log-in.component.html',
  styleUrls: ['./log-in.component.css']
})
export class LogInComponent implements OnInit {
  user: User = { username:"", password:""};
  constructor(private usersApi: UsersApiService, private router: Router) { }
  userData: User = {} as User;
  leng_users = 0;
  isOk = false;
  isCorrect = false;
  arr_user!: User["username"];

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
    for(let i = 1; i<= this.leng_users;i++)
    {
      this.usersApi.getByUsername(username).subscribe(response=>{
        if(response.username === this.userData.username){
          if(response.password === this.userData.password){
            this.navigateToHome();
          }
        }
      })
    }
  }

}
