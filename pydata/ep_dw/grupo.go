package pydata

type Grupo struct {
	Id    int    `json:"id" db:"id"`
	Grupo string `json:"grupo" db:"grupo"`
}
