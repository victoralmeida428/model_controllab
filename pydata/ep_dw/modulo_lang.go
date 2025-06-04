package pydata

type ModuloLang struct {
	IDModulo        int    `json:"id_modulo" db:"id_modulo"`
	Lang            string `json:"lang" db:"lang"`
	Nome            string `json:"nome" db:"nome"`
	NomeCertificado string `json:"nome_cert" db:"nome_certificado"`
	ID              int    `json:"id" db:"id"`
}
