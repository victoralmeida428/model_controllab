package stat

import "time"

type EflmMeta struct {
	AnalyteID  int       `db:"analyte_id" json:"analyte_id"`
	Comments   *string   `db:"comments" json:"comments"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	Display    bool      `db:"display" json:"display"`
	ID         int       `db:"id" json:"id"`
	Lower      *float64  `db:"lower" json:"lower"`
	Median     *float64  `db:"median" json:"median"`
	NumberUsed *int      `db:"number_used" json:"number_used"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	Upper      *float64  `db:"upper" json:"upper"`
	VarType    string    `db:"var_type" json:"var_type"`
}
