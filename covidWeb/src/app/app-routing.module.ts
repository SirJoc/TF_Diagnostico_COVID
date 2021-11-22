import { NgModule } from '@angular/core';
import {RouterModule, Routes} from "@angular/router";
import {LogInComponent} from "./log-in/log-in.component";
import {TestComponent} from "./test/test.component";
import {HistoricComponent} from "./historic/historic.component";


const routes: Routes = [
  {path: '', component: LogInComponent},
  {path: 'home', component: TestComponent},
  {path: 'historic', component: HistoricComponent}
]

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
