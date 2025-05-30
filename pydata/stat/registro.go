package stat

import "time"

type Registro struct {
	Analito       string     `db:"analito" json:"analito"`
	AnalitoVBA    *int       `db:"analito_vba" json:"analito_vba"`
	AnoF          int        `db:"ano_f" json:"ano_f"`
	AnoI          int        `db:"ano_i" json:"ano_i"`
	AprovadoPor   *string    `db:"aprovado_por" json:"aprovado_por"`
	CriadoEm      *time.Time `db:"criado_em" json:"criado_em"`
	EncerradoEm   *time.Time `db:"encerrado_em" json:"encerrado_em"`
	ID            int        `db:"id" json:"id"`
	IDAprovadoPor *int       `db:"id_aprovado_por" json:"id_aprovado_por"`
	IDModAnalito  []int      `db:"id_mod_analito" json:"id_mod_analito"`
	IDModulo      []int      `db:"id_modulo" json:"id_modulo"`
	IDUsuario     int        `db:"id_usuario" json:"id_usuario"`
	MatrizVBA     *int       `db:"matriz_vba" json:"matriz_vba"`
	Status        int16      `db:"status" json:"status"`
	Usuario       string     `db:"usuario" json:"usuario"`
}
