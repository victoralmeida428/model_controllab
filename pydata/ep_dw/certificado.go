package pydata

import "time"

type Certificado struct {
	ID                  int        `db:"id" json:"id"`
	Ano                 int16      `db:"ano" json:"ano"`
	Part                *int       `db:"part" json:"part"`
	IDModulo            *int16     `db:"id_modulo" json:"id_modulo"`
	IDExame             *int16     `db:"id_exame" json:"id_exame"`
	Exame               *string    `db:"exame" json:"exame"`
	Segmento            *int16     `db:"segmento" json:"segmento"`
	IDSistema           *int16     `db:"id_sistema" json:"id_sistema"`
	IDCodAva            *int16     `db:"id_cod_ava" json:"id_cod_ava"`
	IDModuloCertificado *int16     `db:"id_modulo_certificado" json:"id_modulo_certificado"`
	ModuloCertificado   *string    `db:"modulo_certificado" json:"modulo_certificado"`
	GrauDesempenho      *float64   `db:"grau_desempenho" json:"grau_desempenho"`
	Educacao            *int16     `db:"educacao" json:"educacao"`
	Parcial             *bool      `db:"parcial" json:"parcial"`
	DataParcial         *time.Time `db:"data_parcial" json:"data_parcial"`
	IDModExame          *int       `db:"id_mod_exame" json:"id_mod_exame"`
}
