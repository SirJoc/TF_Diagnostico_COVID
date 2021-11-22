export interface Form{
  tos: boolean
  cefalea: boolean
  congNasal: boolean
  difRespiratoria: boolean
  dolorGarganta: boolean
  fiebre: boolean
  diarrea: boolean
  nauseas: boolean
  anosmiaPulmonar: boolean
  dolorAbdominal: boolean
  dolorArticulaciones: boolean
  dolorMuscular: boolean
  dolorPecho: boolean
  otros: boolean
}

export interface User{
  id: number;
  username: string
  password: string
}

export interface  Result{
  createdAt: string
  diagnostic: string
  userId: number
  form: Form
}

