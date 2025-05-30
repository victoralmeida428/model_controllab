package stat

import "time"

type RegistroComentario struct {
	Comentario *string    `db:"comentario" json:"comentario"`
	CriadoEm   *time.Time `db:"criado_em" json:"criado_em"`
	ID         int        `db:"id" json:"id"`
	IDRegistro int        `db:"id_registro" json:"id_registro"`
	IDUsuario  int        `db:"id_usuario" json:"id_usuario"`
	Title      string     `db:"title" json:"title"`
	Usuario    string     `db:"usuario" json:"usuario"`
}
