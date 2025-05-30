package stat

type EflmAnalitoMatriz struct {
	Id        int  `json:"id" db:"id"`
	IdAnalito *int `json:"id_analito" db:"id_analito"`
	IdMatriz  *int `json:"id_matriz" db:"id_matriz"`
}
