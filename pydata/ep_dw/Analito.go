package pydata

type Analito struct {
	Analito      *string `db:"analito" json:"analito"`
	ID           int     `db:"id" json:"id"`
	IDAnalito    *int    `db:"id_analito" json:"id_analito"`
	IDModAnalito *int    `db:"id_mod_analito" json:"id_mod_analito"`
	IDModExame   *int    `db:"id_mod_exame" json:"id_mod_exame"`
}
