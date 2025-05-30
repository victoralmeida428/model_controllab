package stat

type EstatisticaQuanti struct {
	ID             int      `db:"id" json:"id"`
	Ano            int16    `db:"ano" json:"ano"`
	Analito        *string  `db:"analito" json:"analito"`
	CodAva         *int     `db:"cod_ava" json:"cod_ava"`
	CV             *float64 `db:"cv" json:"cv"`
	DP             *float64 `db:"dp" json:"dp"`
	Envio          *string  `db:"envio" json:"envio"`
	FaixaInf       *float64 `db:"faixa_inf" json:"faixa_inf"`
	FaixaSup       *float64 `db:"faixa_sup" json:"faixa_sup"`
	GrupoAnalitico *string  `db:"grupo_analitico" json:"grupo_analitico"`
	GrupoAva       *string  `db:"grupo_ava" json:"grupo_ava"`
	IDAnalito      *int     `db:"id_analito" json:"id_analito"`
	IDExame        *int     `db:"id_exame" json:"id_exame"`
	IDGrupoAva     *int     `db:"id_grupo_ava" json:"id_grupo_ava"`
	IDModulo       *int16   `db:"id_modulo" json:"id_modulo"`
	ItemEnsaio     *string  `db:"item_ensaio" json:"item_ensaio"`
	LimRobustoInf  *float64 `db:"lim_robusto_inf" json:"lim_robusto_inf"`
	LimRobustoSup  *float64 `db:"lim_robusto_sup" json:"lim_robusto_sup"`
	Media          *float64 `db:"media" json:"media"`
	Mediana        *float64 `db:"mediana" json:"mediana"`
	NumItem        *int     `db:"num_item" json:"num_item"`
	QtdParts       *int     `db:"qtd_parts" json:"qtd_parts"`
	QtdResultados  *int     `db:"qtd_resultados" json:"qtd_resultados"`
	Teste          *bool    `db:"teste" json:"teste"`
}
