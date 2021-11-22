import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpHeaders } from "@angular/common/http";
import { Observable, throwError } from "rxjs";
import { User } from "../models/data";
import { catchError, retry } from "rxjs/operators";


@Injectable({providedIn: 'root'})
export class UsersApiService {
  basePath = 'http://localhost:8000/api/users';
  constructor(private http: HttpClient) {  }
  httpOptions = { headers: new HttpHeaders({ 'Content-Type': 'application/json'})};

  handleError(error: HttpErrorResponse): Observable<never> {
    if (error.error instanceof ErrorEvent) {
      console.log('An error ocurred: ', error.error.message);
    }
    else {
      console.log(`Backend returned code ${error.status}, body was: ${error.error}`);
    }
    return throwError('Something happened with request, please try again later.');
  }

  getAll(): Observable<any> {
    return this.http.get<any>(this.basePath, this.httpOptions)
      .pipe(retry(2), catchError(this.handleError));
  }

  getById(id: number): Observable<User> {
    return this.http.get<User>(`${this.basePath}/id/${id}`, this.httpOptions)
      .pipe(retry(2), catchError(this.handleError));
  }
}
