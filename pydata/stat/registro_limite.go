package stat

import "time"

type RegistroLimite struct {
	CriadoEm       *time.Time `db:"criado_em" json:"criado_em"`
	DmpProposto    string     `db:"dmp_proposto" json:"dmp_proposto"` // tipo json tratado como string bruto
	ID             int        `db:"id" json:"id"`
	IDRegistro     *int       `db:"id_registro" json:"id_registro"`
	IDUsuario      int        `db:"id_usuario" json:"id_usuario"`
	InputLimite    string     `db:"input_limite" json:"input_limite"`       // tipo json tratado como string bruto
	LimiteProposto string     `db:"limite_proposto" json:"limite_proposto"` // tipo json tratado como string bruto
	Usuario        string     `db:"usuario" json:"usuario"`
}
