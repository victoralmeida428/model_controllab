package pydata

type Adequacao struct {
	ID               int      `db:"id" json:"id"`
	Ano              *int16   `db:"ano" json:"ano,omitempty"`
	IDGrupo          int16    `db:"id_grupo" json:"id_grupo"`
	Part             *int16   `db:"part" json:"part,omitempty"`
	IDModulo         *int16   `db:"id_modulo" json:"id_modulo,omitempty"`
	IDExame          *int16   `db:"id_exame" json:"id_exame,omitempty"`
	Adequados        *int16   `db:"adequados" json:"adequados,omitempty"`
	Possibilidades   *int16   `db:"possibilidades" json:"possibilidades,omitempty"`
	Total            *int16   `db:"total" json:"total,omitempty"`
	Esp              *int16   `db:"esp" json:"esp,omitempty"`
	NA21             *int16   `db:"na21" json:"na21,omitempty"`
	Subst            *int16   `db:"subst" json:"subst,omitempty"`
	SumNR            *int16   `db:"sum_nr" json:"sum_nr,omitempty"`
	ContEdu          *int16   `db:"cont_edu" json:"cont_edu,omitempty"`
	ContAnalito      *int16   `db:"cont_analito" json:"cont_analito,omitempty"`
	Edu              *int16   `db:"edu" json:"edu,omitempty"`
	ExameCertificado *string  `db:"exame_certificado" json:"exame_certificado,omitempty"`
	Adeq             *float64 `db:"adeq" json:"adeq,omitempty"`
	PctMinCert       *float64 `db:"pct_min_cert" json:"pct_min_cert,omitempty"`
	IDExameGrupo     *int16   `db:"id_exame_grupo" json:"id_exame_grupo,omitempty"`
	Rodadas          *int16   `db:"rodadas" json:"rodadas,omitempty"`
	MenorEnvio       *string  `db:"menor_envio" json:"menor_envio,omitempty"`
	MaiorEnvio       *string  `db:"maior_envio" json:"maior_envio,omitempty"`
	Elegivel         *int16   `db:"elegivel" json:"elegivel,omitempty"`
	Certificado      *int     `db:"certificado" json:"certificado,omitempty"`
	IDModExame       int      `db:"id_mod_exame" json:"id_mod_exame"`
	MaxObs           *int16   `db:"max_obs" json:"max_obs,omitempty"`
	ObsTotalValidas  *int16   `db:"obs_total_validas" json:"obs_total_validas,omitempty"`
	ACum             *int16   `db:"a_cum" json:"a_cum,omitempty"`
	TCum             *int16   `db:"t_cum" json:"t_cum,omitempty"`
	AEsp             *int16   `db:"a_esp" json:"a_esp,omitempty"`
	TEsp             *int16   `db:"t_esp" json:"t_esp,omitempty"`
	AdeqCum          *int16   `db:"adeq_cum" json:"adeq_cum,omitempty"`
	AdeqEsp          *int16   `db:"adeq_esp" json:"adeq_esp,omitempty"`
}
