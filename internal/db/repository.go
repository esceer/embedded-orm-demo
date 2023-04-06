package db

import (
	"github.com/go-pg/pg/v10"
)

type entityRepository struct {
	db *pg.DB
}

type EntityRepository interface {
	CreateEntity(*Entity) (*Entity, error)
	GetEntities() ([]*Entity, error)
	DeleteEntity(string) error
}

func NewEntityRepository(db *pg.DB) EntityRepository {
	return &entityRepository{
		db: db,
	}
}

func (d *entityRepository) CreateEntity(e *Entity) (*Entity, error) {
	_, err := d.db.Model(e).Insert()
	return e, err
}

func (d *entityRepository) GetEntities() (es []*Entity, err error) {
	err = d.db.Model(&es).Select()
	return
}

func (d *entityRepository) DeleteEntity(id string) error {
	e := &Entity{
		ID: id,
	}
	_, err := d.db.Model(e).WherePK().Delete()
	return err
}
