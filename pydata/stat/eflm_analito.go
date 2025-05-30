package stat

import "time"

type EflmAnalito struct {
	Comments    *string   `db:"comments" json:"comments"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	Description string    `db:"description" json:"description"`
	DisplayName string    `db:"display_name" json:"display_name"`
	FullName    *string   `db:"full_name" json:"full_name"`
	ID          int       `db:"id" json:"id"`
	IsDefault   bool      `db:"is_default" json:"is_default"`
	NlmcID      string    `db:"nlmc_id" json:"nlmc_id"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
