package pydata

type Parceiros struct {
	NomeFantasia       *string  `db:"nome_fantasia" json:"nome_fantasia"`
	GrupoEmpresarial   *string  `db:"grupo_empresarial" json:"grupo_empresarial"`
	GrupoRepresentacao *string  `db:"grupo_representacao" json:"grupo_representacao"`
	AtivoEp            int16    `db:"ativo_ep" json:"ativo_ep"`
	AtivoPbil          int16    `db:"ativo_pbil" json:"ativo_pbil"`
	AtivoCi            int16    `db:"ativo_ci" json:"ativo_ci"`
	AtivoGeral         int16    `db:"ativo_geral" json:"ativo_geral"`
	Latitude           *float32 `db:"latitude" json:"latitude"`
	Longitude          *float32 `db:"longitude" json:"longitude"`
	Part               int      `db:"part" json:"part"`
	Cidade             *string  `db:"cidade" json:"cidade"`
	Estado             *string  `db:"estado" json:"estado"`
	Pais               *string  `db:"pais" json:"pais"`
	RazaoSocial        *string  `db:"razao_social" json:"razao_social"`
	NumeroCn           *string  `db:"numero_cn" json:"numero_cn"`
	AtivoPbilGratuito  int16    `db:"ativo_pbil_gratuito" json:"ativo_pbil_gratuito"`
	InstituicaoPublica int16    `db:"instituicao_publica" json:"instituicao_publica"`
	Logradouro         *string  `db:"logradouro" json:"logradouro"`
	Bairro             *string  `db:"bairro" json:"bairro"`
	Cep                *string  `db:"cep" json:"cep"`
	IDGrupo            *int16   `db:"id_grupo" json:"id_grupo"`
}
