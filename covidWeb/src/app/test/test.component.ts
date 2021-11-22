import {Component, OnInit, ViewChild} from '@angular/core';
import {NgForm} from "@angular/forms";
import { Form } from '../models/data';
import {FormsApiService} from "../service/forms_api.service";
import {ActivatedRoute, Router} from "@angular/router";
import {Observable} from "rxjs";


@Component({
  selector: 'app-test',
  templateUrl: './test.component.html',
  styleUrls: ['./test.component.css']
})
export class TestComponent implements OnInit {
  @ViewChild('formForm', {static: false})
  formForm!: NgForm;
  formId!: number;
  formData: Form = {} as Form;
  resultado!: Observable<any>;


  constructor(private formApi: FormsApiService, private router: Router, private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.formId = Number(this.route.params.subscribe(params =>{
      if(params.id){
        const id = params.id;
      }
    }));
  }

  addForm(): void{
    const newForm = {tos: this.formData.tos, cefalea: this.formData.cefalea, congNasal: this.formData.congNasal,
                    difRespiratoria: this.formData.difRespiratoria, dolorGarganta: this.formData.dolorGarganta,
                    fiebre: this.formData.fiebre, diarrea: this.formData.diarrea, nauseas: this.formData.nauseas,
                    anosmiaPulmonar: this.formData.anosmiaPulmonar, dolorAbdominal: this.formData.anosmiaPulmonar,
                    dolorArticulaciones: this.formData.dolorArticulaciones, dolorMuscular: this.formData.dolorMuscular,
                    dolorPecho: this.formData.dolorPecho, otros: this.formData.otros};
    this.resultado = this.formApi.addForm(this.formId, newForm);
    console.log(newForm);
    //this.resultado.subscribe((response: String) => {
    //  console.log(response);
    //});
  }

  onSubmit(): void{
    if (this.formForm.form.valid){
      this.addForm();
    }
  }

}
