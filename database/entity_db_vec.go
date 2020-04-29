package database

import (
	"github.com/kunashu/simple.mud/entity"
)

type EntityDatabaseVector []entity.Entity

func (db EntityDatabaseVector) Has(name string) bool {
	return db.Find(name) != nil
}

func (db EntityDatabaseVector) HasFull(name string) bool {
	return db.FindFull(name) != nil
}

func (db EntityDatabaseVector) Find(name string) entity.Entity {
	c := make(chan entity.Entity, 2)

	go func() {
		c <- db.FindFull(name)
	}()

	go func() {
		for _, e := range db {
			if e.Match(name) {
				c <- e
				return
			}
		}
		c <- nil
	}()

	var final entity.Entity
	for found := range c {
		if final != nil {
			if !final.FullMatch(name) && found != nil {
				final = found
			}
		} else {
			final = found
		}
	}

	return final
}

func (db EntityDatabaseVector) FindFull(name string) entity.Entity {
	for _, entity := range db {
		if entity.FullMatch(name) {
			return entity
		}
	}
	return nil
}
