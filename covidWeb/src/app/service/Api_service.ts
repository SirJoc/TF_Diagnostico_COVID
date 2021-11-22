import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import {Form, Result} from "../models/data"

@Injectable({
  providedIn: 'root'
})
export class DataApiService {
  basePath = "http://localhost:8001";
  constructor(private http: HttpClient ) { }

  getAllData(userId: number): any {
    var link = this.basePath + `/api/users/${userId}/results`
    return this.http.get<any>(link)
  }
}
