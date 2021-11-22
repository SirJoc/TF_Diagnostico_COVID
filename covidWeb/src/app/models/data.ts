export interface Form{
  Tos: boolean
  Cefalea: boolean
  cong_nasal: boolean
  dif_respiratoria: boolean
  dolor_garganta: boolean
  fiebre: boolean
  diarrea: boolean
  nauseas: boolean
  anosmia_pulmonar: boolean
  dolor_abdominal: boolean
  dolor_articulaciones: boolean
  dolor_muscular: boolean
  dolor_pecho: boolean
  otros: boolean
}

export interface User{
  firstName: string
  lastname: string
  username: string
  password: string
}

export interface  Result{
  CreatedAt: string
  Diagnostic: string
  UserId: number
  form: Form
}

