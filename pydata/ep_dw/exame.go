package pydata

type Exame struct {
	ID                  int     `db:"id" json:"id"`
	IDModulo            int     `db:"id_modulo" json:"id_modulo"`
	IDExame             int     `db:"id_exame" json:"id_exame"`
	Exame               *string `db:"exame" json:"exame"`
	ExameCumulativo     *string `db:"exame_cumulativo" json:"exame_cumulativo"`
	ExameCert           *string `db:"exame_cert" json:"exame_cert"`
	ExameCertificado    *string `db:"exame_certificado" json:"exame_certificado"`
	ExameCumulativoCert *string `db:"exame_cumulativo_cert" json:"exame_cumulativo_cert"`
	IDExameCumulativo   *int16  `db:"id_exame_cumulativo" json:"id_exame_cumulativo"`
}
