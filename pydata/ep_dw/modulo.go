package pydata

type Modulo struct {
	IDSegmento *int16  `db:"id_segmento" json:"id_segmento"`
	IDModulo   int16   `db:"id_modulo" json:"id_modulo"`
	IDArea     *int16  `db:"id_area" json:"id_area"`
	Area       *string `db:"area" json:"area"`
	Segmento   *string `db:"segmento" json:"segmento"`
	Modulo     *string `db:"modulo" json:"modulo"`
	ModuloCert *string `db:"modulo_cert" json:"modulo_cert"`
}
