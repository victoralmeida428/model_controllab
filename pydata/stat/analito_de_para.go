package stat

type AnalitoDePara struct {
	ID            int  `db:"id" json:"id"`
	IDAnalito     *int `db:"id_analito" json:"id_analito"`
	IDAnalitoEflm *int `db:"id_analito_eflm" json:"id_analito_eflm"`
}
