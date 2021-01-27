package database

import (
	"log"

	"github.com/leguminosa/kururu-engine/driver"
)

type (
	database struct {
		master driver.OrmItf
		slave  []driver.OrmItf
	}
)

func (db *database) GetMasterDB() driver.OrmItf {
	if db.master == nil {
		log.Fatalln("No master instance found.")
	}

	return db.master
}

func (db *database) GetSlaveDB(index int) driver.OrmItf {
	if db.slave == nil {
		log.Fatalln("No slave instance found.")
	}

	if index >= len(db.slave) {
		index = len(db.slave) - 1
	}

	return db.slave[index]
}
