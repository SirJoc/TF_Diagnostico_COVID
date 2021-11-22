import { Component, OnInit } from '@angular/core';
import {Form, Result} from '../models/data'
import {ResultsApiService} from "../service/results_api.service";
import {isElementScrolledOutsideView} from "@angular/cdk/overlay/position/scroll-clip";

export interface PeriodicElement {
  name: string;
  position: number;
  weight: number;
  symbol: string;
}



@Component({
  selector: 'app-historic',
  templateUrl: './historic.component.html',
  styleUrls: ['./historic.component.css']
})
export class HistoricComponent implements OnInit {
  elements :  Result[] = [];

  constructor(private data : ResultsApiService) { }

  ngOnInit(): void {
    this.data.getResultsByUserId(1)
      .subscribe( (data: any) => {
        this.elements = data
        //console.log(data)
        console.log(this.elements)
      });
  }

  test() {
    this.elements.pop()
    console.log(this.elements)
  }

  displayedColumns: string[] = ['id', 'fecha', 'diagnostico'];
}
