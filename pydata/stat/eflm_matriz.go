package stat

import "time"

type EflmMatriz struct {
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	ID                int       `db:"id" json:"id"`
	MatrixExpansion   string    `db:"matrix_expansion" json:"matrix_expansion"`
	MatrixExplanation *string   `db:"matrix_explanation" json:"matrix_explanation"`
	Ordering          *bool     `db:"ordering" json:"ordering"`
	UpdatedAt         time.Time `db:"updated_at" json:"updated_at"`
}
