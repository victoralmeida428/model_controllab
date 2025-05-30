package pydata

import "time"

type LogAtualizacao struct {
	ID       int        `db:"id" json:"id"`
	Nome     *string    `db:"nome" json:"nome"`
	OK       *int16     `db:"ok" json:"ok"`
	Date     *time.Time `db:"date" json:"date"`
	IDModulo *int16     `db:"id_modulo" json:"id_modulo"`
	Ano      *int16     `db:"ano" json:"ano"`
	DateI    *time.Time `db:"date_i" json:"date_i"`
	Error    *string    `db:"error" json:"error"`
}
