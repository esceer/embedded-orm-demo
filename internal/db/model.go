package db

type Entity struct {
	ID          string `pg:"id,pk"`
	Description string `pg:"description"`
}
