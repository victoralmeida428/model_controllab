package pydata

type ExameLang struct {
	IDModulo        int    `json:"id_modulo" db:"id_modulo"`
	IDExame         int    `json:"id_exame" db:"id_exame"`
	Lang            string `json:"lang" db:"lang"`
	Nome            string `json:"nome" db:"nome"`
	NomeCertificado string `json:"nome_certificado" db:"nome_certificado"`
	ID              int    `json:"id" db:"id"`
	IDModExame      int    `json:"id_mod_exame" db:"id_mod_exame"`
}
