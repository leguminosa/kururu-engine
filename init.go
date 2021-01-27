package database

import (
	"log"

	"github.com/leguminosa/kururu-engine/driver"
	"github.com/leguminosa/kururu-engine/driver/orm"
)

func InitDB(cfg map[string]*DatabaseConnectionConfig) DatabaseAccessorItf {
	var (
		db = &databaseAccessor{
			connections: make(map[string]*database),
		}
	)

	for name, conn := range cfg {
		if conn == nil {
			continue
		}

		var (
			instances = &database{}
			err       error
		)

		log.Printf("Connecting to database %q with dialect %q...\n", name, conn.Dialect)
		instances.master, err = orm.Open(conn.Dialect, conn.Master)
		if err != nil {
			log.Fatalln("Failed connecting to master instance.", err.Error())
		}

		for _, slave := range conn.Slave {
			var slaveDB driver.OrmItf
			slaveDB, err = orm.Open(conn.Dialect, slave)
			if err != nil {
				log.Fatalln("Failed connecting to slave instance.", err.Error())
			}
			instances.slave = append(instances.slave, slaveDB)
		}

		db.connections[name] = instances
	}

	return db
}
