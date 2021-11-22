import {Component, OnInit, ViewChild} from '@angular/core';
import {FormBuilder, FormGroup, NgForm} from "@angular/forms";
import { Form } from '../models/data';
import {FormsApiService} from "../service/forms_api.service";
import {ActivatedRoute, Router} from "@angular/router";
import {Observable} from "rxjs";
import {JsonPipe} from "@angular/common";
import {parseJson} from "@angular/cli/utilities/json-file";


@Component({
  selector: 'app-test',
  templateUrl: './test.component.html',
  styleUrls: ['./test.component.css']
})
export class TestComponent implements OnInit {
  @ViewChild('formForm', {static: false})
  formForm!: NgForm;
  formId!: number;
  data: Form = {} as Form;
  toppings: FormGroup;
  resultado!: Observable<any>;
  realResultado!: String;


  constructor(private formApi: FormsApiService, private router: Router, private route: ActivatedRoute, fb:FormBuilder) {
    this.toppings = fb.group(
      {
        tos: false,
        cefalea: false,
        congNasal: false,
        difRespiratoria : false,
        dolorGarganta : false,
        fiebre :false,
        diarrea : false,
        nauseas : false,
        anosmiaPulmonar :false,
        dolorAbdominal : false,
        dolorArticulaciones : false,
        dolorMuscular : false,
        dolorPecho : false,
        otros : false
      }
    );
  }

  ngOnInit(): void {
    this.formId = Number(this.route.params.subscribe(params =>{
      if(params.id){
        const id = params.id;
      }
    }));

  }

  addForm(): void{
    const newForm = {tos: this.toppings.value.tos, cefalea: this.toppings.value.cefalea, congNasal: this.toppings.value.congNasal,
      difRespiratoria: this.toppings.value.difRespiratoria, dolorGarganta: this.toppings.value.dolorGarganta,
      fiebre: this.toppings.value.fiebre, diarrea: this.toppings.value.diarrea, nauseas: this.toppings.value.nauseas,
      anosmiaPulmonar: this.toppings.value.anosmiaPulmonar, dolorAbdominal: this.toppings.value.anosmiaPulmonar,
      dolorArticulaciones: this.toppings.value.dolorArticulaciones, dolorMuscular: this.toppings.value.dolorMuscular,
      dolorPecho: this.toppings.value.dolorPecho, otros: this.toppings.value.otros};

    console.log(newForm);
    this.resultado = this.formApi.addForm(this.formId, newForm);
    this.resultado.subscribe((response: String) => {
      console.log(response);
      this.realResultado = response;
    });

  }

  onSubmit(): void{
    if (this.formForm.form.valid){
      this.addForm();
    }
  }

}
