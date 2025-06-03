package pydata

type ParceiroRede struct {
	ID            int    `db:"id" json:"id"`
	Login         string `db:"login" json:"login"`
	Participantes []int  `db:"participantes" json:"participantes"`
	Senha         string `db:"senha" json:"senha"`
}
